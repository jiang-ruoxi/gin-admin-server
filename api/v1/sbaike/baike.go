package sbaike

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sbaike"
	sbaikeReq "github.com/flipped-aurora/gin-vue-admin/server/model/sbaike/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaikeApi struct {
}

var baikeService = service.ServiceGroupApp.SbaikeServiceGroup.BaikeService

//CreateBaike 创建数据
func (baikeApi *BaikeApi) CreateBaike(c *gin.Context) {
	var baike sbaike.Baike
	err := c.ShouldBindJSON(&baike)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "CategoryId":{utils.NotEmpty()},
        "Question":{utils.NotEmpty()},
        "OptionA":{utils.NotEmpty()},
        "OptionB":{utils.NotEmpty()},
        "OptionC":{utils.NotEmpty()},
        "OptionD":{utils.NotEmpty()},
        "Answer":{utils.NotEmpty()},
    }
	if err := utils.Verify(baike, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := baikeService.CreateBaike(baike); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//DeleteBaike 删除数据
func (baikeApi *BaikeApi) DeleteBaike(c *gin.Context) {
	var baike sbaike.Baike
	err := c.ShouldBindJSON(&baike)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := baikeService.DeleteBaike(baike); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

//DeleteBaikeByIds 批量删除数据
func (baikeApi *BaikeApi) DeleteBaikeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := baikeService.DeleteBaikeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

//UpdateBaike 更新数据
func (baikeApi *BaikeApi) UpdateBaike(c *gin.Context) {
	var baike sbaike.Baike
	err := c.ShouldBindJSON(&baike)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "CategoryId":{utils.NotEmpty()},
          "Question":{utils.NotEmpty()},
          "OptionA":{utils.NotEmpty()},
          "OptionB":{utils.NotEmpty()},
          "OptionC":{utils.NotEmpty()},
          "OptionD":{utils.NotEmpty()},
          "Answer":{utils.NotEmpty()},
      }
    if err := utils.Verify(baike, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := baikeService.UpdateBaike(baike); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

//FindBaike 查询数据
func (baikeApi *BaikeApi) FindBaike(c *gin.Context) {
	var baike sbaike.Baike
	err := c.ShouldBindQuery(&baike)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebaike, err := baikeService.GetBaike(baike.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebaike": rebaike}, c)
	}
}

//GetBaikeList 列表数据
func (baikeApi *BaikeApi) GetBaikeList(c *gin.Context) {
	var pageInfo sbaikeReq.BaikeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := baikeService.GetBaikeInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
