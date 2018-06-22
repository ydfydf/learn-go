package controller

import (
	"learngo/crawler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strings"
	"strconv"
	"learngo/crawler/frontend/model"
	"context"
	"reflect"
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandle(template string) SearchResultHandler {
	client , err := elastic.NewClient(
		elastic.SetURL(config.ElaticUrl),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page,err := h.getSearchResult(q,from)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
	err = h.view.Render(w,page)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string ,from int) (model.SearchResult,error) {
	var result model.SearchResult
	resp ,err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).//elastic库里有很多NewQuery相关的方法，这里只做最简单的
		From(from).
		Do(context.Background())
	if err != nil {
		return result,err
	}
	result.Hits = int(resp.TotalHits())
	result.Start = from
	//func (r *SearchResult) Each(typ reflect.Type) []interface{}，go语言中也有反射
	for _,v := range resp.Each(reflect.TypeOf(engine.Item{})){//反射
		item := v.(engine.Item)//类型判断，当v的值的类型为engine.Item时，进行赋值
		result.Items = append(result.Items,item)
	}
	return result,nil
}

