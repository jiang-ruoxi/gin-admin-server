// 自动生成模板Answer
package sanswer

// Answer 结构体
type Answer struct {
	Id          int    `json:"id" form:"id" gorm:"primarykey"`
	OpenId      string `json:"openId" form:"openId" gorm:"column:open_id;comment:用户OpendId;size:64;"`
	CategoryId  int    `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:类型Id;size:10;"`
	QuestionId  int    `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目Id;size:10;"`
	IsSelect    string `json:"isSelect" form:"isSelect" gorm:"column:is_select;comment:当前选择的答案;"`
	RightSelect string `json:"rightSelect" form:"rightSelect" gorm:"column:right_select;comment:正确答案;"`
	AddTime     int    `json:"addTime" form:"addTime" gorm:"column:add_time;comment:添加时间;size:10;"`
}

// TableName Answer 表名
func (Answer) TableName() string {
	return "s_answer_log"
}
