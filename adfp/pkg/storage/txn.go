package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type TxnStorage struct {
	Pool *pgxpool.Pool
}

func NewTxnStorage() *TxnStorage {

	config, err := pgxpool.ParseConfig("postgres://root@localhost:26257/adfp?sslmode=disable")
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	// Create a connection pool to the "bank" database.
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	// Success log
	log.Println("Hey! You successfully connected to your local CockroachDB cluster.")

	return &TxnStorage{
		Pool: pool,
	}
}

func (s *TxnStorage) CreateUser(ctx context.Context, username string, balance int64, lat float64, lng float64) error {

	log.Println("Creating user by cockraochdb...")

	if _, err := s.Pool.Exec(ctx,
		"INSERT INTO users (name, balance, lat, lng) VALUES ($1, $2, $3, $4)",
		username, balance, lat, lng); err != nil {
		return err
	}

	return nil
}

func (s *TxnStorage) CreateComment(ctx context.Context, username string, context string, place string, isPay bool) error {

	log.Println("Creating comment by cockraochdb...")

	var placeId string
	var userId string

	if err := s.Pool.QueryRow(ctx,
		"SELECT id FROM places WHERE name = $1",
		place).Scan(&placeId); err != nil {
		return err
	}

	if err := s.Pool.QueryRow(ctx,
		"SELECT id FROM users WHERE name = $1",
		username).Scan(&userId); err != nil {
		return err
	}

	if isPay {

		if _, err := s.Pool.Exec(ctx,
			"UPDATE users SET balance = balance - 10 WHERE name = $1",
			username); err != nil {
			return err
		}

		if _, err := s.Pool.Exec(ctx,
			"INSERT INTO comments (user_id, place_id, context, is_pay) VALUES ($1, $2, $3, $4)",
			userId, placeId, context, true); err != nil {
			return err
		}

	} else {
		if _, err := s.Pool.Exec(ctx,
			"INSERT INTO comments (user_id, place_id, context, is_pay) VALUES ($1, $2, $3, $4)",
			userId, placeId, context, false); err != nil {
			return err
		}
	}

	return nil
}

func (s *TxnStorage) CreatePlace(ctx context.Context, name string, category string, lat float64, lng float64) error {

	log.Println("Creating place by cockraochdb...")

	if _, err := s.Pool.Exec(ctx,
		"INSERT INTO places (name, category, lat, lng) VALUES ($1, $2, $3, $4)",
		name, category, lat, lng); err != nil {
		return err
	}
	return nil
}
