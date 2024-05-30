import service from '@/utils/request'

// @Tags UserAuth
// @Summary 创建userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuth true "创建userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userAuth/createUserAuth [post]
export const createUserAuth = (data) => {
  return service({
    url: '/userAuth/createUserAuth',
    method: 'post',
    data
  })
}

// @Tags UserAuth
// @Summary 删除userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuth true "删除userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAuth/deleteUserAuth [delete]
export const deleteUserAuth = (params) => {
  return service({
    url: '/userAuth/deleteUserAuth',
    method: 'delete',
    params
  })
}

// @Tags UserAuth
// @Summary 批量删除userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAuth/deleteUserAuth [delete]
export const deleteUserAuthByIds = (params) => {
  return service({
    url: '/userAuth/deleteUserAuthByIds',
    method: 'delete',
    params
  })
}

// @Tags UserAuth
// @Summary 更新userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuth true "更新userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAuth/updateUserAuth [put]
export const updateUserAuth = (data) => {
  return service({
    url: '/userAuth/updateUserAuth',
    method: 'put',
    data
  })
}

// @Tags UserAuth
// @Summary 用id查询userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserAuth true "用id查询userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAuth/findUserAuth [get]
export const findUserAuth = (params) => {
  return service({
    url: '/userAuth/findUserAuth',
    method: 'get',
    params
  })
}

// @Tags UserAuth
// @Summary 分页获取userAuth表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userAuth表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAuth/getUserAuthList [get]
export const getUserAuthList = (params) => {
  return service({
    url: '/userAuth/getUserAuthList',
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
export const getUserAuthDataSource = () => {
  return service({
    url: '/userAuth/getUserAuthDataSource',
    method: 'get',
  })
}
