package main

import (
	"context"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

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
