package models

type IdTab struct {
	CODE  string `gorm:"default:006"`
	DIGIT int    `gorm:"default:10"`
	VALUE int
}

func (m *IdTab) TableName() string {
	return "id_tab"
}
