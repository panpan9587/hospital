// 自动生成模板Recipe
package doctorrecipe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 处方表 结构体  Recipe
type Recipe struct {
	global.GVA_MODEL
	UserId     *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:就诊人id;" binding:"required"`                      //就诊人
	Drug       string   `json:"drug" form:"drug" gorm:"column:drug;comment:处方用药;size:50;" binding:"required"`                      //处方用药
	Price      *float64 `json:"price" form:"price" gorm:"column:price;comment:处方价格;" binding:"required"`                           //处方价格
	Prescriber string   `json:"prescriber" form:"prescriber" gorm:"column:prescriber;comment:开药医生;size:20;" binding:"required"`    //开药医生
	Status     string   `json:"status" form:"status" gorm:"column:status;comment:状态（待审核、审核中、已审核、已驳回）;size:10;" binding:"required"` //状态
}

// TableName 处方表 Recipe自定义表名 recipe
func (Recipe) TableName() string {
	return "recipe"
}
