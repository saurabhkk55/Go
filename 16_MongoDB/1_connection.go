package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connectMongoDB establishes a connection to the default/local MongoDB server.
func connectMongoDB() (*mongo.Client, error) {
	// Set MongoDB connection options, specifying the server URI.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to the MongoDB server using the specified options.
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Ping the MongoDB server to ensure a successful connection.
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Connection successful, log a message indicating the successful connection.
	log.Println("Connected to MongoDB!")

	// Return the MongoDB client for further use.
	return client, nil
}

// createDatabase creates a new database in MongoDB and prints a message once it's created.
func createDatabase(client *mongo.Client, dbName string) {
	// Create a new database using the client.Database method.
	client.Database(dbName)

	// Log a message indicating the successful creation of the database.
	log.Printf("Database '%s' created successfully!", dbName)
}

// createCollection creates a new collection inside the specified database and prints a message once it's created.
func createCollection(client *mongo.Client, dbName, collectionName string) {
	// Access the specified database.
	db := client.Database(dbName)

	// Create a new collection using the Database.CreateCollection method.
	err := db.CreateCollection(context.TODO(), collectionName)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Log a message indicating the successful creation of the collection.
	log.Printf("Collection '%s' created successfully in database '%s'!", collectionName, dbName)
}

// insertDocument inserts a document into the specified collection in MongoDB.
func insertDocument(client *mongo.Client, dbName, collectionName string) {
	// Access the specified database and collection.
	collection := client.Database(dbName).Collection(collectionName)

	var user_input string

	for {
		fmt.Println("press 'Y' or 'y' to insert data, otherwise press any key to exit.")
		fmt.Scan(&user_input)
		if user_input == "Y" || user_input == "y" {
			document := getUserInput()
			result, err := collection.InsertOne(context.TODO(), document)
			if err != nil {
				log.Fatal(err)
				return
			}
			// Log a message indicating the successful insertion of the document.
			log.Printf("Document inserted successfully! Document ID: %v", result.InsertedID)
		} else {
			break
		}
	}
}

// Input from user that needs to be inserted
func getUserInput() bson.D {
	var user_name, user_age, user_gen string

	// Get user input for key and value
	fmt.Print("Enter name: ")
	fmt.Scan(&user_name)

	fmt.Print("Enter age: ")
	fmt.Scan(&user_age)

	fmt.Print("Enter gender: ")
	fmt.Scan(&user_gen)

	// Create a BSON document with the user input
	document := bson.D{
		primitive.E{Key: "Name", Value: user_name},
		primitive.E{Key: "Age", Value: user_age},
		primitive.E{Key: "Gender", Value: user_gen},
	}

	return document
}

// fetchDocument retrieves a document from the specified collection based on a field and its value.
func fetchDocument(client *mongo.Client, dbName, collectionName string) {
	// Access the specified database and collection.
	collection := client.Database(dbName).Collection(collectionName)

	var user_name string
	fmt.Print("Enter name to get its correspondng document: ")
	fmt.Scan(&user_name)

	// Define a filter based on the field and its value.
	filter := bson.D{{Key: "Name", Value: user_name}}

	// Find the document using the Collection.FindOne method with the specified filter.
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	var k int
	if err != nil {
		fmt.Println("ERROR: ", err)
		k = 0
		// log.Fatal(err)
		// return
	} else {
		if k == 0 {
			for key, value := range result {
				fmt.Printf("%s: %v\n", key, value)
			}
		}
	}

	// Print the retrieved document.
	// fmt.Println(result)
}

// deleteDocument deletes a document from the specified collection based on a field and its value.
func deleteDocument(client *mongo.Client, dbName, collectionName string) {
	// Access the specified database and collection.
	collection := client.Database(dbName).Collection(collectionName)

	var fieldName, value string

	fieldName = "Name"

	fmt.Print("Enter name to delete its corresponding document: ")
	fmt.Scan(&value)

	// Define a filter based on the field and its value.
	filter := bson.D{{Key: fieldName, Value: value}}

	// Delete the document using the Collection.DeleteOne method with the specified filter.
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Print the number of documents deleted.
	fmt.Printf("Deleted %v document(s) with the specified filter.\n", result.DeletedCount)
}

func main() {
	// Attempt to connect to MongoDB.
	client, err := connectMongoDB()
	if err != nil {
		// If connection fails, log the error and exit the program.
		log.Fatal(err)
		return
	}
	// Defer closing the MongoDB connection until the end of the program.
	defer client.Disconnect(context.TODO())

	// Specify the name of the database you want to create.
	dbName := "db_san"
	// Create a new database.
	createDatabase(client, dbName)

	// Specify the name of the collection you want to create.
	collectionName := "col_san"

	// Create a new collection inside the database.
	createCollection(client, dbName, collectionName)

	// Insert document
	insertDocument(client, dbName, collectionName)

	// Fetch and print the document based on the specified field and value.
	fetchDocument(client, dbName, collectionName)

	// Delete the document based on the specified field and value.
	deleteDocument(client, dbName, collectionName)
}
