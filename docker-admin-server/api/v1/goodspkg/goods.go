package goodspkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/goodspkg"
    goodspkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/goodspkg/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GoodsApi struct {
}

var goodsService = service.ServiceGroupApp.GoodspkgServiceGroup.GoodsService


// CreateGoods 创建goods表
// @Tags Goods
// @Summary 创建goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goodspkg.Goods true "创建goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goods/createGoods [post]
func (goodsApi *GoodsApi) CreateGoods(c *gin.Context) {
	var goods goodspkg.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goodsService.CreateGoods(&goods); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoods 删除goods表
// @Tags Goods
// @Summary 删除goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goodspkg.Goods true "删除goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goods/deleteGoods [delete]
func (goodsApi *GoodsApi) DeleteGoods(c *gin.Context) {
	ID := c.Query("ID")
	if err := goodsService.DeleteGoods(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsByIds 批量删除goods表
// @Tags Goods
// @Summary 批量删除goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goods/deleteGoodsByIds [delete]
func (goodsApi *GoodsApi) DeleteGoodsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := goodsService.DeleteGoodsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoods 更新goods表
// @Tags Goods
// @Summary 更新goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goodspkg.Goods true "更新goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goods/updateGoods [put]
func (goodsApi *GoodsApi) UpdateGoods(c *gin.Context) {
	var goods goodspkg.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := goodsService.UpdateGoods(goods); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoods 用id查询goods表
// @Tags Goods
// @Summary 用id查询goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goodspkg.Goods true "用id查询goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goods/findGoods [get]
func (goodsApi *GoodsApi) FindGoods(c *gin.Context) {
	ID := c.Query("ID")
	if regoods, err := goodsService.GetGoods(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoods": regoods}, c)
	}
}

// GetGoodsList 分页获取goods表列表
// @Tags Goods
// @Summary 分页获取goods表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goodspkgReq.GoodsSearch true "分页获取goods表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/getGoodsList [get]
func (goodsApi *GoodsApi) GetGoodsList(c *gin.Context) {
	var pageInfo goodspkgReq.GoodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := goodsService.GetGoodsInfoList(pageInfo); err != nil {
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

// GetGoodsPublic 不需要鉴权的goods表接口
// @Tags Goods
// @Summary 不需要鉴权的goods表接口
// @accept application/json
// @Produce application/json
// @Param data query goodspkgReq.GoodsSearch true "分页获取goods表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/getGoodsList [get]
func (goodsApi *GoodsApi) GetGoodsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的goods表接口信息",
    }, "获取成功", c)
}
