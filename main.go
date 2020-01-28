package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//User type
/* type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
} */

type FootballPlayer struct {
	ID      string `bson:"_id"`
	Name 	string `bson:"name"`
    Age  	int	   `bson:"age"`
    Club 	string
}

/* type FootballPlayer struct {
    Name string
    Age  int
    Club string
} */
 


func main() {
	

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("local").Collection("hellomongo")
	
	
	/*player1 := FootballPlayer{"Christiano Ronaldo", 34, "Juventus"}
	player2 := FootballPlayer{"Lionel Messi", 33, "Barcelona"}
	player3 := FootballPlayer{"David De Gea", 29, "Manchester United"}
	
	 insertResult, err := collection.InsertOne(context.TODO(), player1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	trainers := []interface{}{player1, player2, player3}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs) */

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(10)

	// Here's an array in which you can store the decoded documents
	var results []*FootballPlayer

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), (bson.M{"age":34}))
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem FootballPlayer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	for _, element := range results {
		fmt.Println(*element)
	}
	
}