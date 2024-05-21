package goods

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goods"
    goodsReq "github.com/flipped-aurora/gin-vue-admin/server/model/goods/request"
)

type GoodsService struct {
}

// CreateGoods 创建商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *GoodsService) CreateGoods(product *goods.Goods) (err error) {
	err = global.GVA_DB.Create(product).Error
	return err
}

// DeleteGoods 删除商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *GoodsService)DeleteGoods(ID string) (err error) {
	err = global.GVA_DB.Delete(&goods.Goods{},"id = ?",ID).Error
	return err
}

// DeleteGoodsByIds 批量删除商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *GoodsService)DeleteGoodsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]goods.Goods{},"id in ?",IDs).Error
	return err
}

// UpdateGoods 更新商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *GoodsService)UpdateGoods(product goods.Goods) (err error) {
	err = global.GVA_DB.Save(&product).Error
	return err
}

// GetGoods 根据ID获取商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *GoodsService)GetGoods(ID string) (product goods.Goods, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&product).Error
	return
}

// GetGoodsInfoList 分页获取商品记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *GoodsService)GetGoodsInfoList(info goodsReq.GoodsSearch) (list []goods.Goods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&goods.Goods{})
    var products []goods.Goods
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.GoodsType != "" {
        db = db.Where("goods_type LIKE ?","%"+ info.GoodsType+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&products).Error
	return  products, total, err
}
