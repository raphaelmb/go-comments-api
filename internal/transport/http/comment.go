package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/raphaelmb/go-comments-api/internal/comment"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, id string) (comment.Comment, error)
	UpdateComment(ctx context.Context, id string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, id string) error
}

type Response struct {
	Message string
}

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func convertPostCommentRequestToComment(c PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(cmt)
	if err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
		return
	}

	convertedCmt := convertPostCommentRequestToComment(cmt)
	postedCmt, err := h.Service.PostComment(r.Context(), convertedCmt)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(postedCmt); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}

}
