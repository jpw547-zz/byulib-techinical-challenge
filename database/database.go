package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllPosts() ([]BlogPost, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %s", err.Error())
		return []BlogPost{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("blogPosts").Collection("posts")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		cursor.Close(ctx)
		return []BlogPost{}, err
	}
	defer cursor.Close(ctx)

	var results []BlogPost

	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	fmt.Println(results)
	return results, nil
}

func GetPostByTitle(postTitle string) (BlogPost, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %s", err.Error())
		return BlogPost{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("blogPosts").Collection("posts")
	var result BlogPost

	err = collection.FindOne(ctx, bson.D{{"title", postTitle}}).Decode(&result)
	if err != nil {
		return BlogPost{}, err
	}

	return result, nil
}

func AddNewPost(data BlogPost) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %s", err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("blogPosts").Collection("posts")

	_, err = collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	fmt.Println("Inserted the document")
	return nil
}
