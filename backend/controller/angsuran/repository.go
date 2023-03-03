package angsuran

import (
	"fmt"
	"github.com/julyusmanurung/Kredit/models"
	"gorm.io/gorm"
	"log"
	"math"
	"strconv"
	"time"
)

type Repository interface {
	GetInstallmentScale() ([]models.CustomerDataTab, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetInstallmentScale() ([]models.CustomerDataTab, error) {
	var customerDataTab []models.CustomerDataTab
	var loanDataTab models.LoanDataTab

	res := r.db.Where("approval_status = ?", "9").Find(&customerDataTab)
	if res.Error != nil {
		log.Println("get installment scale error: ", res.Error)
		return nil, res.Error
	}

	for _, item := range customerDataTab {
		res := r.db.Where("custcode = ?", item.Custcode).First(&loanDataTab)
		if res.Error != nil {
			log.Println("Get Data error : ", res.Error)
			return nil, res.Error
		}

		loanPeriod, err := strconv.ParseInt(loanDataTab.LoanPeriod, 10, 8)
		if err != nil {
			fmt.Println(err)
		}

		osBalance := loanDataTab.LoanAmount
		monthlyPayment := loanDataTab.MonthlyPayment
		installmentScale := make([]models.SkalaRentalTab, loanPeriod+1)

		for i := range installmentScale {
			osBalance = r.insertCustomerToSkalaRentalTab(item.Custcode, osBalance, monthlyPayment, loanDataTab, i)
		}
	}

	return customerDataTab, res.Error
}

func (r *repository) insertCustomerToSkalaRentalTab(
	customerCode string,
	osBalance float64,
	monthlyPayment float64,
	loanDataTab models.LoanDataTab,
	index int) float64 {
	installmentScale := models.SkalaRentalTab{}
	if index == 0 {
		installmentScale = models.SkalaRentalTab{
			Custcode:   customerCode,
			Counter:    int8(index),
			Osbalance:  osBalance,
			EndBalance: osBalance,
			DueDate:    time.Now(),
			EffRate:    float64(loanDataTab.InterestEffective),
			Rental:     monthlyPayment,
			Principle:  0,
			Interest:   0,
			Inputdate:  time.Now(),
		}

		r.db.Create(&installmentScale)
		return osBalance
	} else {
		Interest := math.Floor(osBalance * float64(loanDataTab.InterestEffective) * 30 / 36000)
		Principle := monthlyPayment - Interest
		endBalance := osBalance - Principle
		DueDate := time.Now().AddDate(0, index, 0)

		installmentScale = models.SkalaRentalTab{
			Custcode:   customerCode,
			Counter:    int8(index),
			Osbalance:  osBalance,
			EndBalance: endBalance,
			DueDate:    DueDate,
			EffRate:    float64(loanDataTab.InterestEffective),
			Rental:     monthlyPayment,
			Principle:  Principle,
			Interest:   Interest,
			Inputdate:  time.Now(),
		}

		if endBalance < 0 {
			installmentScale.EndBalance = 0
			installmentScale.Principle = osBalance
		}
		r.db.Create(&installmentScale)
		return endBalance
	}

}
