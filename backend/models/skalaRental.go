package models

import "time"

type SkalaRentalTab struct {
	Custcode      string    `json:"custcode" gorm:"not null; type: varchar(25);"`
	Counter       int8      `json:"counter" gorm:"type:smallint"`
	Osbalance     float64   `json:"osbalance"  gorm:"type:decimal"`
	EndBalance    float64   `json:"end_balance"  gorm:"type:decimal"`
	DueDate       time.Time `json:"due_date"  gorm:"type:timestamp"`
	EffRate       float64   `json:"eff_rate"  gorm:"type:float"`
	Rental        float64   `json:"rental"  gorm:"type:decimal"`
	Principle     float64   `json:"principle"  gorm:"type:decimal"`
	Interest      float64   `json:"interest"  gorm:"type:decimal"`
	Inputdate     time.Time `json:"inputdate"  gorm:"type:timestamp"`
	Inputby       string    `json:"inputby"  gorm:"type:varchar(50)"`
	Lastmodified  time.Time `json:"lastmodified"  gorm:"type:timestamp"`
	Modifiedby    string    `json:"modifiedby"  gorm:"type:varchar(50)"`
	PaymentDate   time.Time `json:"payment_date"  gorm:"type:timestamp"`
	Penalty       float64   `json:"penalty"  gorm:"type:decimal"`
	PaymentAmount float64   `json:"payment_amount"  gorm:"type:decimal"`
	PaymentType   int8      `json:"payment_type" gorm:"type:smallint"`
}

func (m *SkalaRentalTab) TableName() string {
	return "skala_rental_tab"
}
