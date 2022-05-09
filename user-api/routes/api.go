package routes

import (
	. "github.com/869413421/micro-service/user-api/app/http/controllers"
	"github.com/gin-gonic/gin"
)

var userController = NewUserController()
var passwordController = NewPasswordResetController()



// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	// 用户管理路由,所有路由必须包含user，因为micro网关只会映射路径中包含user的路由
	api := router.Group("/")
	{
		// 登录
		api.POST("/user/token", userController.Auth)
		// 注册
		api.POST("/user", userController.Store)
	}
	{
		// 发起重置密码
		api.POST("user/password", passwordController.Store)
		// 重置密码
		api.PUT("user/password", passwordController.ResetPassword)
	}

	userApi := api.Group("user")
	{
		// 用户列表
		userApi.GET("", userController.Index)
		// 获取单个用户
		userApi.GET("/:id", userController.Show)
		// 更新用户
		userApi.PUT("/:id", userController.Update)
		// 删除用户
		userApi.DELETE("/:id", userController.Delete)
	}
}
