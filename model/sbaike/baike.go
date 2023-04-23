// 自动生成模板Baike
package sbaike

// Baike 结构体
type Baike struct {
	ID         int    `gorm:"primarykey"` // 主键ID
	CategoryId int    `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:类型;"`
	Question   string `json:"question" form:"question" gorm:"column:question;comment:标题;size:256;"`
	OptionA    string `json:"optionA" form:"optionA" gorm:"column:option_a;comment:选项a;size:1024;"`
	OptionB    string `json:"optionB" form:"optionB" gorm:"column:option_b;comment:选项b;size:1024;"`
	OptionC    string `json:"optionC" form:"optionC" gorm:"column:option_c;comment:选项c;size:1024;"`
	OptionD    string `json:"optionD" form:"optionD" gorm:"column:option_d;comment:选项d;size:1024;"`
	Answer     string `json:"answer" form:"answer" gorm:"column:answer;comment:答案;size:191;"`
	Analytic   string `json:"analytic" form:"analytic" gorm:"column:analytic;comment:原因;size:1024;"`
	AddTime    int    `json:"addTime" form:"addTime" gorm:"column:add_time;comment:添加时间;"`
}

// TableName Baike 表名
func (Baike) TableName() string {
	return "s_baike"
}
