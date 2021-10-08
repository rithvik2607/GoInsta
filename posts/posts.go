package posts

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	id        string `json:"id"`
	caption   string `json:"name"`
	img_url   string `json:"email"`
	timestamp string `json:"password"`
}

// Database instance
var collection *mongo.Collection

func PostCollection(c *mongo.Database) {
	collection = c.Collection("posts")
}
