package service

import (
	pb "adfp/pkg/pb/adfp"
	"adfp/pkg/stream"
	"context"
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

type ADFPService struct {
	Stream *nats.EncodedConn
	pb.UnimplementedADFPServiceServer
}

func NewADFPService() *ADFPService {
	lis, err := net.Listen("tcp", "localhost:30000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	s := &ADFPService{
		Stream: stream.NewStream(),
	}
	pb.RegisterADFPServiceServer(grpcServer, s)
	grpcServer.Serve(lis)
	return s
}

func (s *ADFPService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {

	// Log
	log.Println(req)

	// Publish
	if err := s.Stream.Publish("adfp.user.create", req); err != nil {
		log.Println(err)
	}

	return &pb.User{}, nil
}

func (s *ADFPService) CreatePlace(ctx context.Context, req *pb.CreatePlaceRequest) (*pb.Place, error) {

	// Log
	log.Println(req)

	// Publish
	if err := s.Stream.Publish("adfp.place.create", req); err != nil {
		log.Println(err)
	}
	return &pb.Place{}, nil
}

func (s *ADFPService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.Comment, error) {

	// Log
	log.Println(req)

	// Publish
	if err := s.Stream.Publish("adfp.comment.create", req); err != nil {
		log.Println(err)
	}
	return &pb.Comment{}, nil
}
