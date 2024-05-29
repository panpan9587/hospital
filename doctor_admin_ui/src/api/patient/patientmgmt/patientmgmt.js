import service from '@/utils/request'

// @Tags User
// @Summary 创建患者表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User true "创建患者表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /user/createUser [post]
export const createUser = (data) => {
  return service({
    url: '/user/createUser',
    method: 'post',
    data
  })
}

// @Tags User
// @Summary 删除患者表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User true "删除患者表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (params) => {
  return service({
    url: '/user/deleteUser',
    method: 'delete',
    params
  })
}

// @Tags User
// @Summary 批量删除患者表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除患者表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
export const deleteUserByIds = (params) => {
  return service({
    url: '/user/deleteUserByIds',
    method: 'delete',
    params
  })
}

// @Tags User
// @Summary 更新患者表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User true "更新患者表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateUser [put]
export const updateUser = (data) => {
  return service({
    url: '/user/updateUser',
    method: 'put',
    data
  })
}

// @Tags User
// @Summary 用id查询患者表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.User true "用id查询患者表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findUser [get]
export const findUser = (params) => {
  return service({
    url: '/user/findUser',
    method: 'get',
    params
  })
}

// @Tags User
// @Summary 分页获取患者表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取患者表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
export const getUserList = (params) => {
  return service({
    url: '/user/getUserList',
    method: 'get',
    params
  })
}
