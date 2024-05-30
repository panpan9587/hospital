import service from '@/utils/request'

// @Tags UserAuthS
// @Summary 创建userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuthS true "创建userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userAuth/createUserAuthS [post]
export const createUserAuthS = (data) => {
  return service({
    url: '/userAuth/createUserAuthS',
    method: 'post',
    data
  })
}

// @Tags UserAuthS
// @Summary 删除userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuthS true "删除userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAuth/deleteUserAuthS [delete]
export const deleteUserAuthS = (params) => {
  return service({
    url: '/userAuth/deleteUserAuthS',
    method: 'delete',
    params
  })
}

// @Tags UserAuthS
// @Summary 批量删除userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAuth/deleteUserAuthS [delete]
export const deleteUserAuthSByIds = (params) => {
  return service({
    url: '/userAuth/deleteUserAuthSByIds',
    method: 'delete',
    params
  })
}

// @Tags UserAuthS
// @Summary 更新userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuthS true "更新userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAuth/updateUserAuthS [put]
export const updateUserAuthS = (data) => {
  return service({
    url: '/userAuth/updateUserAuthS',
    method: 'put',
    data
  })
}

// @Tags UserAuthS
// @Summary 用id查询userAuth表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserAuthS true "用id查询userAuth表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAuth/findUserAuthS [get]
export const findUserAuthS = (params) => {
  return service({
    url: '/userAuth/findUserAuthS',
    method: 'get',
    params
  })
}

// @Tags UserAuthS
// @Summary 分页获取userAuth表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userAuth表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAuth/getUserAuthSList [get]
export const getUserAuthSList = (params) => {
  return service({
    url: '/userAuth/getUserAuthSList',
    method: 'get',
    params
  })
}
