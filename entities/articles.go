package entities

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

type Article struct {
	Id          string    `json:"Id"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	UrlToImage  string    `json:"url_to_image"`
	PublishedAt time.Time `json:published_at`
}

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

const mapping = `{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"articles":{
			"properties": {
				"Id":{
					"type":"keyword"
				},
				"author":{
					"type":"text"
				},
				"title":{
					"type":"text"
				},
				"description":{
					"type":"text"
				},
				"url":{
					"type":"text"
				},
				"url_to_image":{
					"type":"text"
				},
				"published_at":{
					"type":"text"
				}
			}
		}
	}
}`

func (a *Article) Save() error {
	client, err := elastic.NewDefaultClient()
	if err != nil {
		return err
	}

	resp, err := client.Search(
		client.Search.WithIndex("articles"),
		client.Search.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to connect to elasticsearch")
	}

	var responseBody []byte
	_, err = resp.Body.Read(responseBody)
	if err != nil {
		// Handle error
		return err
	}

	fmt.Printf("Indexed articles %s to index %s, type %s\n", string(responseBody))
	return nil
}
