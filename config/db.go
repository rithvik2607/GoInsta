package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rithvik2607/GoInsta/posts"
	"github.com/rithvik2607/GoInsta/users"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() {
	// DB config
	mongoUri, ok := os.LookupEnv("MONGODBURI")
	if ok != true {
		log.Fatal("error: unable to find uri in the environment")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}

	// Set up context required by mongo.Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//To close the connection at the end
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to MongoDB!")
	}

	db := client.Database("appData")
	users.UserCollection(db)
	posts.PostCollection(db)
	return
}
