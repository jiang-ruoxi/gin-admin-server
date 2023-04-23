package scategory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scategory"
	scategoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/scategory/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryApi struct {
}

var categoryService = service.ServiceGroupApp.ScategoryServiceGroup.CategoryService

//CreateCategory 创建数据
func (categoryApi *CategoryApi) CreateCategory(c *gin.Context) {
	var category scategory.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
    }
	if err := utils.Verify(category, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := categoryService.CreateCategory(category); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

//DeleteCategory 删除记录
func (categoryApi *CategoryApi) DeleteCategory(c *gin.Context) {
	var category scategory.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := categoryService.DeleteCategory(category); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

//DeleteCategoryByIds 批量删除
func (categoryApi *CategoryApi) DeleteCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := categoryService.DeleteCategoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

//UpdateCategory 更新数据
func (categoryApi *CategoryApi) UpdateCategory(c *gin.Context) {
	var category scategory.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
      }
    if err := utils.Verify(category, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := categoryService.UpdateCategory(category); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

//FindCategory 用查询数据
func (categoryApi *CategoryApi) FindCategory(c *gin.Context) {
	var category scategory.Category
	err := c.ShouldBindQuery(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recategory, err := categoryService.GetCategory(category.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recategory": recategory}, c)
	}
}

//GetCategoryList 列表数据分页
func (categoryApi *CategoryApi) GetCategoryList(c *gin.Context) {
	var pageInfo scategoryReq.CategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := categoryService.GetCategoryInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

//GetCategoryListAll 所有数据
func (categoryApi *CategoryApi) GetCategoryListAll(c *gin.Context) {
	var pageInfo scategoryReq.CategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := categoryService.GetCategoryListAll(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
		}, "获取成功", c)
	}
}
