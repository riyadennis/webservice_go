package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type item struct {
	Title   string
	Url     string
	Comment int `json:"num_comments"`
}

type Response struct {
	Data struct {
		Children []struct {
			Data item
		}
	}
}

func (i item) String() string {
	return fmt.Sprintf("Title %s Url %s ( Comments %d )", i.Title, i.Url, i.Comment)
}

func ReadReddit(topic string) ([]item, error) {
	url := fmt.Sprintf("https://www.reddit.com/r/%s.json", topic)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	items := make([]item, len(r.Data.Children))
	for i, item := range r.Data.Children {
		items[i] = item.Data
	}
	return items, nil
}
