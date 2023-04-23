package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/sanswer"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/sbaike"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/scategory"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/slike"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/suser"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	ScategoryApiGroup scategory.ApiGroup
	SbaikeApiGroup    sbaike.ApiGroup
	SuserApiGroup     suser.ApiGroup
	SanswerApiGroup   sanswer.ApiGroup
	SlikeApiGroup     slike.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
