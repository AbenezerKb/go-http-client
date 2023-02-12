package client

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) GetPokemonByName(
	ctx context.Context,
	pokemon string,
) (*Custome, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/api/v2/pokemon/"+pokemon,
		nil)
	if err != nil {
		return nil, PokeManError{
			Message: err.Error(),
			Status:  -1,
		}
	}
	req.Header.Add("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, PokeManError{
			Message: err.Error(),
			Status:  -1,
		}
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, PokeManError{
			Message: "non 200 status code",
			Status:  resp.StatusCode,
		}
	}
	var custom Custome
	err = json.NewDecoder(resp.Body).Decode(&custom)
	if err != nil {
		return nil, PokeManError{
			Message: err.Error(),
			Status:  resp.StatusCode,
		}
	}
	return &custom, nil
}
