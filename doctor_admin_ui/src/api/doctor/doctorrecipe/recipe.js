import service from '@/utils/request'

// @Tags Recipe
// @Summary 创建处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recipe true "创建处方表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /recipe/createRecipe [post]
export const createRecipe = (data) => {
  return service({
    url: '/recipe/createRecipe',
    method: 'post',
    data
  })
}

// @Tags Recipe
// @Summary 删除处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recipe true "删除处方表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recipe/deleteRecipe [delete]
export const deleteRecipe = (params) => {
  return service({
    url: '/recipe/deleteRecipe',
    method: 'delete',
    params
  })
}

// @Tags Recipe
// @Summary 批量删除处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除处方表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recipe/deleteRecipe [delete]
export const deleteRecipeByIds = (params) => {
  return service({
    url: '/recipe/deleteRecipeByIds',
    method: 'delete',
    params
  })
}

// @Tags Recipe
// @Summary 更新处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recipe true "更新处方表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recipe/updateRecipe [put]
export const updateRecipe = (data) => {
  return service({
    url: '/recipe/updateRecipe',
    method: 'put',
    data
  })
}

// @Tags Recipe
// @Summary 用id查询处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Recipe true "用id查询处方表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recipe/findRecipe [get]
export const findRecipe = (params) => {
  return service({
    url: '/recipe/findRecipe',
    method: 'get',
    params
  })
}

// @Tags Recipe
// @Summary 分页获取处方表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取处方表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recipe/getRecipeList [get]
export const getRecipeList = (params) => {
  return service({
    url: '/recipe/getRecipeList',
    method: 'get',
    params
  })
}
// @Tags Recipe
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recipe/findRecipeDataSource [get]
export const getRecipeDataSource = () => {
return service({
url: '/recipe/getRecipeDataSource',
method: 'get',
})
}
