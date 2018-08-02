package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const (
	threadURL      = "https://news.ycombinator.com/item?id=%d"
	storiesFeedURL = "https://hacker-news.firebaseio.com/v0/%sstories.json"
	oneItemURL     = "https://hacker-news.firebaseio.com/v0/item/%d.json"
)

var mux sync.Mutex
var results = make(map[int]Story)
var done = make(chan bool)

type Story struct {
	id    int
	url   string
	title string
	votes int
}

// Fetch the list of ids
func fetchStoryIDs(feedType string) []int {
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

func makeStory(body map[string]interface{}) Story {
	id := int(body["id"].(float64))
	url, err := body["url"].(string)
	if err == false {
		url = fmt.Sprintf(threadURL, id)
	}
	votes := int(body["score"].(float64))
	title := body["title"].(string)
	return Story{id: id, url: url, title: title, votes: votes}
}

func fetchOne(item int, ch chan int) {
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
	story := makeStory(m)
	mux.Lock()
	results[item] = story
	mux.Unlock()
	// Increment the queue
	ch <- item
	return
}

// Returns an array of Story structs
// Fetches them in parallel
func fetchStories(count int, feedType string) []Story {
	var stories []Story
	var feed = fetchStoryIDs(feedType)
	var ch = make(chan int, count)

	for i := 0; i < count; i++ {
		go fetchOne(feed[i], ch)
	}

	for j := 0; j < count; j++ {
		<-ch
	}

	for k := 0; k < count; k++ {
		stories = append(stories, results[feed[k]])
	}

	return stories

}
