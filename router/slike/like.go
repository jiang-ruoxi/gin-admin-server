package slike

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type LikeRouter struct {
}

// InitLikeRouter 初始化 Like 路由信息
func (s *LikeRouter) InitLikeRouter(Router *gin.RouterGroup) {
	likeRouterWithoutRecord := Router.Group("like")
	var likeApi = v1.ApiGroupApp.SlikeApiGroup.LikeApi
	{
		likeRouterWithoutRecord.GET("getLikeList", likeApi.GetLikeList) // 获取Like列表
	}
}
