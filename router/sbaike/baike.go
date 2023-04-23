package sbaike

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BaikeRouter struct {
}

// InitBaikeRouter 初始化 Baike 路由信息
func (s *BaikeRouter) InitBaikeRouter(Router *gin.RouterGroup) {
	baikeRouter := Router.Group("baike").Use(middleware.OperationRecord())
	baikeRouterWithoutRecord := Router.Group("baike")
	var baikeApi = v1.ApiGroupApp.SbaikeApiGroup.BaikeApi
	{
		baikeRouter.POST("createBaike", baikeApi.CreateBaike)   // 新建Baike
		baikeRouter.DELETE("deleteBaike", baikeApi.DeleteBaike) // 删除Baike
		baikeRouter.DELETE("deleteBaikeByIds", baikeApi.DeleteBaikeByIds) // 批量删除Baike
		baikeRouter.PUT("updateBaike", baikeApi.UpdateBaike)    // 更新Baike
	}
	{
		baikeRouterWithoutRecord.GET("findBaike", baikeApi.FindBaike)        // 根据ID获取Baike
		baikeRouterWithoutRecord.GET("getBaikeList", baikeApi.GetBaikeList)  // 获取Baike列表
	}
}
