package stream

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NewStream() *nats.EncodedConn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	return ec
}
