package lib

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
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

func ReadFileWriteToKafka(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	msg := new(Message)
	for scanner.Scan() {
		msg.Jsonmsg = scanner.Text()
		msg.Save()
		i++
	}
	return fmt.Sprintf("Line put to kafka %d", i)
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
