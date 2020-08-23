package models

type Model struct {
	ID          int    `gorm:"primary_key" json:"id"`
	CreatedBy   string `json:"created_by"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	UpdatedBy   string `json:"updated_by"`
	DeletedDate string `json:"deleted_date"`
	IsDel       uint8  `json:"is_del"`
}
