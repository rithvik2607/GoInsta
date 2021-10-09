package users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	generate "github.com/rithvik2607/GoInsta/genrate"
	"github.com/rithvik2607/GoInsta/hashing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	id       string `json:"id"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}

type UserInfo struct {
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}

// Database instance
var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	var userId string
	if len(p) <= 1 {
		log.Fatal("id not found")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		userId = strconv.Quote(p[2])
	}

	user := User{}
	err := collection.FindOne(context.TODO(), bson.M{"id": userId}).Decode(&user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserInfo

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	salt, ok := os.LookupEnv("salt")
	if ok != true {
		log.Fatal("error: unable to find uri in the environment")
	}
	Id := generate.GenId()
	newSalt := []byte(salt)
	user.password = hashing.HashPassword(user.password, newSalt)

	newUser := User{
		id:       Id,
		name:     user.name,
		email:    user.email,
		password: user.password,
	}

	res, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

	return
}
