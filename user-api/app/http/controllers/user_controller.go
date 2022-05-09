package controllers

import (
	base "github.com/869413421/micro-service/common/api/http/controller"
	"github.com/869413421/micro-service/common/pkg/container"
	"github.com/869413421/micro-service/common/pkg/types"
	pb "github.com/869413421/micro-service/user/proto/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	base.BaseController
}

// NewUserController 初始化用户控制器
func NewUserController() *UserController {
	return &UserController{}
}

// Index 用户分页
func (controller *UserController) Index(context *gin.Context) {
	// 1.构建发起请求参数
	pagination := &base.Pagination{}
	err := context.BindQuery(pagination)
	if err != nil {
		controller.ResponseJson(context, http.StatusForbidden, "pagination params required", []string{})
	}

	// 2.请求用户服务
	req := &pb.PaginationRequest{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
	}
	client := container.GetUserServiceClient()
	rsp, err := client.Pagination(context, req)
	if err != nil {
		controller.ResponseJson(context, http.StatusInternalServerError, err.Error(), []string{})
		return
	}

	//3.响应用户信息
	controller.ResponseJson(context, http.StatusOK, "", rsp)
}

// Store 创建用户
func (controller *UserController) Store(context *gin.Context) {
	// 1.构建微服务请求体
	req := &pb.CreateRequest{}
	client := container.GetUserServiceClient()
	err := context.BindJSON(req)
	if err != nil {
		controller.ResponseJson(context, http.StatusForbidden, "json params error", []string{})
		return
	}

	// 2.发起创建请求
	rsp, err := client.Create(context, req)
	if err != nil {
		controller.ResponseJson(context, http.StatusInternalServerError, err.Error(), []string{})
		return
	}

	// 3.响应用户信息
	controller.ResponseJson(context, http.StatusOK, "", rsp.User)
}

// Update 更新用户
func (controller *UserController) Update(context *gin.Context) {
	// 1.获取路由中的ID
	idStr := context.Param("id")
	if idStr == "" {
		controller.ResponseJson(context, http.StatusForbidden, "route id required", []string{})
		return
	}

	// 2.构建微服务请求体
	req := &pb.UpdateRequest{}
	err := context.BindJSON(req)
	if err != nil {
		controller.ResponseJson(context, http.StatusForbidden, "json params error", []string{})
		return
	}
	id, _ := types.StringToInt(idStr)
	req.Id = uint64(id)

	// 3.调用服务请求
	client := container.GetUserServiceClient()
	rsp, err := client.Update(context, req)
	if err != nil {
		controller.ResponseJson(context, http.StatusInternalServerError, err.Error(), []string{})
		return
	}

	//5.响应用户信息
	controller.ResponseJson(context, http.StatusOK, "", rsp.User)
}

// Delete 删除用户
func (controller *UserController) Delete(context *gin.Context) {
	//1.获取路由中的ID
	idStr := context.Param("id")
	if idStr == "" {
		controller.ResponseJson(context, http.StatusForbidden, "route id required", []string{})
		return
	}

	//2.构建微服务请求体发起请求
	id, _ := types.StringToInt(idStr)
	req := &pb.DeleteRequest{Id: uint64(id)}
	client := container.GetUserServiceClient()
	rsp, err := client.Delete(context, req)
	if err != nil {
		controller.ResponseJson(context, http.StatusInternalServerError, err.Error(), []string{})
		return
	}

	//3.响应用户信息
	controller.ResponseJson(context, http.StatusOK, "", rsp.User)
}

// Show 展示单个用户
func (controller *UserController) Show(context *gin.Context) {
	// 1.获取路由中的ID
	idStr := context.Param("id")
	if idStr == "" {
		controller.ResponseJson(context, http.StatusForbidden, "route id required", []string{})
		return
	}

	// 2.构建微服务请求体发起请求
	id, _ := types.StringToInt(idStr)
	req := &pb.GetRequest{Id: uint64(id)}
	client := container.GetUserServiceClient()
	rsp, err := client.Get(context, req)
	if err != nil {
		controller.ResponseJson(context, http.StatusInternalServerError, err.Error(), []string{})
		return
	}

	// 3.响应用户信息
	controller.ResponseJson(context, http.StatusOK, "", rsp.User)
}

// Auth 认证
func (controller *UserController) Auth(context *gin.Context) {
	//1.构建微服务请求体
	req := &pb.AuthRequest{}
	err := context.BindJSON(req)
	if err != nil {
		controller.ResponseJson(context, http.StatusForbidden, "json params error", []string{})
		return
	}

	//2.发起请求
	client := container.GetUserServiceClient()
	rsp, err := client.Auth(context, req)
	if err != nil {
		controller.ResponseJson(context, http.StatusInternalServerError, err.Error(), []string{})
		return
	}

	//3.响应用户信息
	controller.ResponseJson(context, http.StatusOK, "", rsp)
}
