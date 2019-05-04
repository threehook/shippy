package main

import (
	pb "github.com/threehook/shippy/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

// Our gRPC service handler
type service struct {
	client *mongo.Client
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.client}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close(ctx)

	// Find the next available vessel
	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close(ctx)
	if err := repo.Create(ctx, req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
