import service from '@/utils/request'

// @Tags Health
// @Summary 创建体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Health true "创建体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /health/createHealth [post]
export const createHealth = (data) => {
  return service({
    url: '/health/createHealth',
    method: 'post',
    data
  })
}

// @Tags Health
// @Summary 删除体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Health true "删除体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /health/deleteHealth [delete]
export const deleteHealth = (params) => {
  return service({
    url: '/health/deleteHealth',
    method: 'delete',
    params
  })
}

// @Tags Health
// @Summary 批量删除体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /health/deleteHealth [delete]
export const deleteHealthByIds = (params) => {
  return service({
    url: '/health/deleteHealthByIds',
    method: 'delete',
    params
  })
}

// @Tags Health
// @Summary 更新体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Health true "更新体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /health/updateHealth [put]
export const updateHealth = (data) => {
  return service({
    url: '/health/updateHealth',
    method: 'put',
    data
  })
}

// @Tags Health
// @Summary 用id查询体检表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Health true "用id查询体检表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /health/findHealth [get]
export const findHealth = (params) => {
  return service({
    url: '/health/findHealth',
    method: 'get',
    params
  })
}

// @Tags Health
// @Summary 分页获取体检表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取体检表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /health/getHealthList [get]
export const getHealthList = (params) => {
  return service({
    url: '/health/getHealthList',
    method: 'get',
    params
  })
}
