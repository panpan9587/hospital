package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsRouter struct {
}

// InitGoodsRouter 初始化 商品 路由信息
func (s *GoodsRouter) InitGoodsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	productRouter := Router.Group("product").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("product")
	productRouterWithoutAuth := PublicRouter.Group("product")

	var productApi = v1.ApiGroupApp.GoodsApiGroup.GoodsApi
	{
		productRouter.POST("createGoods", productApi.CreateGoods)   // 新建商品
		productRouter.DELETE("deleteGoods", productApi.DeleteGoods) // 删除商品
		productRouter.DELETE("deleteGoodsByIds", productApi.DeleteGoodsByIds) // 批量删除商品
		productRouter.PUT("updateGoods", productApi.UpdateGoods)    // 更新商品
	}
	{
		productRouterWithoutRecord.GET("findGoods", productApi.FindGoods)        // 根据ID获取商品
		productRouterWithoutRecord.GET("getGoodsList", productApi.GetGoodsList)  // 获取商品列表
	}
	{
	    productRouterWithoutAuth.GET("getGoodsPublic", productApi.GetGoodsPublic)  // 获取商品列表
	}
}
