package users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	generate "github.com/rithvik2607/GoInsta/genrate"
	"github.com/rithvik2607/GoInsta/hashing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type UserInfo struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

// Database instance
var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

/*
GetUser - collects users from DB and matches ID
 of users with the ID given in the link
*/
func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Split the URL to obtain the ID, store it in userId
		p := strings.Split(r.URL.Path, "/")
		var userId string
		if len(p) <= 1 {
			log.Fatal("id not found")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			userId = p[2]
		}

		// Initialize users array and collect users from DB
		users := []User{}
		cursor, err := collection.Find(context.TODO(), bson.D{{}})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Iterate through users and match ID with userId
		for cursor.Next(context.TODO()) {
			var user User
			cursor.Decode(&user)
			if user.Id == userId {
				users = append(users, user)
			}
		}

		// Convert the result to JSON
		jsonBytes, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Structure the response
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)

		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

/*
CreateUser - receives data in JSON format,
creates ID, hashes password and creates User object.
User object is then converted to JSON and added to DB
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Create UserInfo object
		var user UserInfo

		// Decode the input JSON and match it with UserInfo struct
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Retrieve salt from environment variables
		salt, ok := os.LookupEnv("salt")
		if ok != true {
			log.Fatal("error: unable to find uri in the environment")
		}

		// Generate ID and hashed password
		Id := generate.GenId()
		newSalt := []byte(salt)
		user.Password = hashing.HashPassword(user.Password, newSalt)

		// Create User object, name it newUser
		newUser := User{
			Id:       Id,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}

		// Insert newUser into DB
		res, err := collection.InsertOne(context.TODO(), newUser)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		log.Println(res) // Handling response from Mongo driver

		// Convert newUser to JSON
		jsonBytes, err := json.Marshal(newUser)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Structure the response
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)

		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
