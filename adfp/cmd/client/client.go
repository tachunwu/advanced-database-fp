package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "adfp/pkg/pb/adfp"

	"google.golang.org/grpc"
)

func RandomString(n int) string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	// Connect to grpc server
	conn, err := grpc.Dial("localhost:30000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// New client instance
	client := pb.NewADFPServiceClient(conn)

	// Test var
	demoUserName := "User_" + RandomString(10)
	demoPlaceName := "Place_" + RandomString(10)

	// Create a demo user
	createDemoUser(client, &pb.CreateUserRequest{
		User: &pb.User{
			Name:    demoUserName,
			Balance: 1000,
			Location: &pb.LatLng{
				Latitude:  25.033964,
				Longitude: 121.564468,
			},
		},
	})
	// Create a demo place
	createDemoPlace(client, &pb.CreatePlaceRequest{
		Place: &pb.Place{
			Name:     demoPlaceName,
			Category: "MRT_Station",
			Location: &pb.LatLng{
				Latitude:  25.033964,
				Longitude: 121.564468,
			},
		},
	})
	// Create a demo comment with pay transaction
	createDemoComment(client, &pb.CreateCommentRequest{
		Comment: &pb.Comment{
			Name:    "",
			Context: "Nice place!",
			IsPay:   true,
			User: &pb.User{
				Name: demoUserName,
			},
			Place: &pb.Place{
				Name: demoPlaceName,
			},
		},
	})

}

func createDemoUser(client pb.ADFPServiceClient, request *pb.CreateUserRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user, err := client.CreateUser(ctx, request)
	if err != nil {
		log.Println(err)
	}
	log.Println(user)
}
func createDemoPlace(client pb.ADFPServiceClient, request *pb.CreatePlaceRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	place, err := client.CreatePlace(ctx, request)
	if err != nil {
		log.Println(err)
	}
	log.Println(place)
}
func createDemoComment(client pb.ADFPServiceClient, request *pb.CreateCommentRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	comment, err := client.CreateComment(ctx, request)
	if err != nil {
		log.Println(err)
	}
	log.Println(comment)
}
