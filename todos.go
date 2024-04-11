package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

type Todo struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Owner  *User  `json:"owner"`
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting todos...")
	if !validUser(r) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	coll := client.Database(viper.GetString("mongo.db")).Collection("todos")
	cursor, err := coll.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, "could not find todos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())
	todos := []*Todo{}
	for cursor.Next(r.Context()) {
		todo := &Todo{}
		err := cursor.Decode(todo)
		if err != nil {
			http.Error(w, "could not decode todo: "+err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	err = cursor.Err()
	if err != nil {
		http.Error(w, "cursor error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "could not encode todos: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting todo...")
	params := mux.Vars(r)
	todoID := params["todoID"]
	if todoID == "" {
		log.Println("todoID is required")
		http.Error(w, "todoID is required", http.StatusBadRequest)
		return
	}
	coll := client.Database(viper.GetString("mongo.db")).Collection("todos")
	todo := &Todo{}
	err := coll.FindOne(r.Context(), bson.M{"_id": todoID}).Decode(todo)
	if err != nil {
		http.Error(w, "could not find todo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		http.Error(w, "could not encode todo: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating todo...")
	todo := &Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		http.Error(w, "could not decode todo: "+err.Error(), http.StatusBadRequest)
		return
	}
	coll := client.Database(viper.GetString("mongo.db")).Collection("todos")
	_, err = coll.InsertOne(r.Context(), todo)
	if err != nil {
		http.Error(w, "could not create todo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		http.Error(w, "could not encode todo: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating todo...")
	params := mux.Vars(r)
	todoID := params["todoID"]
	if todoID == "" {
		log.Println("todoID is required")
		http.Error(w, "todoID is required", http.StatusBadRequest)
		return
	}
	todo := &Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		http.Error(w, "could not decode todo: "+err.Error(), http.StatusBadRequest)
		return
	}
	coll := client.Database(viper.GetString("mongo.db")).Collection("todos")
	_, err = coll.ReplaceOne(r.Context(), bson.M{"_id": todoID}, todo)
	if err != nil {
		http.Error(w, "could not update todo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		http.Error(w, "could not encode todo: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting todo...")
	params := mux.Vars(r)
	todoID := params["todoID"]
	if todoID == "" {
		log.Println("todoID is required")
		http.Error(w, "todoID is required", http.StatusBadRequest)
		return
	}
	coll := client.Database(viper.GetString("mongo.db")).Collection("todos")
	_, err := coll.DeleteOne(r.Context(), bson.M{"_id": todoID})
	if err != nil {
		http.Error(w, "could not delete todo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
