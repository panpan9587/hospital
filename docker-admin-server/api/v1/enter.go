package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/goodsPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/goodspkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	GoodsApiGroup    goods.ApiGroup
	GoodspkgApiGroup goodspkg.ApiGroup
	GoodsPkgApiGroup goodsPkg.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
