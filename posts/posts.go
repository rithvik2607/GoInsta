package posts

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	generate "github.com/rithvik2607/GoInsta/genrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	Id        string `bson:"id" json:"id"`
	User_id   string `bson:"user_id" json:"user_id"`
	Caption   string `bson:"caption" json:"caption"`
	Img_url   string `bson:"img_url" json:"img_url"`
	Timestamp string `bson:"timestamp" json:"timestamp"`
}

type PostInfo struct {
	User_id string `bson:"user_id" json:"user_id"`
	Caption string `bson:"caption" json:"caption"`
	Img_url string `bson:"img_url" json:"img_url"`
}

// Database instance
var collection *mongo.Collection

func PostCollection(c *mongo.Database) {
	collection = c.Collection("posts")
}

/*
GetPost - collects posts from DB and matches ID
 of posts with the ID given in the link
*/
func GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Split the URL to obtain the ID, store it in postId
		p := strings.Split(r.URL.Path, "/")
		var postId string
		if len(p) <= 1 {
			log.Fatal("id not found")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			postId = p[2]
		}

		// Initialize posts array and collect posts from DB
		posts := []Post{}
		cursor, err := collection.Find(context.TODO(), bson.D{{}})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Iterate through posts and match ID with postId
		for cursor.Next(context.TODO()) {
			var post Post
			cursor.Decode(&post)
			if post.Id == postId {
				posts = append(posts, post)
			}
		}

		// Convert the result to JSON
		jsonBytes, err := json.Marshal(posts)
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
CreatePost - receives data in JSON format,
creates ID, records timestamp and creates Post object.
Post object is then converted to JSON and added to DB
*/
func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Create PostInfo object
		var post PostInfo

		// Decode the input JSON and match it with PostInfo struct
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Generate ID and record timestamp
		Id := generate.GenId()
		Timestamp := time.Now().Format(time.RFC3339)

		// Create Post object, name it newPost
		newPost := Post{
			Id:        Id,
			User_id:   post.User_id,
			Caption:   post.Caption,
			Img_url:   post.Img_url,
			Timestamp: Timestamp,
		}

		// Insert newPost into DB
		res, err := collection.InsertOne(context.TODO(), newPost)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		log.Println(res) // Handling response from Mongo driver

		// Convert newPost to JSON
		jsonBytes, err := json.Marshal(newPost)
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
GetUsersPost - collects posts from DB and compares
ID of the user who posted it with the ID of given
in the link
*/
func GetUsersPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Split the URL to obtain the ID, store it in userId
		p := strings.Split(r.URL.Path, "/")
		var userId string
		if len(p) <= 1 {
			log.Fatal("id not found")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			userId = p[3]
		}

		// Initialize posts array and collect posts from DB
		posts := []Post{}
		cursor, err := collection.Find(context.TODO(), bson.D{{}})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Iterate through posts and match User_iD with userId
		for cursor.Next(context.TODO()) {
			var post Post
			cursor.Decode(&post)
			if post.User_id == userId {
				posts = append(posts, post)
			}
		}

		// Convert the result to JSON
		jsonBytes, err := json.Marshal(posts)
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
