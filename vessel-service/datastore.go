package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func CreateDb(host string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// CreateSession creates the main session to our mongodb instance
//func CreateSession(host string) (*mgo.Session, error) {
//	session, err := mgo.Dial(host)
//	if err != nil {
//		return nil, err
//	}
//
//	session.SetMode(mgo.Monotonic, true)
//
//	return session, nil
//}

//func WriteJSONResult(w http.ResponseWriter, v interface{}) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(v)
//}
//
//func ConnectDb() *mongo.Client {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	clientOptions := options.Client().ApplyURI(mongoURI)
//	client, err := mongo.Connect(ctx, clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return client
//}
