package mongodb

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection Guide: https://www.mongodb.com/docs/drivers/go/current/fundamentals/connections/connection-guide/

type Operator struct {
	Kind       string
	Uri        string
	database   string
	collection string
	Client     *mongo.Client
}

// NewOperator return
func NewOperator(connStr string) *Operator {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().SetServerAPIOptions(serverAPI).ApplyURI(connStr)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("***** [DATABASE][FAIL] ***** Failed to connect to MongoDB server:: %s", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("***** [DATABASE][FAIL] ***** Failed to Ping to MongoDB:: %s", err)

	}
	log.Println("***** [DATABASE:%s] ***** Create connection to MongoDB!")

	// defer func(client *mongo.Client, ctx context.Context) {
	// 	err := client.Disconnect(ctx)
	// 	if err != nil {
	// 		log.Infof("Failed to disconnect from MongoDB server: %+v", err)
	// 	}
	// }(client, ctx)

	return &Operator{
		Kind:   "MongoDB",
		Uri:    connStr,
		Client: client,
	}
}

func (o *Operator) FetchAll()            {}
func (o *Operator) FetchById(id string)  {}
func (o *Operator) UpdateById(id string) {}
func (o *Operator) DeleteById(id string) {}

/*
"Functional Options" pattern.
The Functional Options pattern is a design pattern in Go where you pass in functions that alter the state of a type.
These functions are often called "option functions". They provide a way to cleanly design APIs and offer a more flexible
and readable way to interact with a function or type.

The Most Efficient Struct Configuration Pattern For Golang: https://www.youtube.com/watch?v=MDy7JQN5MN4
*/
type optsFunc func(*Operator)

// Below are option functions which can specify Connection String Options
// Ref: https://www.mongodb.com/docs/manual/reference/connection-string/#connection-string-options
