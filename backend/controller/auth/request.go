package auth

import (
	"errors"
)

type DataRequest struct {
	Nik      string `json:"nik"`
	Password string `json:"password"`
}

func (r *DataRequest) Validation() error {
	if r.Nik == "" {
		return errors.New("invalid")
	}

	if r.Password == "" {
		return errors.New("invalid")
	}

	return nil
}
