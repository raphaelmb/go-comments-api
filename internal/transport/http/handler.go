package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Router  *chi.Mux
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = chi.NewRouter()
	h.Router.Use(JSONMiddleware, LoggingMiddleware, TimeoutMiddleware)
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.Route("/api/v1/comment", func(r chi.Router) {
		r.Post("/", JWTAuth(h.PostComment))
		r.Get("/{id}", h.GetComment)
		r.Put("/{id}", JWTAuth(h.UpdateComment))
		r.Delete("/{id}", JWTAuth(h.DeleteComment))
	})
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")

	return nil
}
