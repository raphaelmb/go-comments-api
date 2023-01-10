package main

import (
	"context"
	"fmt"

	"github.com/raphaelmb/go-comments-api/internal/comment"
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(context.Background(), "35bc94f8-5624-49de-809a-289c83768259"))

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
