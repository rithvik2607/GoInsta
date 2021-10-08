package users

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	id       string `json:"id"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}

// Database instance
var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

}
