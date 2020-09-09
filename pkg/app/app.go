package app

import (
	"github.com/frankie/go-app/pkg/ecode"
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

func (r *Respone) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Respone) ToErrorResponse(err *ecode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg()}
	details := err.Detauls()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
