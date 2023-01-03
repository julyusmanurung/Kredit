package models

type BranchTab struct {
	Code        string `json:"code" gorm:"not null; type: varchar(50); unique"`
	Description string `json:"description" gorm:"type: varchar(50)"`
	Address1    string `json:"address1" gorm:"type: varchar(50)"`
	Address2    string `json:"address2" gorm:"type: varchar(50)"`
	City        string `json:"city" gorm:"type: varchar(50)"`
	Zip         string `json:"zip" gorm:"type: varchar(6)"`
	Phone       string `json:"phone" gorm:"type: varchar(15)"`
	Fax         string `json:"fax" gorm:"type: varchar(15)"`
}

func (m *BranchTab) TableName() string {
	return "branch_tab"
}
