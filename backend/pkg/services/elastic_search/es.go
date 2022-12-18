package elastic_search

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type EsFullTextData struct {
	FileId string
	Text   string
}

const fullTextSeachIndex = "fulltextsearch"

var esClient *elasticsearch.Client

func Connect() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	esClient = es

	createIndex()
}

func createIndex() {
	indexExistRequest := esapi.IndicesExistsRequest{
		Index: []string{fullTextSeachIndex},
	}

	indexExistResponse, err := indexExistRequest.Do(context.Background(), esClient)

	if err != nil {
		log.Fatalf("Error checking for index: %s", err)
	}

	if indexExistResponse.StatusCode == 200 {
		return
	}

	indexReq := esapi.IndicesCreateRequest{
		Index: fullTextSeachIndex,
	}

	_, err = indexReq.Do(context.Background(), esClient)

	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
}

func IndexFullFileText(id string, text *string) {
	data, err := json.Marshal(EsFullTextData{FileId: id, Text: *text})

	if err != nil {
		log.Print("Error marshaling document", err)
		return
	}

	req := esapi.IndexRequest{
		Index:   fullTextSeachIndex,
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), esClient)
	defer res.Body.Close()

	if err != nil || res.IsError() {
		log.Print("Error indexing document", err)
		return
	}
}

func Search(text string) ([]string, error) {
	var (
		r               map[string]interface{}
		buf             bytes.Buffer
		searchResultIds []string
	)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Text": text,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(fullTextSeachIndex),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Check for hits and make sure they actually exists
	hits, ok := r["hits"].(map[string]interface{})

	if !ok {
		return searchResultIds, nil
	}

	for _, hit := range hits["hits"].([]interface{}) {
		fileId := hit.(map[string]interface{})["_source"].(map[string]interface{})["FileId"].(string)
		searchResultIds = append(searchResultIds, fileId)
	}

	return searchResultIds, nil
}

func GetFileText(fileId string) (*string, error) {
	var (
		r   map[string]interface{}
		buf bytes.Buffer
	)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"FileId": fileId,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(fullTextSeachIndex),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	hits, ok := r["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, nil
	}

	text, ok := hits[0].(map[string]interface{})["_source"].(map[string]interface{})["Text"].(string)

	if !ok {
		return nil, nil
	}

	return &text, err

}
