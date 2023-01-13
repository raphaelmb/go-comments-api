package main

import (
	"fmt"

	"github.com/raphaelmb/go-comments-api/internal/comment"
	"github.com/raphaelmb/go-comments-api/internal/db"
	transportHttp "github.com/raphaelmb/go-comments-api/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
