package goodspkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsRouter struct {
}

// InitGoodsRouter 初始化 goods表 路由信息
func (s *GoodsRouter) InitGoodsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	goodsRouter := Router.Group("goods").Use(middleware.OperationRecord())
	goodsRouterWithoutRecord := Router.Group("goods")
	goodsRouterWithoutAuth := PublicRouter.Group("goods")

	var goodsApi = v1.ApiGroupApp.GoodspkgApiGroup.GoodsApi
	{
		goodsRouter.POST("createGoods", goodsApi.CreateGoods)   // 新建goods表
		goodsRouter.DELETE("deleteGoods", goodsApi.DeleteGoods) // 删除goods表
		goodsRouter.DELETE("deleteGoodsByIds", goodsApi.DeleteGoodsByIds) // 批量删除goods表
		goodsRouter.PUT("updateGoods", goodsApi.UpdateGoods)    // 更新goods表
	}
	{
		goodsRouterWithoutRecord.GET("findGoods", goodsApi.FindGoods)        // 根据ID获取goods表
		goodsRouterWithoutRecord.GET("getGoodsList", goodsApi.GetGoodsList)  // 获取goods表列表
	}
	{
	    goodsRouterWithoutAuth.GET("getGoodsPublic", goodsApi.GetGoodsPublic)  // 获取goods表列表
	}
}
