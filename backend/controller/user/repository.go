package user

import (
	"github.com/julyusmanurung/Kredit/models"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	GetUserDetails(userID string) (models.User, error)
	UpdatePassword(userID string, req DataRequestUpdatePassword) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUserDetails(userID string) (models.User, error) {
	var user models.User

	res := r.db.Where("user_id = ?", userID).First(&user)

	if res.Error != nil {
		log.Println("get recent credit applicant error: ", res.Error)
		return models.User{}, res.Error
	}

	return user, res.Error
}

func (r *repository) UpdatePassword(userID string, req DataRequestUpdatePassword) (string, error) {
	var user models.User
	newPassword := models.User{Password: req.NewPassword}

	res := r.db.Where("user_id =?", userID).Find(&user)
	if res.Error != nil {
		log.Println(res.Error)
		return "terjadi error: ", res.Error
	}

	currentPassword := user.Password
	if req.OldPassword != currentPassword {
		return "password lama salah", nil
	} else {
		res2 := r.db.Model(models.User{}).Where("user_id = ?", userID).Updates(&newPassword)
		if res2.Error != nil {
			log.Println("update password failed: ", res.Error)
			return "password gagal diubah: ", res.Error
		}
	}
	return "password berhasil diubah", nil
}
