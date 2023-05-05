package suser

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/suser"
	suserReq "github.com/flipped-aurora/gin-vue-admin/server/model/suser/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

var userService = service.ServiceGroupApp.SuserServiceGroup.UserService

// CreateUser 创建数据
func (userApi *UserApi) CreateUser(c *gin.Context) {
	var user suser.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"OpenId":   {utils.NotEmpty()},
		"NickName": {utils.NotEmpty()},
		"HeadUrl":  {utils.NotEmpty()},
	}
	if err := utils.Verify(user, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.CreateUser(user); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateUser 更新数据
func (userApi *UserApi) UpdateUser(c *gin.Context) {
	var user suser.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"OpenId":   {utils.NotEmpty()},
		"NickName": {utils.NotEmpty()},
		"HeadUrl":  {utils.NotEmpty()},
	}
	if err := utils.Verify(user, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.UpdateUser(user); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetUserList 分页获取User列表
func (userApi *UserApi) GetUserList(c *gin.Context) {
	var pageInfo suserReq.UserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userService.GetUserInfoList(pageInfo); err != nil {
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

// FindUser 查询数据
func (userApi *UserApi) FindUser(c *gin.Context) {
	var user suser.User
	err := c.ShouldBindQuery(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebaike, err := userService.GetUser(user.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebaike": rebaike}, c)
	}
}
