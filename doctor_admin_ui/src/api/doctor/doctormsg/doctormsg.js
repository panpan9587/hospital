import service from '@/utils/request'

// @Tags Doctor
// @Summary 创建医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Doctor true "创建医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /doctor/createDoctor [post]
export const createDoctor = (data) => {
  return service({
    url: '/doctor/createDoctor',
    method: 'post',
    data
  })
}

// @Tags Doctor
// @Summary 删除医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Doctor true "删除医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /doctor/deleteDoctor [delete]
export const deleteDoctor = (params) => {
  return service({
    url: '/doctor/deleteDoctor',
    method: 'delete',
    params
  })
}

// @Tags Doctor
// @Summary 批量删除医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /doctor/deleteDoctor [delete]
export const deleteDoctorByIds = (params) => {
  return service({
    url: '/doctor/deleteDoctorByIds',
    method: 'delete',
    params
  })
}

// @Tags Doctor
// @Summary 更新医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Doctor true "更新医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /doctor/updateDoctor [put]
export const updateDoctor = (data) => {
  return service({
    url: '/doctor/updateDoctor',
    method: 'put',
    data
  })
}

// @Tags Doctor
// @Summary 用id查询医生表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Doctor true "用id查询医生表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /doctor/findDoctor [get]
export const findDoctor = (params) => {
  return service({
    url: '/doctor/findDoctor',
    method: 'get',
    params
  })
}

// @Tags Doctor
// @Summary 分页获取医生表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取医生表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /doctor/getDoctorList [get]
export const getDoctorList = (params) => {
  return service({
    url: '/doctor/getDoctorList',
    method: 'get',
    params
  })
}



// @Tags Doctoress
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /doctoress/findDoctoressDataSource [get]
export const getDoctorDataSource = () => {
  return service({
    url: '/doctor/getDoctorDataSource',
    method: 'get',
  })
}
