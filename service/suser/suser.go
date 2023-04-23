package suser

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/suser"
	suserReq "github.com/flipped-aurora/gin-vue-admin/server/model/suser/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"strconv"
)

type UserService struct {
}

// GetUserInfoList 分页获取User记录
func (userService *UserService) GetUserInfoList(info suserReq.UserSearch) (list []suser.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.WebGDB().Model(suser.User{})
	var users []suser.User
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != "" && info.EndCreatedAt != "" {
		db = db.Where("add_time >= ? AND add_time < ?", utils.FormatDateToUnixTime(info.StartCreatedAt), utils.FormatDateToUnixTime(info.EndCreatedAt))
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&users).Error

	for idx, item := range users {
		if aTime, err := strconv.ParseInt(item.AddTime, 10, 64); err != nil {
			users[idx].AddTime = "-"
		} else {
			users[idx].AddTime = utils.FormatDateFromUnix(aTime)
		}
	}

	return users, total, err
}
