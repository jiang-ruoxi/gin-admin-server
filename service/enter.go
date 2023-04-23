package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/sanswer"
	"github.com/flipped-aurora/gin-vue-admin/server/service/sbaike"
	"github.com/flipped-aurora/gin-vue-admin/server/service/scategory"
	"github.com/flipped-aurora/gin-vue-admin/server/service/slike"
	"github.com/flipped-aurora/gin-vue-admin/server/service/suser"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	ScategoryServiceGroup scategory.ServiceGroup
	SbaikeServiceGroup    sbaike.ServiceGroup
	SuserServiceGroup     suser.ServiceGroup
	SanswerServiceGroup   sanswer.ServiceGroup
	SlikeServiceGroup     slike.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
