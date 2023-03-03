package auth

import "net/http"

type Service interface {
	Login(req DataRequest) (DataResponse, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Login(req DataRequest) (DataResponse, int, error) {
	if err := req.Validation(); err != nil {
		return DataResponse{}, http.StatusInternalServerError, err
	}

	user, err := s.repo.Login(req)
	if err != nil {
		return DataResponse{}, http.StatusInternalServerError, err
	}

	res := DataResponse{
		Nik: user.UserId,
	}

	return res, http.StatusOK, nil
}
