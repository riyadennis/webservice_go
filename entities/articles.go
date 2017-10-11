package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"fmt"
)

type Article struct {
	Id          string         `json:"Id"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string      `json:"url"`
	UrlToImage  string      `json:"url_to_image"`
	PublishedAt time.Time   `json:published_at`
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

func (a *Article) Save() (error){
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		return err
	}
	exists, err := client.IndexExists("articles").Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		_, err := client.CreateIndex("articles").BodyString(mapping).Do(ctx)
		if err != nil {
			return err
		}
	}
	put, err := client.Index().
		Index("articles").
		Type("properties").
		Id(a.Id).
		BodyJson(a).
		Do(ctx)
	if err != nil {
		// Handle error
		return err
	}

	fmt.Printf("Indexed articles %s to index %s, type %s\n", put.Id, put.Index, put.Type)
	return nil
}
