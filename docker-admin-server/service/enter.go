package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/service/goodsPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/goodspkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	GoodsServiceGroup    goods.ServiceGroup
	GoodspkgServiceGroup goodspkg.ServiceGroup
	GoodsPkgServiceGroup goodsPkg.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
