package models

import "github.com/frankie/go-app/pkg/app"

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
