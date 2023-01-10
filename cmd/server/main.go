package main

import (
	"context"
	"fmt"

	db "github.com/raphaelmb/go-comments-api/internal/database"
)

// Responsible for the instantiation and startup of the app
func Run() error {
	fmt.Println("Startup...")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully connected and pinged database")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
