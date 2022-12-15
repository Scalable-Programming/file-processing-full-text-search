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
}

func IndexFullFileText(id string, text *string) {
	data, err := json.Marshal(EsFullTextData{FileId: id, Text: *text})

	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}

	req := esapi.IndexRequest{
		Index:   fullTextSeachIndex,
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
	}
}

func Search(text string) []string {
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
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(fullTextSeachIndex),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	defer res.Body.Close()

	// Check for hits and make sure they actually exists
	hits, ok := r["hits"].(map[string]interface{})

	if !ok {
		return searchResultIds
	}

	for _, hit := range hits["hits"].([]interface{}) {
		fileId := hit.(map[string]interface{})["_source"].(map[string]interface{})["FileId"].(string)
		searchResultIds = append(searchResultIds, fileId)
	}

	return searchResultIds
}
