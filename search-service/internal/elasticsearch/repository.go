package es

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
)

type SearchResponse[T any] struct {
	Documents []T
	Total     int64
	Highlight map[string]string
	Error     error
}

func GetDocument[T any](indexKey, id string) (T, error) {
	es := GetConnection() //TODO: this is singleton connection

	var record T
	if indexKey == "" {
		return record, errors.New("empty index")
	}

	res, err := es.Client.Get(indexKey, id).Do(context.Background())

	b, err := json.Marshal(res.Source_)

	if err != nil {
		return record, err
	}

	json.Unmarshal(b, &record)

	return record, nil
}

func Search[T any](indexKey string, query *search.Request) SearchResponse[T] {
	es := GetConnection() //TODO: this is singleton connection

	var records []T
	var total int64 = 0
	highlight := make(map[string]string)

	res, err := es.Client.Search().Index(indexKey).Request(query).Do(context.Background())

	if err != nil {
		return SearchResponse[T]{
			Documents: []T{},
			Total:     0,
			Highlight: map[string]string{},
			Error:     err,
		}
	}

	total = res.Hits.Total.Value

	for _, data := range res.Hits.Hits {
		var record T

		json.Unmarshal(data.Source_, &record)

		for _, values := range data.Highlight {
			highlight[data.Id_] = values[0]
			break
		}

		records = append(records, record)
	}

	return SearchResponse[T]{
		Documents: records,
		Total:     total,
		Highlight: highlight,
		Error:     nil,
	}
}

func CreateIndex(indexName string, analyze *create.Request) (bool, error) {
	es := GetConnection() //TODO: this is singleton connection

	exists, _ := es.Client.Indices.Exists(indexName).IsSuccess(context.Background())

	if exists {
		return true, nil
	}

	_, err := es.Client.Indices.Create(indexName).Request(analyze).Do(context.Background())

	if err != nil {
		panic(err)
	}

	return true, err
}

func IndexDocument(indexKey, id string, document interface{}) (*index.Response, error) {
	es := GetConnection()

	if indexKey == "" {
		return nil, errors.New("empty index")
	}

	indices := es.Client.Index(indexKey)

	if id != "" {
		indices = indices.Id(id)
	}

	res, err := indices.Request(document).Do(context.Background())

	return res, err
}

func UpdateDocument(indexKey, id string, document *update.Request) (*update.Response, error) {
	es := GetConnection()

	res, err := es.Client.Update(indexKey, id).Request(document).Do(context.Background())

	return res, err
}

func DeleteDocument(indexKey, id string) (bool, error) {
	es := GetConnection()

	_, err := es.Client.Delete(indexKey, id).Do(context.Background())

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetOneDocument(indexKey string, query *search.Request) (*search.Response, error) {
	es := GetConnection()

	res, err := es.Client.Search().Index(indexKey).Request(query).Do(context.Background())

	if err != nil {
		return nil, err
	}

	return res, nil
}
