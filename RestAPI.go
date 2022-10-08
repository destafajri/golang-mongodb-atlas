package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"bufio"
	"os"
	"strings"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

func main() {
	//credential input
	fmt.Print("Username: ")
	scannerUser := bufio.NewScanner(os.Stdin)
   	scannerUser.Scan()
    username := scannerUser.Text()
    fmt.Print("Password: ")
	scannerPass := bufio.NewScanner(os.Stdin)
   	scannerPass.Scan()
    pass := scannerPass.Text()
	fmt.Print(username)
	credentialArray := []string {"mongodb+srv://", username, ":", pass, "@cluster0.u7lz8u5.mongodb.net/?retryWrites=true&w=majority"}
	credential := strings.Join(credentialArray,"")
	fmt.Println(credential)

	//mongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI(credential))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	database := client.Database("Learn")
	podcastsCollection := database.Collection("Learn")

	podcast := Podcast{
		Title:  "The Polyglot Developer",
		Author: "Nic Raboy",
		Tags:   []string{"development", "programming", "coding"},
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)

}