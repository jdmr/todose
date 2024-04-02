package main

import (
	"context"
	"log"
	"net/http/httptest"
	"testing"

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
	todo = &Todo{}
	err = coll.FindOne(ctx, bson.M{"_id": "testtodo"}).Decode(todo)
	if err != nil {
		t.Fatalf("Error finding todo: %s\n", err)
	}
	if todo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got '%s'", todo.Title)
	}
}

/* func TestGetUser(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create a user
	coll := getUsersCollection(client)
	user := &User{
		ID:   "testuser",
		Name: "Alice",
	}
	_, err = coll.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		t.Fatalf("Error creating user: %s\n", err)
	}

	// Get the user
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/users/testuser", nil)
	r = mux.SetURLVars(r, map[string]string{"userID": "testuser"})
	getUser(w, r)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the user
	user = &User{}
	err = coll.FindOne(ctx, bson.M{"_id": "testuser"}).Decode(user)
	if err != nil {
		t.Fatalf("Error finding user: %s\n", err)
	}
	if user.Name != "Alice" {
		t.Errorf("Expected name Alice, got %s", user.Name)
	}
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create the user
	body := `{"id":"testuser2","name": "Bob"}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/v1/users", strings.NewReader(body))
	createUser(w, r)

	if w.Code != 201 {
		t.Errorf("Expected status code 201, got %d", w.Code)
	}

	// Check the user
	user := &User{}
	coll := getUsersCollection(client)
	err = coll.FindOne(ctx, bson.M{"_id": "testuser2"}).Decode(user)
	if err != nil {
		t.Fatalf("Error finding user: %s\n", err)
	}
	if user.Name != "Bob" {
		t.Errorf("Expected name Bob, got %s", user.Name)
	}
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create a user
	coll := getUsersCollection(client)
	user := &User{
		ID:   "testuser",
		Name: "Alice",
	}
	_, err = coll.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		t.Fatalf("Error creating user: %s\n", err)
	}

	// Update the user
	body := `{"id":"testuser","name": "Bob"}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/api/v1/users/testuser", strings.NewReader(body))
	r = mux.SetURLVars(r, map[string]string{"userID": "testuser"})
	updateUser(w, r)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the user
	user = &User{}
	err = coll.FindOne(ctx, bson.M{"_id": "testuser"}).Decode(user)
	if err != nil {
		t.Fatalf("Error finding user: %s\n", err)
	}
	if user.Name != "Bob" {
		t.Errorf("Expected name Bob, got %s", user.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()
	var err error
	client, err = getMongoClient(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %s\n", err)
	}
	defer client.Disconnect(ctx)

	// Create a user
	coll := getUsersCollection(client)
	user := &User{
		ID:   "testuser",
		Name: "Alice",
	}
	_, err = coll.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		t.Fatalf("Error creating user: %s\n", err)
	}

	// Delete the user
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/api/v1/users/testuser", nil)
	r = mux.SetURLVars(r, map[string]string{"userID": "testuser"})
	deleteUser(w, r)

	if w.Code != 204 {
		t.Errorf("Expected status code 204, got %d", w.Code)
	}

	// Check the user
	user = &User{}
	err = coll.FindOne(ctx, bson.M{"_id": "testuser"}).Decode(user)
	if err == nil {
		t.Fatalf("Error not finding user: %s\n", err)
	}
	if user.Name == "Alice" {
		t.Errorf("Expected nothing, got %s", user.Name)
	}
} */
