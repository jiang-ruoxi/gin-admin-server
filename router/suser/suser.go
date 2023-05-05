package suser

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// InitUserRouter 初始化 User 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterWithoutRecord := Router.Group("user")
	var userApi = v1.ApiGroupApp.SuserApiGroup.UserApi
	{
		userRouterWithoutRecord.POST("createUser", userApi.CreateUser) // 新建Baike
		userRouterWithoutRecord.PUT("updateUser", userApi.UpdateUser)  // 更新Baike
	}
	{
		userRouterWithoutRecord.GET("getSUserList", userApi.GetUserList) // 获取User列表
		userRouterWithoutRecord.GET("findUser", userApi.FindUser)        // 获取User列表
	}
}
