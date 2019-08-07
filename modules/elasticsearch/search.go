package elasticsearch

import (
	"chs/common"
	"chs/models"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func Match(alias string, wid int64) (reply *models.ReplyModel) {
	var (
		res       *elastic.SearchResult
		boolQuery *elastic.BoolQuery
		err       error
	)

	boolQuery = elastic.NewBoolQuery()
	boolQuery.Must(
		elastic.NewTermQuery("Disabled", common.NO_VALUE),
		elastic.NewTermQuery("Wid", wid))
	boolQuery.Should(
		elastic.NewTermQuery("ClickKey", alias),
		elastic.NewMatchQuery("Alias", alias))
	res, err = client.Search("replies").Type("reply").Query(boolQuery).Size(1).Pretty(true).Do(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	if res.Hits.TotalHits > 0 {
		err := json.Unmarshal(*res.Hits.Hits[0].Source, reply)
		if err != nil {
			log.Println("es hit json to struct fail: res:%v, err:%v", res, err)
		}
	}
	return
}
