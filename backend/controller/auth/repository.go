package auth

import (
	"github.com/julyusmanurung/Kredit/models"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	Login(req DataRequest) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(req DataRequest) (models.User, error) {
	var user models.User

	res := r.db.Where("user_id = ?", req.Nik).Find(&user)
	//fmt.Println("User: %+v\n", user)

	if res.Error != nil {
		log.Println(res.Error)
		return user, res.Error
	}

	if user.UserId == "" {
		return models.User{}, res.Error
	}

	err := user.Password == req.Password
	if !err {
		return models.User{}, res.Error
	}

	return user, nil
}
