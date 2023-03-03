package dashboard

import (
	"github.com/julyusmanurung/Kredit/models"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	GetCustomerData() ([]models.CustomerDataTab, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCustomerData() ([]models.CustomerDataTab, error) {
	var data []models.CustomerDataTab

	res := r.db.Find(&data)
	if res.Error != nil {
		log.Println("count customer data error: ", res.Error)
		return nil, res.Error
	}
	return data, nil
}
