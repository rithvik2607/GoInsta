package posts

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	generate "github.com/rithvik2607/GoInsta/genrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	id        string `json:"id"`
	user_id   string `json:"user_id"`
	caption   string `json:"caption"`
	img_url   string `json:"image_url"`
	timestamp string `json:"timestamp"`
}

type PostInfo struct {
	user_id string `json:"user_id"`
	caption string `json:"caption"`
	img_url string `json:"image_url"`
}

// Database instance
var collection *mongo.Collection

func PostCollection(c *mongo.Database) {
	collection = c.Collection("posts")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	var postId string
	if len(p) <= 1 {
		log.Fatal("id not found")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		postId = strconv.Quote(p[2])
	}

	post := Post{}
	err := collection.FindOne(context.TODO(), bson.M{"id": postId}).Decode(&post)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(post)
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post PostInfo

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	Id := generate.GenId()
	Timestamp := time.Now().Format(time.RFC3339)

	newPost := Post{
		id:        Id,
		user_id:   post.user_id,
		caption:   post.caption,
		img_url:   post.img_url,
		timestamp: Timestamp,
	}

	res, err := collection.InsertOne(context.TODO(), newPost)

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

func GetUsersPost(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	var userId string
	if len(p) <= 1 {
		log.Fatal("id not found")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		userId = strconv.Quote(p[3])
	}

	posts := []Post{}
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userId})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for cursor.Next(context.TODO()) {
		var post Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	jsonBytes, err := json.Marshal(posts)
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
