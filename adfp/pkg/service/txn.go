package service

import (
	pb "adfp/pkg/pb/adfp"
	"adfp/pkg/storage"
	"adfp/pkg/stream"
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

type TxnService struct {
	Stream  *nats.EncodedConn
	Storage *storage.TxnStorage
}

func NewTxnService() *TxnService {
	svc := &TxnService{
		Stream:  stream.NewStream(),
		Storage: storage.NewTxnStorage(),
	}

	svc.Stream.Subscribe("adfp.comment.create", svc.CreateComment)
	svc.Stream.Subscribe("adfp.user.create", svc.CreateUser)
	svc.Stream.Subscribe("adfp.place.create", svc.CreatePlace)

	return svc
}

func (s *TxnService) CreateUser(req *pb.CreateUserRequest) {
	ctx := context.Background()
	err := s.Storage.CreateUser(
		ctx,
		req.User.Name,
		req.User.Balance,
		req.User.Location.Latitude,
		req.User.Location.Longitude,
	)

	if err != nil {
		log.Println(err)
	}
}

func (s *TxnService) CreatePlace(req *pb.CreatePlaceRequest) {
	ctx := context.Background()
	err := s.Storage.CreatePlace(
		ctx,
		req.Place.Name,
		req.Place.Category,
		req.Place.Location.Latitude,
		req.Place.Location.Longitude,
	)

	if err != nil {
		log.Println(err)
	}
}

func (s *TxnService) CreateComment(req *pb.CreateCommentRequest) {
	ctx := context.Background()
	err := s.Storage.CreateComment(
		ctx,
		req.Comment.User.Name,
		req.Comment.Context,
		req.Comment.Place.Name,
		req.Comment.IsPay,
	)

	if err != nil {
		log.Println(err)
	}
}
