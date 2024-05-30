package doctorrecipe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/doctorrecipe"
	doctorrecipeReq "github.com/flipped-aurora/gin-vue-admin/server/model/doctorrecipe/request"
)

type RecipeService struct {
}

// CreateRecipe 创建处方表记录
// Author [piexlmax](https://github.com/piexlmax)
func (recipeService *RecipeService) CreateRecipe(recipe *doctorrecipe.Recipe) (err error) {
	err = global.GVA_DB.Create(recipe).Error
	return err
}

// DeleteRecipe 删除处方表记录
// Author [piexlmax](https://github.com/piexlmax)
func (recipeService *RecipeService) DeleteRecipe(ID string) (err error) {
	err = global.GVA_DB.Delete(&doctorrecipe.Recipe{}, "id = ?", ID).Error
	return err
}

// DeleteRecipeByIds 批量删除处方表记录
// Author [piexlmax](https://github.com/piexlmax)
func (recipeService *RecipeService) DeleteRecipeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]doctorrecipe.Recipe{}, "id in ?", IDs).Error
	return err
}

// UpdateRecipe 更新处方表记录
// Author [piexlmax](https://github.com/piexlmax)
func (recipeService *RecipeService) UpdateRecipe(recipe doctorrecipe.Recipe) (err error) {
	err = global.GVA_DB.Model(&doctorrecipe.Recipe{}).Where("id = ?", recipe.ID).Updates(&recipe).Error
	return err
}

// GetRecipe 根据ID获取处方表记录
// Author [piexlmax](https://github.com/piexlmax)
func (recipeService *RecipeService) GetRecipe(ID string) (recipe doctorrecipe.Recipe, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&recipe).Error
	return
}

// GetRecipeInfoList 分页获取处方表记录
// Author [piexlmax](https://github.com/piexlmax)
func (recipeService *RecipeService) GetRecipeInfoList(info doctorrecipeReq.RecipeSearch) (list []doctorrecipe.Recipe, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&doctorrecipe.Recipe{})
	var recipes []doctorrecipe.Recipe
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&recipes).Error
	return recipes, total, err
}
func (recipeService *RecipeService) GetRecipeDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("user").Select("full_name as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}
