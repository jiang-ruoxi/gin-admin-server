package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/suser"
)

type UserSearch struct{
    suser.User
    StartCreatedAt string `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   string `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
