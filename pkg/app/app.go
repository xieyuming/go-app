package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Respone struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewRespone(ctx *gin.Context) *Respone {
	return &Respone{ctx}
}

func (r *Respone) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}
