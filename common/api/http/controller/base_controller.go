package controller

import "github.com/gin-gonic/gin"

// ResponseErrors 异常信息统一格式
type ResponseErrors map[string]interface{}

// ResponseData 统一响应格式
type ResponseData struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg"`
	Data     interface{} `json:"data"`
}

// Pagination 分页统一结构体
type Pagination struct {
	Page    uint64 `form:"page"`
	PerPage uint32 `form:"perPage"`
}

// BaseController base
type BaseController struct {
}

// NewBaseController 初始化控制器
func NewBaseController() *BaseController {
	return &BaseController{}
}

// ResponseJson 响应json
func (*BaseController) ResponseJson(ctx *gin.Context, code int, errorMsg string, data interface{}) {
	responseData := ResponseData{
		Code:     code,
		ErrorMsg: errorMsg,
		Data:     data,
	}

	ctx.JSON(code, responseData)
	ctx.Abort()
}
