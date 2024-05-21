import service from '@/utils/request'

// @Tags Goods
// @Summary 创建goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Goods true "创建goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goods/createGoods [post]
export const createGoods = (data) => {
  return service({
    url: '/goods/createGoods',
    method: 'post',
    data
  })
}

// @Tags Goods
// @Summary 删除goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Goods true "删除goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goods/deleteGoods [delete]
export const deleteGoods = (params) => {
  return service({
    url: '/goods/deleteGoods',
    method: 'delete',
    params
  })
}

// @Tags Goods
// @Summary 批量删除goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goods/deleteGoods [delete]
export const deleteGoodsByIds = (params) => {
  return service({
    url: '/goods/deleteGoodsByIds',
    method: 'delete',
    params
  })
}

// @Tags Goods
// @Summary 更新goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Goods true "更新goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goods/updateGoods [put]
export const updateGoods = (data) => {
  return service({
    url: '/goods/updateGoods',
    method: 'put',
    data
  })
}

// @Tags Goods
// @Summary 用id查询goods表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Goods true "用id查询goods表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goods/findGoods [get]
export const findGoods = (params) => {
  return service({
    url: '/goods/findGoods',
    method: 'get',
    params
  })
}

// @Tags Goods
// @Summary 分页获取goods表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取goods表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/getGoodsList [get]
export const getGoodsList = (params) => {
  return service({
    url: '/goods/getGoodsList',
    method: 'get',
    params
  })
}
