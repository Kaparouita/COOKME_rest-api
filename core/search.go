package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"rest-api/models"
	"rest-api/repositories"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type SearchService struct {
	ES *elasticsearch.Client
	DB *repositories.Db
}

func NewSearchService(db *repositories.Db) *SearchService {
	rows, err := db.Raw("SELECT * FROM keywords").Rows()
	if err != nil {
		log.Fatal(err)
	}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	createIndex(es, "keywords")
	createIndex(es, "recipes")

	// Prepare bulk request
	var bulkRequest strings.Builder

	for rows.Next() {
		keyword := &models.Keyword{}
		if err := rows.Scan(&keyword.ID, &keyword.Keyword); err != nil {
			log.Fatal(err)
		}

		// Add metadata for each document
		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "keywords", "_id" : "%d" } }%s`, keyword.ID, "\n"))
		data, _ := json.Marshal(keyword)
		bulkRequest.Write(meta)
		bulkRequest.Write(data)
		bulkRequest.WriteString("\n")
	}

	// Perform the bulk request
	res, err := es.Bulk(strings.NewReader(bulkRequest.String()), es.Bulk.WithContext(context.Background()))
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error in response: %s", res.String())
	} else {
		log.Printf("Bulk indexing completed successfully.")
	}
	return &SearchService{
		ES: es,
		DB: db,
	}
}

func createIndex(es *elasticsearch.Client, indexName string) {
	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		log.Fatalf("Error checking if index exists: %s", err)
	}
	if res.StatusCode == 404 {
		// Index does not exist, create it
		mapping := `{
            "mappings": {
                "properties": {
                    "keyword": { "type": "text" }
                }
            }
        }`
		req := esapi.IndicesCreateRequest{
			Index: indexName,
			Body:  strings.NewReader(mapping),
		}
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error creating index: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Fatalf("Error creating index: %s", res.String())
		}
	}
}

func (s *SearchService) SearchKeywords(query string) ([]string, error) {

	var buf bytes.Buffer
	queryJSON := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"name", "keywords", "keyword"},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(queryJSON); err != nil {
		return nil, fmt.Errorf("error encoding query: %s", err)
	}

	res, err := s.ES.Search(
		s.ES.Search.WithContext(context.Background()),
		s.ES.Search.WithIndex("recipes", "keywords"),
		s.ES.Search.WithBody(&buf),
		s.ES.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error in response: %s", res.String())
	}

	searchResponse := &models.SearchResponse{}
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	results := make([]string, len(searchResponse.Hits.Hits))
	for i, hit := range searchResponse.Hits.Hits {
		results[i] = hit.Source["keyword"].(string)
	}

	return results, nil
}

func (s *SearchService) GetAllKeywords() []models.Keyword {
	keywords := s.DB.GetAllKeywords()
	return keywords
}
