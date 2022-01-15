package storage

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
)

type GraphStorage struct {
	Client *dgo.Dgraph
}

// Schema
type loc struct {
	Type   string    `json:"type,omitempty"`
	Coords []float64 `json:"coordinates,omitempty"`
}

type Place struct {
	Id       string    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Comments []Comment `json:"comment,omitempty"`
	Category string    `json:"category,omitempty"`
	Location loc       `json:"loc,omitempty"`
	DType    []string  `json:"dgraph.type,omitempty"`
}

type Comment struct {
	Id      string   `json:"id,omitempty"`
	User    User     `json:"user,omitempty"`
	Place   Place    `json:"place,omitempty"`
	Star    int      `json:"star,omitempty"`
	IsPay   bool     `json:"is_pay,omitempty"`
	Context string   `json:"context,omitempty"`
	DType   []string `json:"dgraph.type,omitempty"`
}

type User struct {
	Id                string   `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Collection_places []Place  `json:"collection_places,omitempty"`
	Location          loc      `json:"loc,omitempty"`
	DType             []string `json:"dgraph.type,omitempty"`
}

func NewGraphStorage() *GraphStorage {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return &GraphStorage{
		Client: dgo.NewDgraphClient(
			api.NewDgraphClient(d),
		),
	}
}

func (s *GraphStorage) CreateUser(ctx context.Context, username string, lat float64, lng float64) error {

	log.Println("Creating user by dgraph...")

	user := User{
		DType: []string{"User"},
		Name:  username,
		Location: loc{
			Type:   "Point",
			Coords: []float64{lat, lng},
		},
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	userb, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return err
	}

	mu.SetJson = userb
	response, err := s.Client.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println(response)
		return nil
	}
}

func (s *GraphStorage) CreateComment(ctx context.Context, username string, context string, place string, isPay bool) error {

	log.Println("Creating comment by dgraph...")

	comment := Comment{
		DType:   []string{"Comment"},
		Context: context,
		User: User{
			Name: username,
		},
		Place: Place{
			Name: place,
		},
		IsPay: isPay,
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	commentb, err := json.Marshal(comment)
	if err != nil {
		log.Println(err)
		return err
	}

	mu.SetJson = commentb
	response, err := s.Client.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println(response)
		return nil
	}
}

func (s *GraphStorage) CreatePlace(ctx context.Context, name string, category string, lat float64, lng float64) error {
	log.Println("Creating place by dgraph...")

	place := Place{
		DType:    []string{"Place"},
		Name:     name,
		Category: category,
		Location: loc{
			Type:   "Point",
			Coords: []float64{lat, lng},
		},
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	placeb, err := json.Marshal(place)
	if err != nil {
		log.Println(err)
		return err
	}

	mu.SetJson = placeb
	response, err := s.Client.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println(response)
		return nil
	}
}
