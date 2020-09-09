package models

import (
	"github.com/frankie/go-app/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (a Tag) TableName() string {
	return "bolg_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.Status)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
