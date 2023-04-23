package scategory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scategory"
	scategoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/scategory/request"
	"time"
)

type CategoryService struct {
}

//CreateCategory 创建Category记录
func (categoryService *CategoryService) CreateCategory(category scategory.Category) (err error) {
	db := global.WebGDB().Model(scategory.Category{})
	now := int(time.Now().UTC().Unix())
	category.AddTime = &now
	db = db.Create(&category)
	err = db.Error
	return err
}

//DeleteCategory 删除Category记录
func (categoryService *CategoryService) DeleteCategory(category scategory.Category) (err error) {
	db := global.WebGDB().Model(scategory.Category{})
	db = db.Delete(&category)
	err = db.Error
	return err
}

//DeleteCategoryByIds 批量删除Category记录
func (categoryService *CategoryService) DeleteCategoryByIds(ids request.IdsReq) (err error) {
	db := global.WebGDB().Model(scategory.Category{})
	db = db.Delete(&[]scategory.Category{}, "id in ?", ids.Ids)
	err = db.Error
	return err
}

//UpdateCategory 更新Category记录
func (categoryService *CategoryService) UpdateCategory(category scategory.Category) (err error) {
	db := global.WebGDB().Model(scategory.Category{})
	err = db.Omit("Id", "add_time").Where("id = ?", category.ID).Updates(category).Error
	err = db.Error
	return err
}

//GetCategory 根据id获取Category记录
func (categoryService *CategoryService) GetCategory(id int) (category scategory.Category, err error) {
	db := global.WebGDB().Model(scategory.Category{})
	db = db.Where("id = ?", id).First(&category)
	err = db.Error
	return
}

//GetCategoryInfoList 分页获取Category记录
func (categoryService *CategoryService) GetCategoryInfoList(info scategoryReq.CategorySearch) (list []scategory.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.WebGDB().Model(scategory.Category{})
	var categorys []scategory.Category
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&categorys).Error
	return categorys, total, err
}

//GetCategoryListAll 获取Category所有记录
func (categoryService *CategoryService) GetCategoryListAll() (list []scategory.Category, err error) {
	db := global.WebGDB().Model(scategory.Category{})
	var categoryList []scategory.Category
	err = db.Where("status = ?", 1).Find(&categoryList).Error
	return categoryList, err
}
