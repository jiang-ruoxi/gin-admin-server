package sanswer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sanswer"
	sanswerReq "github.com/flipped-aurora/gin-vue-admin/server/model/sanswer/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sbaike"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scategory"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AnswerService struct {
}

//GetAnswerInfoList 分页获取Answer记录
func (answerService *AnswerService) GetAnswerInfoList(info sanswerReq.AnswerSearch) (list []map[string]interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.WebGDB().Model(sanswer.Answer{})
	var answers []sanswer.Answer
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != "" && info.EndCreatedAt != "" {
		db = db.Where("add_time >= ? AND add_time < ?", utils.FormatDateToUnixTime(info.StartCreatedAt), utils.FormatDateToUnixTime(info.EndCreatedAt))
	}
	if info.CategoryId > 0 {
		db = db.Where("category_id = ?",info.CategoryId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&answers).Error

	//查询问题
	questionIds := make([]int, 0)
	for _, item := range answers {
		questionIds = append(questionIds, item.QuestionId)
	}

	// 创建db
	db1 := global.WebGDB().Model(sbaike.Baike{})
	var baiKeList []sbaike.Baike
	db1.Select("id", "question").Where("id in ?", questionIds).Find(&baiKeList)

	//查询栏目
	categoryIds := make([]int, 0)
	for _, item := range answers {
		categoryIds = append(categoryIds, item.CategoryId)
	}
	db2 := global.WebGDB().Model(scategory.Category{})
	var categoryList []scategory.Category
	db2.Select("id", "name").Where("id in ?", categoryIds).Find(&categoryList)

	oneMap := make(map[string]interface{}, 0)
	for _, v := range answers {
		oneMap["id"] = v.Id
		oneMap["open_id"] = v.OpenId
		oneMap["is_select"] = v.IsSelect
		oneMap["right_select"] = v.RightSelect
		oneMap["add_time"] = utils.FormatDateFromUnix(int64(v.AddTime))
		for _, d := range baiKeList {
			if v.QuestionId == d.ID {
				oneMap["question"] = d.Question
			}
		}
		for _, d := range categoryList {
			if v.CategoryId == d.ID {
				oneMap["category"] = d.Name
			}
		}
		list = append(list, oneMap)
	}

	return list, total, err
}
