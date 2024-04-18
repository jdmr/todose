package main

import (
	"context"
	"crypto/rsa"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var privKey *rsa.PrivateKey
var pubKey *rsa.PublicKey

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}
	ctx := context.Background()
	client, err = getMongoClient(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s\n", err)
	}

	pem := viper.GetString("rsa.private")
	privKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(pem))
	if err != nil {
		log.Fatalf("Error parsing private key: %s\n", err)
	}

	pem = viper.GetString("rsa.public")
	pubKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
	if err != nil {
		log.Fatalf("Error parsing public key: %s\n", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/healthz", getHealthz).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/users", getUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users", createUser).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users/{userID}", getUser).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users/{userID}", updateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/users/{userID}", deleteUser).Methods(http.MethodDelete)

	router.HandleFunc("/api/v1/todos", getTodos).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/todos", createTodo).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/todos/{todoID}", getTodo).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/todos/{todoID}", updateTodo).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/todos/{todoID}", deleteTodo).Methods(http.MethodDelete)

	router.HandleFunc("/api/v1/login", getLogin).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		log.Println("Starting server on :8080")
		log.Fatal(srv.ListenAndServe())
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	stopCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Disconnect(stopCtx)
	if err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %s\n", err)
	}
	log.Println("Disconnected from MongoDB!")

	err = srv.Shutdown(stopCtx)
	if err != nil {
		log.Fatalf("Error shutting down server: %s\n", err)
	}
	log.Println("Server gracefully stopped!")

	cancel()
}
