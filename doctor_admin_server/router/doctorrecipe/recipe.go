package doctorrecipe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RecipeRouter struct {
}

// InitRecipeRouter 初始化 处方表 路由信息
func (s *RecipeRouter) InitRecipeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	recipeRouter := Router.Group("recipe").Use(middleware.OperationRecord())
	recipeRouterWithoutRecord := Router.Group("recipe")
	recipeRouterWithoutAuth := PublicRouter.Group("recipe")

	var recipeApi = v1.ApiGroupApp.DoctorrecipeApiGroup.RecipeApi
	{
		recipeRouter.POST("createRecipe", recipeApi.CreateRecipe)             // 新建处方表
		recipeRouter.DELETE("deleteRecipe", recipeApi.DeleteRecipe)           // 删除处方表
		recipeRouter.DELETE("deleteRecipeByIds", recipeApi.DeleteRecipeByIds) // 批量删除处方表
		recipeRouter.PUT("updateRecipe", recipeApi.UpdateRecipe)              // 更新处方表
	}
	{
		recipeRouterWithoutRecord.GET("findRecipe", recipeApi.FindRecipe)       // 根据ID获取处方表
		recipeRouterWithoutRecord.GET("getRecipeList", recipeApi.GetRecipeList) // 获取处方表列表
	}
	{
		recipeRouterWithoutAuth.GET("getRecipeDataSource", recipeApi.GetRecipeDataSource) // 获取处方表数据源
		recipeRouterWithoutAuth.GET("getRecipePublic", recipeApi.GetRecipePublic)         // 获取处方表列表
	}
}
