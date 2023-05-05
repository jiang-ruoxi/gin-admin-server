package suser

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/suser"
	suserReq "github.com/flipped-aurora/gin-vue-admin/server/model/suser/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type UserService struct {
}

// CreateUser 创建Baike记录
func (userService *UserService) CreateUser(user suser.User) (err error) {
	db := global.WebGDB().Model(suser.User{}).Debug()
	now := int(time.Now().UTC().Unix())
	user.AddTime = strconv.Itoa(now)
	db = db.Create(&user)
	err = db.Error
	return err
}

// UpdateUser 更新Baike记录
func (userService *UserService) UpdateUser(user suser.User) (err error) {
	db := global.WebGDB().Model(suser.User{}).Debug()
	db = db.Omit("Id", "add_time").Where("id = ?", user.Id).Updates(user)
	err = db.Error
	return err
}

// DeleteUser 删除Baike记录
func (userService *UserService) DeleteUser(user suser.User) (err error) {
	db := global.WebGDB().Model(suser.User{}).Debug()
	db = db.Delete(&user)
	err = db.Error
	return err
}

// GetUser 根据id获取user记录
func (userService *UserService) GetUser(id int) (user suser.User, err error) {
	db := global.WebGDB().Model(suser.User{})
	db = db.Where("id = ?", id).First(&user)
	err = db.Error
	return
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
