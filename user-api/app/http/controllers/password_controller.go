package controllers

import (
	"github.com/869413421/micro-service/common/api/http/controller"
	"github.com/869413421/micro-service/common/pkg/container"
	pb "github.com/869413421/micro-service/user/proto/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PasswordResetController 密码控制器
type PasswordResetController struct {
	controller.BaseController
}

// NewPasswordResetController 初始化密码控制器
func NewPasswordResetController() *PasswordResetController {
	return &PasswordResetController{}
}

// Store 创建
func (controller *PasswordResetController) Store(context *gin.Context) {
	// 1.构建微服务请求体
	req := &pb.CreatePasswordResetRequest{}
	client := container.GetUserServiceClient()
	err := context.BindJSON(req)
	if err != nil {
		controller.ResponseJson(context, http.StatusForbidden, "json params error", []string{})
		return
	}

	// 2.发起创建请求
	rsp, err := client.CreatePasswordReset(context, req)
	if err != nil {
		controller.ResponseJson(context, http.StatusInternalServerError, err.Error(), []string{})
		return
	}

	// 3.响应信息
	controller.ResponseJson(context, http.StatusOK, "", rsp)
}

// ResetPassword 重置密码
func (controller *PasswordResetController) ResetPassword(context *gin.Context) {
	// 1.构建微服务请求体
	req := &pb.ResetPasswordRequest{}
	client := container.GetUserServiceClient()
	err := context.BindJSON(req)
	if err != nil {
		controller.ResponseJson(context, http.StatusForbidden, "json params error", []string{})
		return
	}

	// 2.发起创建请求
	rsp, err := client.ResetPassword(context, req)
	if err != nil {
		controller.ResponseJson(context, http.StatusInternalServerError, err.Error(), []string{})
		return
	}

	// 3.响应信息
	controller.ResponseJson(context, http.StatusOK, "", rsp)
}
