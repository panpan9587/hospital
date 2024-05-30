package doctorrecipe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/doctorrecipe"
	doctorrecipeReq "github.com/flipped-aurora/gin-vue-admin/server/model/doctorrecipe/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RecipeApi struct {
}

var recipeService = service.ServiceGroupApp.DoctorrecipeServiceGroup.RecipeService

// CreateRecipe 创建处方表
// @Tags Recipe
// @Summary 创建处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctorrecipe.Recipe true "创建处方表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /recipe/createRecipe [post]
func (recipeApi *RecipeApi) CreateRecipe(c *gin.Context) {
	var recipe doctorrecipe.Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := recipeService.CreateRecipe(&recipe); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRecipe 删除处方表
// @Tags Recipe
// @Summary 删除处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctorrecipe.Recipe true "删除处方表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /recipe/deleteRecipe [delete]
func (recipeApi *RecipeApi) DeleteRecipe(c *gin.Context) {
	ID := c.Query("ID")
	if err := recipeService.DeleteRecipe(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRecipeByIds 批量删除处方表
// @Tags Recipe
// @Summary 批量删除处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /recipe/deleteRecipeByIds [delete]
func (recipeApi *RecipeApi) DeleteRecipeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := recipeService.DeleteRecipeByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRecipe 更新处方表
// @Tags Recipe
// @Summary 更新处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body doctorrecipe.Recipe true "更新处方表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /recipe/updateRecipe [put]
func (recipeApi *RecipeApi) UpdateRecipe(c *gin.Context) {
	var recipe doctorrecipe.Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := recipeService.UpdateRecipe(recipe); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRecipe 用id查询处方表
// @Tags Recipe
// @Summary 用id查询处方表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query doctorrecipe.Recipe true "用id查询处方表"
// @Success 200 {object} response.Response{data=object{rerecipe=doctorrecipe.Recipe},msg=string} "查询成功"
// @Router /recipe/findRecipe [get]
func (recipeApi *RecipeApi) FindRecipe(c *gin.Context) {
	ID := c.Query("ID")
	if rerecipe, err := recipeService.GetRecipe(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerecipe": rerecipe}, c)
	}
}

// GetRecipeList 分页获取处方表列表
// @Tags Recipe
// @Summary 分页获取处方表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query doctorrecipeReq.RecipeSearch true "分页获取处方表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /recipe/getRecipeList [get]
func (recipeApi *RecipeApi) GetRecipeList(c *gin.Context) {
	var pageInfo doctorrecipeReq.RecipeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := recipeService.GetRecipeInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetRecipeDataSource 获取Recipe的数据源
// @Tags Recipe
// @Summary 获取Recipe的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /recipe/getRecipeDataSource [get]
func (recipeApi *RecipeApi) GetRecipeDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := recipeService.GetRecipeDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetRecipePublic 不需要鉴权的处方表接口
// @Tags Recipe
// @Summary 不需要鉴权的处方表接口
// @accept application/json
// @Produce application/json
// @Param data query doctorrecipeReq.RecipeSearch true "分页获取处方表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /recipe/getRecipePublic [get]
func (recipeApi *RecipeApi) GetRecipePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的处方表接口信息",
	}, "获取成功", c)
}
