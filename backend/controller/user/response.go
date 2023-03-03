package user

type DataResponse struct {
	UserId  string `json:"user_id" gorm:"type: varchar(6);not null"`
	Name    string `json:"name" gorm:"type: varchar(20);not null"`
	Email   string `json:"email" gorm:"unique;not null;type:varchar(30)"`
	Level   string `json:"level" gorm:"type: varchar(20);not null"`
	Jabatan string `json:"jabatan" gorm:"type: varchar(128);not null"`
}
