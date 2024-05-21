package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/goods"
    goodsReq "github.com/flipped-aurora/gin-vue-admin/server/model/goods/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GoodsApi struct {
}

var productService = service.ServiceGroupApp.GoodsServiceGroup.GoodsService


// CreateGoods 创建商品
// @Tags Goods
// @Summary 创建商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.Goods true "创建商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /product/createGoods [post]
func (productApi *GoodsApi) CreateGoods(c *gin.Context) {
	var product goods.Goods
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := productService.CreateGoods(&product); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoods 删除商品
// @Tags Goods
// @Summary 删除商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.Goods true "删除商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /product/deleteGoods [delete]
func (productApi *GoodsApi) DeleteGoods(c *gin.Context) {
	ID := c.Query("ID")
	if err := productService.DeleteGoods(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsByIds 批量删除商品
// @Tags Goods
// @Summary 批量删除商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /product/deleteGoodsByIds [delete]
func (productApi *GoodsApi) DeleteGoodsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := productService.DeleteGoodsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoods 更新商品
// @Tags Goods
// @Summary 更新商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goods.Goods true "更新商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /product/updateGoods [put]
func (productApi *GoodsApi) UpdateGoods(c *gin.Context) {
	var product goods.Goods
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := productService.UpdateGoods(product); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoods 用id查询商品
// @Tags Goods
// @Summary 用id查询商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goods.Goods true "用id查询商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /product/findGoods [get]
func (productApi *GoodsApi) FindGoods(c *gin.Context) {
	ID := c.Query("ID")
	if reproduct, err := productService.GetGoods(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproduct": reproduct}, c)
	}
}

// GetGoodsList 分页获取商品列表
// @Tags Goods
// @Summary 分页获取商品列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goodsReq.GoodsSearch true "分页获取商品列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/getGoodsList [get]
func (productApi *GoodsApi) GetGoodsList(c *gin.Context) {
	var pageInfo goodsReq.GoodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := productService.GetGoodsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetGoodsPublic 不需要鉴权的商品接口
// @Tags Goods
// @Summary 不需要鉴权的商品接口
// @accept application/json
// @Produce application/json
// @Param data query goodsReq.GoodsSearch true "分页获取商品列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/getGoodsList [get]
func (productApi *GoodsApi) GetGoodsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的商品接口信息",
    }, "获取成功", c)
}
