import service from '@/utils/request'

// @Tags Attendingphysician
// @Summary 创建挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attendingphysician true "创建挂号管理表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /attendingphysician/createAttendingphysician [post]
export const createAttendingphysician = (data) => {
  return service({
    url: '/attendingphysician/createAttendingphysician',
    method: 'post',
    data
  })
}

// @Tags Attendingphysician
// @Summary 删除挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attendingphysician true "删除挂号管理表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /attendingphysician/deleteAttendingphysician [delete]
export const deleteAttendingphysician = (params) => {
  return service({
    url: '/attendingphysician/deleteAttendingphysician',
    method: 'delete',
    params
  })
}

// @Tags Attendingphysician
// @Summary 批量删除挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除挂号管理表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /attendingphysician/deleteAttendingphysician [delete]
export const deleteAttendingphysicianByIds = (params) => {
  return service({
    url: '/attendingphysician/deleteAttendingphysicianByIds',
    method: 'delete',
    params
  })
}

// @Tags Attendingphysician
// @Summary 更新挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attendingphysician true "更新挂号管理表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /attendingphysician/updateAttendingphysician [put]
export const updateAttendingphysician = (data) => {
  return service({
    url: '/attendingphysician/updateAttendingphysician',
    method: 'put',
    data
  })
}

// @Tags Attendingphysician
// @Summary 用id查询挂号管理表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Attendingphysician true "用id查询挂号管理表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /attendingphysician/findAttendingphysician [get]
export const findAttendingphysician = (params) => {
  return service({
    url: '/attendingphysician/findAttendingphysician',
    method: 'get',
    params
  })
}

// @Tags Attendingphysician
// @Summary 分页获取挂号管理表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取挂号管理表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /attendingphysician/getAttendingphysicianList [get]
export const getAttendingphysicianList = (params) => {
  return service({
    url: '/attendingphysician/getAttendingphysicianList',
    method: 'get',
    params
  })
}
// @Tags Attendingphysician
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /attendingphysician/findAttendingphysicianDataSource [get]
export const getAttendingphysicianDataSource = () => {
return service({
url: '/attendingphysician/getAttendingphysicianDataSource',
method: 'get',
})
}
