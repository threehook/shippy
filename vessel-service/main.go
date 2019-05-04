package main

import (
	"context"
	"fmt"
	"time"

	"os"

	"github.com/micro/go-micro"
	pb "github.com/threehook/shippy/vessel-service/proto/vessel"
)

//const (
//	defaultHost = "localhost:27017"
const mongoURI = "mongodb://mongo:27017"

//)

func createDummyData(repo Repository) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	defer repo.Close(ctx)
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(ctx, v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = mongoURI
	}

	client := CreateDb(host)
	repo := &VesselRepository{client}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{client})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
