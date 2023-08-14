package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	"github.com/jutionck/simple-golang-sqlc/tutorial"
	_ "github.com/lib/pq"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}

	queries := tutorial.New(db)

	// list all users
	users, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(users)

	// create an user
	insertedUser, err := queries.CreateUSer(ctx, tutorial.CreateUSerParams{
		ID:       "US001",
		Username: sql.NullString{String: "admin"},
		Password: sql.NullString{String: "password", Valid: true},
		IsActive: sql.NullBool{Bool: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedUser)

	// get the user we just inserted
	fetcheUser, err := queries.GetUser(ctx, insertedUser.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedUser, fetcheUser))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
