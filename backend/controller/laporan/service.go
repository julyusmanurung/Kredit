package laporan

import (
	"log"
	"net/http"
)

type Service interface {
	GetLaporanData(data DataRequest) ([]CustomerJoinLoan, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetLaporanData(data DataRequest) ([]CustomerJoinLoan, int, error) {
	laporan, err := s.repo.GetLaporanData(data.Branch, data.ChannelingCompany, data.StartDate, data.EndDate)
	if err != nil {
		log.Println("Internal server error: ", err)
		return nil, http.StatusInternalServerError, err
	}

	return laporan, http.StatusOK, nil
}
