package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/sanswer"
	"github.com/flipped-aurora/gin-vue-admin/server/router/sbaike"
	"github.com/flipped-aurora/gin-vue-admin/server/router/scategory"
	"github.com/flipped-aurora/gin-vue-admin/server/router/slike"
	"github.com/flipped-aurora/gin-vue-admin/server/router/suser"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Scategory scategory.RouterGroup
	Sbaike    sbaike.RouterGroup
	Suser     suser.RouterGroup
	Sanswer   sanswer.RouterGroup
	Slike     slike.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
