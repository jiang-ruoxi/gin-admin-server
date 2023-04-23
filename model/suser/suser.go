// 自动生成模板User
package suser

// User 结构体
type User struct {
	Id         int    `json:"id" form:"id" gorm:"primarykey"`
	OpenId     string `json:"openId" form:"openId" gorm:"column:open_id;comment:用户OpendId;size:64;"`
	NickName   string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:64;"`
	HeadUrl    string `json:"headUrl" form:"headUrl" gorm:"column:head_url;comment:用户头像;size:255;"`
	Area       string `json:"area" form:"area" gorm:"column:area;comment:地区;size:255;"`
	AddTime    string    `json:"addTime" form:"addTime" gorm:"column:add_time;comment:添加时间;size:10;"`
	UpdateTime string    `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;size:10;"`
}

// TableName User 表名
func (User) TableName() string {
	return "s_user"
}
