package user

import (
	"log"
	"net/http"
)

type Service interface {
	GetUserDetails(userID string) (DataResponse, error)
	UpdatePassword(userID string, req DataRequestUpdatePassword) (string, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetUserDetails(userID string) (DataResponse, error) {
	user, err := s.repo.GetUserDetails(userID)
	if err != nil {
		log.Println("Internal server error: ", err)
		return DataResponse{}, err
	}

	res := DataResponse{
		UserId:  user.UserId,
		Name:    user.Name,
		Email:   user.Email,
		Level:   user.Level,
		Jabatan: user.Jabatan,
	}

	return res, nil
}

func (s *service) UpdatePassword(userID string, req DataRequestUpdatePassword) (string, int, error) {
	res, err := s.repo.UpdatePassword(userID, req)
	if err != nil {
		log.Println(err, "gabisa ganti password")
		return "gabisa ganti password", http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}
