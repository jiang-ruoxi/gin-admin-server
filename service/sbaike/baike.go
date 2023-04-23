package sbaike

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sbaike"
	sbaikeReq "github.com/flipped-aurora/gin-vue-admin/server/model/sbaike/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scategory"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

type BaikeService struct {
}

//CreateBaike 创建Baike记录
func (baikeService *BaikeService) CreateBaike(baike sbaike.Baike) (err error) {
	db := global.WebGDB().Model(sbaike.Baike{})
	now := int(time.Now().UTC().Unix())
	baike.AddTime = now
	db = db.Create(&baike)
	err = db.Error
	baikeService.HttpGetMethod(baike.CategoryId)
	return err
}

//DeleteBaike 删除Baike记录
func (baikeService *BaikeService) DeleteBaike(baike sbaike.Baike) (err error) {
	db := global.WebGDB().Model(sbaike.Baike{})
	db = db.Delete(&baike)
	err = db.Error
	return err
}

//DeleteBaikeByIds 批量删除Baike记录
func (baikeService *BaikeService) DeleteBaikeByIds(ids request.IdsReq) (err error) {
	db := global.WebGDB().Model(sbaike.Baike{})
	db = db.Delete(&[]sbaike.Baike{}, "id in ?", ids.Ids)
	err = db.Error
	return err
}

//UpdateBaike 更新Baike记录
func (baikeService *BaikeService) UpdateBaike(baike sbaike.Baike) (err error) {
	db := global.WebGDB().Model(sbaike.Baike{})
	err = db.Omit("Id", "add_time").Where("id = ?", baike.ID).Updates(baike).Error
	err = db.Error
	baikeService.HttpGetMethod(baike.CategoryId)
	return err
}

//GetBaike 根据id获取Baike记录
func (baikeService *BaikeService) GetBaike(id int) (baike sbaike.Baike, err error) {
	db := global.WebGDB().Model(sbaike.Baike{})
	db = db.Where("id = ?", id).First(&baike)
	err = db.Error
	return
}

// GetBaikeInfoList 分页获取Baike记录
func (baikeService *BaikeService) GetBaikeInfoList(info sbaikeReq.BaikeSearch) (list []map[string]interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.WebGDB().Model(sbaike.Baike{})
	var listData []sbaike.Baike
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CategoryId > 0 {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	if info.Question != "" {
		db = db.Where("question LIKE ?", "%"+info.Question+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	db = baikeService.MakeBaiKeListOrder(db, info.Sort, info.Order)
	err = db.Limit(limit).Offset(offset).Find(&listData).Error

	// 创建db
	db1 := global.WebGDB().Model(scategory.Category{})
	var categoryList []scategory.Category

	err = db1.Find(&categoryList).Error
	mapData := make(map[interface{}]interface{})
	for _, item := range categoryList {
		mapData[int(item.ID)] = item.Name
	}

	for _, item := range listData {
		list = append(list, map[string]interface{}{
			"ID":       item.ID,
			"question": item.Question,
			"optionA":  item.OptionA,
			"optionB":  item.OptionB,
			"optionC":  item.OptionC,
			"optionD":  item.OptionD,
			"category": mapData[item.CategoryId],
			"answer":   item.Answer,
		})
	}
	return list, total, err
}

//MakeBaiKeListOrder 组合排序
func (baikeService *BaikeService) MakeBaiKeListOrder(db *gorm.DB, sort, order string) *gorm.DB {
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["ID"] = true
	if orderMap[sort] {
		OrderStr = sort
		if order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}
	return db
}

func (baikeService *BaikeService) HttpGetMethod(categoryId int) {
	client := &http.Client{}
	url := fmt.Sprintf(global.DELETE_QUEUE_URL, categoryId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-token", "wechat")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}
