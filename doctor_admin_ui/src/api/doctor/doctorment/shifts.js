import service from '@/utils/request'

// @Tags Shift
// @Summary 创建排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shift true "创建排班表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shift/createShift [post]
export const createShift = (data) => {
  return service({
    url: '/shift/createShift',
    method: 'post',
    data
  })
}

// @Tags Shift
// @Summary 删除排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shift true "删除排班表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shift/deleteShift [delete]
export const deleteShift = (params) => {
  return service({
    url: '/shift/deleteShift',
    method: 'delete',
    params
  })
}

// @Tags Shift
// @Summary 批量删除排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除排班表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shift/deleteShift [delete]
export const deleteShiftByIds = (params) => {
  return service({
    url: '/shift/deleteShiftByIds',
    method: 'delete',
    params
  })
}

// @Tags Shift
// @Summary 更新排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shift true "更新排班表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shift/updateShift [put]
export const updateShift = (data) => {
  return service({
    url: '/shift/updateShift',
    method: 'put',
    data
  })
}

// @Tags Shift
// @Summary 用id查询排班表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Shift true "用id查询排班表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shift/findShift [get]
export const findShift = (params) => {
  return service({
    url: '/shift/findShift',
    method: 'get',
    params
  })
}

// @Tags Shift
// @Summary 分页获取排班表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取排班表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shift/getShiftList [get]
export const getShiftList = (params) => {
  return service({
    url: '/shift/getShiftList',
    method: 'get',
    params
  })
}
// @Tags Shift
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shift/findShiftDataSource [get]
export const getShiftDataSource = () => {
return service({
url: '/shift/getShiftDataSource',
method: 'get',
})
}
