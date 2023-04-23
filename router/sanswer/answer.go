package sanswer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AnswerRouter struct {
}

// InitAnswerRouter 初始化 Answer 路由信息
func (s *AnswerRouter) InitAnswerRouter(Router *gin.RouterGroup) {
	answerRouterWithoutRecord := Router.Group("answer")
	var answerApi = v1.ApiGroupApp.SanswerApiGroup.AnswerApi
	{
		answerRouterWithoutRecord.GET("getAnswerList", answerApi.GetAnswerList) // 获取Answer列表
	}
}
