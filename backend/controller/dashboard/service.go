package dashboard

import (
	"github.com/julyusmanurung/Kredit/models"
	"log"
	"net/http"
)

type Service interface {
	GetCustomerData() ([]models.CustomerDataTab, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetCustomerData() ([]models.CustomerDataTab, int, error) {
	data, err := s.repo.GetCustomerData()
	if err != nil {
		log.Println("Internal server error: ", err)
		return nil, http.StatusInternalServerError, err
	}
	return data, http.StatusOK, nil
}
