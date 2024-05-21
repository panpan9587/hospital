// 自动生成模板Goods
package goodspkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// goods表 结构体  Goods
type Goods struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:50;"`  //name字段 
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:;"`  //price字段 
      GoodsType  string `json:"goodsType" form:"goodsType" gorm:"column:goods_type;comment:;size:20;"`  //goodsType字段 
      Cover  string `json:"cover" form:"cover" gorm:"column:cover;comment:;size:150;"`  //cover字段 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;size:1;"`  //status字段 
}


// TableName goods表 Goods自定义表名 goods
func (Goods) TableName() string {
  return "goods"
}

