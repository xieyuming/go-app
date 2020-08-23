package models

type ArticleTag struct {
	*Model
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}
