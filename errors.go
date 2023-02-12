package client

import (
	"errors"
	"fmt"
)

var (
	ErrFetchingResources = errors.New("error ocurred while fetching resource")
)

type PokeManError struct {
	Message string
	Status  int
}

func (p PokeManError) Error() string {
	return fmt.Sprintf("failed to fetch pokemon: %s with status code %d", p.Message, p.Status)
}
