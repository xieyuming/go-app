package models

type Tag struct {
	*Model
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (a Tag) TableName() string {
	return "bolg_tag"
}
