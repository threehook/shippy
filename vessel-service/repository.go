package main

import (
	"context"
	pb "github.com/threehook/shippy/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(ctx context.Context, vessel *pb.Vessel) error
	Close(ctx context.Context)
}

type VesselRepository struct {
	client *mongo.Client
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	// Here we define a more complex query than our consignment-service's
	// GetAll function. Here we're asking for a vessel who's max weight and
	// capacity are greater than and equal to the given capacity and weight.
	// We're also using the `One` function here as that's all we want.
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := repo.collection().FindOne(ctx, bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).Decode(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *VesselRepository) Create(ctx context.Context, vessel *pb.Vessel) error {
	_, err := repo.collection().InsertOne(ctx, vessel)
	return err
}

func (repo *VesselRepository) Close(ctx context.Context) {
	repo.client.Disconnect(ctx)
}

func (repo *VesselRepository) collection() *mongo.Collection {
	return repo.client.Database(dbName).Collection(vesselCollection)
}
