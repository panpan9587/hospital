package goodspkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goodspkg"
    goodspkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/goodspkg/request"
)

type GoodsService struct {
}

// CreateGoods 创建goods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService) CreateGoods(goods *goodspkg.Goods) (err error) {
	err = global.GVA_DB.Create(goods).Error
	return err
}

// DeleteGoods 删除goods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)DeleteGoods(ID string) (err error) {
	err = global.GVA_DB.Delete(&goodspkg.Goods{},"id = ?",ID).Error
	return err
}

// DeleteGoodsByIds 批量删除goods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)DeleteGoodsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]goodspkg.Goods{},"id in ?",IDs).Error
	return err
}

// UpdateGoods 更新goods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)UpdateGoods(goods goodspkg.Goods) (err error) {
	err = global.GVA_DB.Save(&goods).Error
	return err
}

// GetGoods 根据ID获取goods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)GetGoods(ID string) (goods goodspkg.Goods, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&goods).Error
	return
}

// GetGoodsInfoList 分页获取goods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (goodsService *GoodsService)GetGoodsInfoList(info goodspkgReq.GoodsSearch) (list []goodspkg.Goods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&goodspkg.Goods{})
    var goodss []goodspkg.Goods
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&goodss).Error
	return  goodss, total, err
}
