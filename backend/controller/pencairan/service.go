package pencairan

import (
	"github.com/julyusmanurung/Kredit/models"
	"log"
	"net/http"
)

type Service interface {
	GetRecentCreditApplicant() ([]models.StagingCustomer, int, error)
	GetAllBranch() ([]models.BranchTab, int, error)
	GetAllCompany() ([]models.MstCompanyTab, int, error)
	GetAllApprovalStatusNine() ([]CustomerJoinLoan, int, error)
	GetAllApprovalStatusNineFilter(data DataRequest) ([]CustomerJoinLoan, int, error)
	UpdateApprovalStatus(checkedList DataPostPPK) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetRecentCreditApplicant() ([]models.StagingCustomer, int, error) {
	user, err := s.repo.GetRecentCreditApplicant()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func (s *service) GetAllBranch() ([]models.BranchTab, int, error) {
	branch, err := s.repo.GetAllBranch()
	if err != nil {
		log.Println("Internal server error: ", err)
		return nil, http.StatusInternalServerError, err
	}

	return branch, http.StatusOK, nil
}

func (s *service) GetAllCompany() ([]models.MstCompanyTab, int, error) {
	company, err := s.repo.GetAllCompany()
	if err != nil {
		log.Println("Internal server error: ", err)
		return nil, http.StatusInternalServerError, err
	}

	return company, http.StatusOK, nil
}

func (s *service) GetAllApprovalStatusNine() ([]CustomerJoinLoan, int, error) {
	approval, err := s.repo.GetAllApprovalStatusNine()
	if err != nil {
		log.Println("Internal server error: ", err)
		return nil, http.StatusInternalServerError, err
	}

	return approval, http.StatusOK, nil
}

func (s *service) GetAllApprovalStatusNineFilter(data DataRequest) ([]CustomerJoinLoan, int, error) {
	approval, err := s.repo.GetAllApprovalStatusNineFilter(data.Branch, data.ChannelingCompany, data.StartDate, data.EndDate)
	if err != nil {
		log.Println("Internal server error: ", err)
		return nil, http.StatusInternalServerError, err
	}

	return approval, http.StatusOK, nil
}

func (s *service) UpdateApprovalStatus(checkedList DataPostPPK) error {
	err := s.repo.UpdateApprovalStatus(checkedList)
	if err != nil {
		log.Println("Internal server error: ", err)
		return err
	}

	return nil
}
