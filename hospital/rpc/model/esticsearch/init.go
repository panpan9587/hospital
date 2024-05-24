package esticsearch

import (
	"context"
	"demo/config"
	"encoding/json"
	"fmt"
	"math"

	"go.uber.org/zap"

	"github.com/olivere/elastic/v7"
)

var (
	ES  *elastic.Client
	err error
)

func init() {
	dsn := fmt.Sprintf("http://%s:%d", config.GlobalConfig.Es.Host, config.GlobalConfig.Es.Port)
	ES, err = elastic.NewClient(elastic.SetURL(dsn), elastic.SetSniff(false))
	if err != nil {
		zap.S().Info("es连接失败的原因:::", err.Error())
		return
	}

}

// 查询所有+分页
func QueryEs(index string, query elastic.Query, page int) (int, int, []map[string]interface{}, error) {
	//分页

	//当前页
	if page == 0 {
		page = 1
	}

	//每页展示条数
	size := 1
	//总条数
	count, _ := ES.Count().Index(index).Query(query).Do(context.Background())
	//总页数
	totalPage := math.Ceil(float64(int(count) / size))

	res, err := ES.Search().Index(index).Query(query).From((page - 1) * size).Size(size).Do(context.Background())
	if err != nil {
		return 0, 0, nil, err
	}
	var list []map[string]interface{}
	for _, hit := range res.Hits.Hits {
		data := make(map[string]interface{})
		err = json.Unmarshal(hit.Source, &data)
		if err != nil {
			return 0, 0, nil, err
		}
		list = append(list, data)
	}
	return int(totalPage), page, list, nil
}

// 单条件查询
func FilterSearch(index string, content string) ([]map[string]interface{}, error) {
	res, err := ES.Search().Index(index).Query(elastic.NewQueryStringQuery(content)).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var list []map[string]interface{}
	for _, hit := range res.Hits.Hits {
		data := make(map[string]interface{})
		err = json.Unmarshal(hit.Source, &data)
		if err != nil {
			return nil, err
		}
		list = append(list, data)
	}
	return list, nil
}
