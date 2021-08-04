package domain

import (
	"errors"
)

type Summary struct {
	value string
}

func (Summary) New(value string) (*Summary, error) {
	if len(value) > 2500 {
		return nil, errors.New("the summary cannot be longer than 2500 characters.")
	}

	return &Summary{
		value: value,
	}, nil
}
