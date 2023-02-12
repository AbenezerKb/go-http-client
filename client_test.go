package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Run("the request can hit the API and return", func(*testing.T) {
		myClient := NewClient()
		poke, err := myClient.GetPokemonByName(context.Background(), "clefairy")
		assert.NoError(t, err)
		assert.Equal(t, "clefairy", poke.Name)
	})

	t.Run("return error", func(*testing.T) {
		myClient := NewClient()
		_, err := myClient.GetPokemonByName(context.Background(), "non-pokemon")
		assert.Error(t, err)
		assert.Error(t, PokeManError{Message: "non 200 status code", Status: 404} ,err)
	})
	t.Run("testing the return with URL ", func(*testing.T) {
		myClient := NewClient(
			WithAPIURL("test-URL"),
		)
		assert.Equal(t, "test-URL", myClient.apiURL)
	})

	t.Run("testing the return with HTTP client ", func(*testing.T) {
		myClient := NewClient(
			WithAPIURL("test-URL"),
			WithHTTPClient(&http.Client{
				Timeout: 10 * time.Second,
			}),
		)
		assert.Equal(t, 10*time.Second, myClient.httpClient.Timeout)
	})

	t.Run("testing the local running server ", func(*testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"name":"clefairy", "height":30}`)
			}),
		)
		defer ts.Close()

		myClient := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := myClient.GetPokemonByName(context.Background(), "clefairy")
		assert.NoError(t, err)
		assert.Equal(t, 30, poke.Height)
	})

	t.Run("testing internal server error ", func(*testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}),
		)
		defer ts.Close()

		myClient := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := myClient.GetPokemonByName(context.Background(), "clefairy")
		assert.Error(t, err)
		assert.Error(t, PokeManError{Message: "non 200 status code", Status: 500} ,err)
		assert.Equal(t, (*Custome)(nil), poke)
	})

}
