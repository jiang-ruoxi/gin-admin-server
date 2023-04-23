// 自动生成模板Category
package scategory

// Category 结构体
type Category struct {
	ID      int    `gorm:"primarykey"` // 主键ID
	Name    string `json:"name" form:"name" gorm:"column:name;comment:栏目名称;size:256;"`
	Status  *bool  `json:"status" form:"status" gorm:"column:status;comment:状态,1启用,0禁用;"`
	AddTime *int   `json:"addTime" form:"addTime" gorm:"column:add_time;comment:添加时间;"`
}

// TableName Category 表名
func (Category) TableName() string {
	return "s_category"
}
