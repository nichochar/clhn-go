package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	storiesFeedURL = "https://hacker-news.firebaseio.com/v0/%sstories.json"
	oneItemURL     = "https://hacker-news.firebaseio.com/v0/item/%d.json"
)

// Returns an array of story IDs
func fetchStories(feedType string) []int {
	url := fmt.Sprintf(storiesFeedURL, feedType)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var m []int
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func fetchOne(item int) map[string]interface{} {
	url := fmt.Sprintf(oneItemURL, item)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}
