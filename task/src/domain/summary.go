package domain

import (
	"errors"
	"fmt"
)

const (
	summaryMinCharacters      = 4
	summaryMinCharactersError = "The summary cannot be at least %d characters long."
	summaryMaxCharacters      = 2500
	summaryMaxCharactersError = "The summary cannot be at least %d characters long."
)

type Summary struct {
	value string
}

func (Summary) New(value string) (*Summary, error) {
	if len(value) < summaryMinCharacters {
		return nil, errors.New(
			fmt.Sprintf(
				summaryMinCharactersError,
				summaryMinCharacters,
			),
		)
	}

	if len(value) > summaryMaxCharacters {
		return nil, errors.New(
			fmt.Sprintf(
				summaryMaxCharactersError,
				summaryMaxCharacters,
			),
		)
	}

	return &Summary{
		value: value,
	}, nil
}
