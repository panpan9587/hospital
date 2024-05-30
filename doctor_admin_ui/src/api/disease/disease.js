import service from '@/utils/request'

// @Tags Diseases
// @Summary 创建病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Diseases true "创建病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /diseases/createDiseases [post]
export const createDiseases = (data) => {
  return service({
    url: '/diseases/createDiseases',
    method: 'post',
    data
  })
}

// @Tags Diseases
// @Summary 删除病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Diseases true "删除病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /diseases/deleteDiseases [delete]
export const deleteDiseases = (params) => {
  return service({
    url: '/diseases/deleteDiseases',
    method: 'delete',
    params
  })
}

// @Tags Diseases
// @Summary 批量删除病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /diseases/deleteDiseases [delete]
export const deleteDiseasesByIds = (params) => {
  return service({
    url: '/diseases/deleteDiseasesByIds',
    method: 'delete',
    params
  })
}

// @Tags Diseases
// @Summary 更新病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Diseases true "更新病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /diseases/updateDiseases [put]
export const updateDiseases = (data) => {
  return service({
    url: '/diseases/updateDiseases',
    method: 'put',
    data
  })
}

// @Tags Diseases
// @Summary 用id查询病历表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Diseases true "用id查询病历表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /diseases/findDiseases [get]
export const findDiseases = (params) => {
  return service({
    url: '/diseases/findDiseases',
    method: 'get',
    params
  })
}

// @Tags Diseases
// @Summary 分页获取病历表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取病历表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /diseases/getDiseasesList [get]
export const getDiseasesList = (params) => {
  return service({
    url: '/diseases/getDiseasesList',
    method: 'get',
    params
  })
}
