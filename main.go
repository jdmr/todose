package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}
	ctx := context.Background()
	mongoURL := viper.GetString("mongo.prod")
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	log.Println("Connected to MongoDB!")
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s\n", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/healthz", getHealthz).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/users", getUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users", createUser).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users/{userID}", getUser).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users/{userID}", updateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/users/{userID}", deleteUser).Methods(http.MethodDelete)

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
