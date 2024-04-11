package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}
}

func TestGetTodos(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create a todo
	coll := getTodosCollection(client)
	user := &User{
		ID:   "testuser",
		Name: "Alice",
	}
	todo := &Todo{
		ID:     "testtodo",
		Title:  "Test Todo",
		Status: "status",
		Owner:  user,
	}
	_, err = coll.UpdateOne(
		ctx,
		bson.M{"_id": todo.ID},
		bson.M{"$set": todo},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		t.Fatalf("Error creating todo: %s\n", err)
	}

	// Get the todos
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/todos", nil)
	getTodos(w, r)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the todo
	todos := []*Todo{}
	err = json.NewDecoder(w.Body).Decode(&todos)
	if err != nil {
		t.Fatalf("Error decoding response: %s\n", err)
	}
	if len(todos) < 1 {
		t.Errorf("Expected at least one todo, got %d", len(todos))
	}
	found := false
	for _, todo := range todos {
		if todo.ID == "testtodo" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected to find testtodo, got %v", todos)
	}

	coll.DeleteOne(ctx, bson.M{"_id": "testtodo"})
}

func TestGetTodo(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create a todo
	coll := getTodosCollection(client)
	user := &User{
		ID:   "testuser",
		Name: "Alice",
	}
	todo := &Todo{
		ID:     "testtodo",
		Title:  "Test Todo",
		Status: "status",
		Owner:  user,
	}
	_, err = coll.UpdateOne(
		ctx,
		bson.M{"_id": todo.ID},
		bson.M{"$set": todo},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		t.Fatalf("Error creating todo: %s\n", err)
	}

	// Get the todo
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/todos/testtodo", nil)
	r = mux.SetURLVars(r, map[string]string{"todoID": "testtodo"})
	getTodo(w, r)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the todo
	todo = &Todo{}
	err = json.NewDecoder(w.Body).Decode(todo)
	if err != nil {
		t.Fatalf("Error decoding response: %s\n", err)
	}
	if todo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got '%s'", todo.Title)
	}
	if todo.Owner.Name != "Alice" {
		t.Errorf("Expected name Alice, got %s", todo.Owner.Name)
	}

	coll.DeleteOne(ctx, bson.M{"_id": "testtodo"})
}

func TestCreateTodo(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create the todo
	body := `{"id":"testtodo","title":"Test Todo","status":"status","owner":{"id":"testuser","name":"Alice"}}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/v1/todos", strings.NewReader(body))
	createTodo(w, r)

	if w.Code != 201 {
		t.Errorf("Expected status code 201, got %d", w.Code)
	}

	// Check the todo
	todo := &Todo{}
	coll := getTodosCollection(client)
	err = coll.FindOne(ctx, bson.M{"_id": "testtodo"}).Decode(todo)
	if err != nil {
		t.Fatalf("Error finding todo: %s\n", err)
	}
	if todo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got '%s'", todo.Title)
	}
	if todo.Owner.Name != "Alice" {
		t.Errorf("Expected name Alice, got %s", todo.Owner.Name)
	}

	coll.DeleteOne(ctx, bson.M{"_id": "testtodo"})
}

func TestUpdateTodo(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create a todo
	coll := getTodosCollection(client)
	user := &User{
		ID:   "testuser",
		Name: "Alice",
	}
	todo := &Todo{
		ID:     "testtodo",
		Title:  "Test Todo",
		Status: "status",
		Owner:  user,
	}
	_, err = coll.UpdateOne(
		ctx,
		bson.M{"_id": todo.ID},
		bson.M{"$set": todo},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		t.Fatalf("Error creating todo: %s\n", err)
	}

	// Update the todo
	body := `{"id":"testtodo","title":"Testing Update","status":"status","owner":{"id":"testuser","name":"Alice"}}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/api/v1/todos/testtodo", strings.NewReader(body))
	r = mux.SetURLVars(r, map[string]string{"todoID": "testtodo"})
	updateTodo(w, r)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the todo
	todo = &Todo{}
	err = coll.FindOne(ctx, bson.M{"_id": "testtodo"}).Decode(todo)
	if err != nil {
		t.Fatalf("Error finding todo: %s\n", err)
	}
	if todo.Title != "Testing Update" {
		t.Errorf("Expected title 'Testing Update', got '%s'", todo.Title)
	}

	coll.DeleteOne(ctx, bson.M{"_id": "testtodo"})
}

func TestDeleteTodo(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create a todo
	coll := getTodosCollection(client)
	user := &User{
		ID:   "testuser",
		Name: "Alice",
	}
	todo := &Todo{
		ID:     "testtodo",
		Title:  "Test Todo",
		Status: "status",
		Owner:  user,
	}
	_, err = coll.UpdateOne(
		ctx,
		bson.M{"_id": todo.ID},
		bson.M{"$set": todo},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		t.Fatalf("Error creating todo: %s\n", err)
	}

	// Delete the todo
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/api/v1/todos/testtodo", nil)
	r = mux.SetURLVars(r, map[string]string{"todoID": "testtodo"})
	deleteTodo(w, r)

	if w.Code != 204 {
		t.Errorf("Expected status code 204, got %d", w.Code)
	}

	// Check the todo
	todo = &Todo{}
	err = coll.FindOne(ctx, bson.M{"_id": "testtodo"}).Decode(todo)
	if err == nil {
		t.Fatalf("Error not finding todo: %s\n", err)
	}
}
