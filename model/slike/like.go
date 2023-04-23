// 自动生成模板Like
package slike

// Like 结构体
type Like struct {
	Id         int    `json:"id" form:"id" gorm:"primarykey"`
	OpenId     string `json:"openId" form:"openId" gorm:"column:open_id;comment:用户OpendId;size:64;"`
	CategoryId int    `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:类型Id;size:10;"`
	QuestionId int    `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目Id;size:10;"`
	AddTime    int    `json:"addTime" form:"addTime" gorm:"column:add_time;comment:添加时间;size:10;"`
}

// TableName Like 表名
func (Like) TableName() string {
	return "s_like"
}
