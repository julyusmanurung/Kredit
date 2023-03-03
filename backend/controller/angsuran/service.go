package angsuran

import (
	"github.com/julyusmanurung/Kredit/models"
	"log"
)

type Service interface {
	GetInstallmentScale() ([]models.CustomerDataTab, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetInstallmentScale() ([]models.CustomerDataTab, error) {
	user, err := s.repo.GetInstallmentScale()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, err
	}

	return user, nil
}
