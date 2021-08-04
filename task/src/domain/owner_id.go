package domain

import "errors"

type OwnerId struct {
	value int
}

func (OwnerId) New(value int) (*OwnerId, error) {
	if value <= 0 {
		return nil, errors.New("Invalid Owner ID.")
	}

	return &OwnerId{
		value: value,
	}, nil
}
