package tests

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("rambo"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		body := `{"slug": "/", "author": "John", "body": "test"}`
		url := "http://localhost:8080/api/v1/comment"
		token := "Bearer " + createToken()

		client := resty.New()
		resp, err := client.R().SetHeader("Authorization", token).SetBody(body).Post(url)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("cannot post comment without JWT", func(t *testing.T) {
		body := `{"slug": "/", "author": "John", "body": "test"}`
		url := "http://localhost:8080/api/v1/comment"

		client := resty.New()
		resp, err := client.R().SetBody(body).Post(url)
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})

}
