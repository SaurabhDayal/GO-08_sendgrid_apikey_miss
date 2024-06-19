package main

import (
	"GO-08/database"
	"GO-08/handlers"
	"GO-08/providers/emailprovider"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	handlers.EmailProvider = emailprovider.NewSendGridEmailProvider(os.Getenv("SENDGRID_KEY"))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Connect to MongoDB
	database.ConnectDatabase()

	r.Post("/contact-user", handlers.ContactUser)

	// Start server
	server := &http.Server{
		Addr:    database.Port,
		Handler: r,
	}

	go func() {
		fmt.Printf("Server is listening on port %s\n", database.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v\n", database.Port, err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	fmt.Println("Closing MongoDB connection...")
	if err := database.DB.Disconnect(database.MongoCtx); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
	}

	fmt.Println("Server exited")
}
