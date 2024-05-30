import service from '@/utils/request'

// @Tags Medical
// @Summary 创建病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Medical true "创建病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /medical/createMedical [post]
export const createMedical = (data) => {
  return service({
    url: '/medical/createMedical',
    method: 'post',
    data
  })
}

// @Tags Medical
// @Summary 删除病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Medical true "删除病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /medical/deleteMedical [delete]
export const deleteMedical = (params) => {
  return service({
    url: '/medical/deleteMedical',
    method: 'delete',
    params
  })
}

// @Tags Medical
// @Summary 批量删除病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /medical/deleteMedical [delete]
export const deleteMedicalByIds = (params) => {
  return service({
    url: '/medical/deleteMedicalByIds',
    method: 'delete',
    params
  })
}

// @Tags Medical
// @Summary 更新病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Medical true "更新病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /medical/updateMedical [put]
export const updateMedical = (data) => {
  return service({
    url: '/medical/updateMedical',
    method: 'put',
    data
  })
}

// @Tags Medical
// @Summary 用id查询病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Medical true "用id查询病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /medical/findMedical [get]
export const findMedical = (params) => {
  return service({
    url: '/medical/findMedical',
    method: 'get',
    params
  })
}

// @Tags Medical
// @Summary 分页获取病历表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取病历表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /medical/getMedicalList [get]
export const getMedicalList = (params) => {
  return service({
    url: '/medical/getMedicalList',
    method: 'get',
    params
  })
}
