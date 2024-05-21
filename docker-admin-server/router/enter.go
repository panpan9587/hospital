package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/router/goodsPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/goodspkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Goods    goods.RouterGroup
	Goodspkg goodspkg.RouterGroup
	GoodsPkg goodsPkg.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
