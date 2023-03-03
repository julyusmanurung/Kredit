package pencairan

import (
	"github.com/julyusmanurung/Kredit/models"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository interface {
	GetRecentCreditApplicant() ([]models.StagingCustomer, error)
	GetAllBranch() ([]models.BranchTab, error)
	GetAllCompany() ([]models.MstCompanyTab, error)
	GetAllApprovalStatusNine() ([]CustomerJoinLoan, error)
	GetAllApprovalStatusNineFilter(branch string, channeling_company string, startDate string, endDate string) ([]CustomerJoinLoan, error)
	UpdateApprovalStatus(checkedList DataPostPPK) error
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

func (r *repository) GetRecentCreditApplicant() ([]models.StagingCustomer, error) {
	currentTime := time.Now()
	var stagingCustomer []models.StagingCustomer

	res := r.db.Where("sc_create_date = ? AND sc_flag= ?", currentTime.Format("2006-01-02"), "0").Find(&stagingCustomer)
	if res.Error != nil {
		log.Println("get recent credit applicant error: ", res.Error)
		return nil, res.Error
	}

	for _, item := range stagingCustomer {
		r.startValidation(item)
	}

	return stagingCustomer, res.Error
}

func (r *repository) GetAllBranch() ([]models.BranchTab, error) {
	var branchTab []models.BranchTab

	res := r.db.Find(&branchTab)
	if res.Error != nil {
		log.Println("get all branch error: ", res.Error)
		return nil, res.Error
	}

	return branchTab, res.Error
}

func (r *repository) GetAllCompany() ([]models.MstCompanyTab, error) {
	var companyTab []models.MstCompanyTab

	res := r.db.Find(&companyTab)
	if res.Error != nil {
		log.Println("get all company error: ", res.Error)
		return nil, res.Error
	}

	return companyTab, res.Error
}

func (r *repository) GetAllApprovalStatusNine() ([]CustomerJoinLoan, error) {
	var customerLoan []CustomerJoinLoan

	res := r.db.Raw(`	SELECT 
	    						ROW_NUMBER() OVER (Order by cdt.name) AS RowNumber,
	    						cdt.ppk, 
	    						cdt.name,
	    						cdt.channeling_company ,  
	    						ldt.loan_amount,
	    						cdt.drawdown_date,
	    						ldt.loan_period ,
	    						ldt.interest_effective,
	    						ldt.branch
							FROM 
							    customer_data_tab cdt left join Loan_Data_Tab ldt 
							ON 
							    cdt.custcode = ldt.custcode
							WHERE
								approval_status = '9'`).Scan(&customerLoan)

	if res.Error != nil {
		log.Println("get approval status nine error")
	}

	return customerLoan, res.Error
}

func (r *repository) GetAllApprovalStatusNineFilter(branch string, channeling_company string, startDate string, endDate string) ([]CustomerJoinLoan, error) {
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

	res := r.db.Raw(`	SELECT
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
						    approval_status='9'
						`+queryDrawdownDate+queryBranch+queryChannelingCompany, branch, channeling_company, startDate, endDate).Scan(&customerLoan)

	if res.Error != nil {
		log.Println("get approval status nine error")
	}

	return customerLoan, res.Error
}

func (r *repository) UpdateApprovalStatus(checkedList DataPostPPK) error {
	for _, ppk := range checkedList.PPK {
		res := r.db.Where("ppk = ?", ppk).Updates(models.CustomerDataTab{ApprovalStatus: "0"})
		if res.Error != nil {
			log.Println("update approval status error: ", res.Error)
			return res.Error
		}
	}
	return nil
}

/*
	-- Use this to check generated query --

	query := r.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Raw("put your query here")
	})

	fmt.Println(query)
*/
