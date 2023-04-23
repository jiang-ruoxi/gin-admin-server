package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/scategory"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CategorySearch struct{
    scategory.Category
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
