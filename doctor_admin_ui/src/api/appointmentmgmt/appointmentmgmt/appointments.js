import service from '@/utils/request'

// @Tags Appointment
// @Summary 创建预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Appointment true "创建预约记录表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /appointment/createAppointment [post]
export const createAppointment = (data) => {
  return service({
    url: '/appointment/createAppointment',
    method: 'post',
    data
  })
}

// @Tags Appointment
// @Summary 删除预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Appointment true "删除预约记录表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appointment/deleteAppointment [delete]
export const deleteAppointment = (params) => {
  return service({
    url: '/appointment/deleteAppointment',
    method: 'delete',
    params
  })
}

// @Tags Appointment
// @Summary 批量删除预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除预约记录表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appointment/deleteAppointment [delete]
export const deleteAppointmentByIds = (params) => {
  return service({
    url: '/appointment/deleteAppointmentByIds',
    method: 'delete',
    params
  })
}

// @Tags Appointment
// @Summary 更新预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Appointment true "更新预约记录表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appointment/updateAppointment [put]
export const updateAppointment = (data) => {
  return service({
    url: '/appointment/updateAppointment',
    method: 'put',
    data
  })
}

// @Tags Appointment
// @Summary 用id查询预约记录表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Appointment true "用id查询预约记录表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appointment/findAppointment [get]
export const findAppointment = (params) => {
  return service({
    url: '/appointment/findAppointment',
    method: 'get',
    params
  })
}

// @Tags Appointment
// @Summary 分页获取预约记录表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取预约记录表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appointment/getAppointmentList [get]
export const getAppointmentList = (params) => {
  return service({
    url: '/appointment/getAppointmentList',
    method: 'get',
    params
  })
}
// @Tags Appointment
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appointment/findAppointmentDataSource [get]
export const getAppointmentDataSource = () => {
return service({
url: '/appointment/getAppointmentDataSource',
method: 'get',
})
}
