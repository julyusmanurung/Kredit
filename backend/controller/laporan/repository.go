package laporan

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository interface {
	GetLaporanData(branch string, channeling_company string, startDate string, endDate string) ([]CustomerJoinLoan, error)
}

type repository struct {
	db *gorm.DB
}

type CustomerJoinLoan struct {
	Rownumber         int       `json:"rownumber"`
	PPK               string    `json:"ppk"`
	Name              string    `json:"name"`
	ChannelingCompany string    `json:"channeling_company"`
	DrawdownDate      time.Time `json:"drawdown_date"`
	LoanAmount        float64   `json:"loan_amount"`
	LoanPeriod        string    `json:"loan_period"`
	InterestEffective float32   `json:"interest_effective"`
	Branch            string    `json:"branch"`
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetLaporanData(branch string, channeling_company string, startDate string, endDate string) ([]CustomerJoinLoan, error) {
	var customerLoan []CustomerJoinLoan

	queryBranch := ""

	if branch == "000" {
		queryBranch += " AND ldt.branch LIKE $1"
		branch = "%%"
	} else {
		queryBranch += " AND ldt.branch = $1"
	}

	queryChannelingCompany := ""
	if channeling_company == "" {
		queryChannelingCompany += " AND cdt.channeling_company LIKE $2"
		channeling_company = "%%"
	} else {
		queryChannelingCompany += " AND cdt.channeling_company = $2"
	}

	queryDrawdownDate := " AND drawdown_date BETWEEN $3 AND $4"
	if startDate != "" && endDate == "" {
		endDate = "31-12-9999"
	} else if startDate == "" && endDate != "" {
		startDate = "01-01-0001"
	} else if startDate == "" && endDate == "" {
		startDate = "01-01-0001"
		endDate = "31-12-9999"
	}

	res := r.db.Raw(`SELECT
							ROW_NUMBER() OVER (Order by cdt.name) AS RowNumber, 
							cdt.ppk, 
    						cdt.name,
							cdt.channeling_company,
							ldt.loan_amount,
							cdt.drawdown_date,
							ldt.loan_period,
							ldt.interest_effective,
							ldt.branch
						FROM 
						    customer_data_tab cdt left join Loan_Data_Tab ldt ON cdt.custcode = ldt.custcode 
						WHERE 
						    approval_status='0'
						`+queryBranch+queryChannelingCompany+queryDrawdownDate, branch, channeling_company, startDate, endDate).Scan(&customerLoan)

	if res.Error != nil {
		log.Println("get approval 0 error")
	}

	return customerLoan, res.Error
}
