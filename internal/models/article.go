package models

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	status        uint8  `json:"status"`
}

func (a Article) TableName() string {
	return "blog_article"
}
