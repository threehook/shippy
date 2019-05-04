package main

import (
	"context"
	pb "github.com/threehook/shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(context.Context, *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close(ctx context.Context)
}

type ConsignmentRepository struct {
	client *mongo.Client
}

// Create a new consignment
func (repo *ConsignmentRepository) Create(ctx context.Context, consignment *pb.Consignment) error {
	_, err := repo.collection().InsertOne(ctx, consignment)
	return err
}

// GetAll consignments
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.collection().Find(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result pb.Consignment
		//var result pb.Consignment
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		consignments = append(consignments, &result)
	}
	return consignments, err
}

// Close closes the database session after each query has ran.
// Mgo creates a 'master' session on start-up, it's then good practice
// to clone a new session for each request that's made. This means that
// each request has its own database session. This is safer and more efficient,
// as under the hood each session has its own database socket and error handling.
// Using one main database socket means requests having to wait for that session.
// I.e this approach avoids locking and allows for requests to be processed concurrently. Nice!
// But... it does mean we need to ensure each session is closed on completion. Otherwise
// you'll likely build up loads of dud connections and hit a connection limit. Not nice!
func (repo *ConsignmentRepository) Close(ctx context.Context) {
	repo.client.Disconnect(ctx)
}

func (repo *ConsignmentRepository) collection() *mongo.Collection {
	return repo.client.Database(dbName).Collection(consignmentCollection)
}
