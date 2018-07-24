package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const topStoriesURL = "https://hacker-news.firebaseio.com/v0/topstories.json"
const oneItemURL = "https://hacker-news.firebaseio.com/v0/item/%d.json"

// Returns an array of story IDs
func fetchStories(count int) []int {
	resp, err := http.Get(topStoriesURL)
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

func fetchOne(item int) interface{} {
	url := fmt.Sprintf(oneItemURL, item)
	fmt.Printf("Fetching the following URL:\n%s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var m interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func main() {
	storyIDs := fetchStories(10)
	storyIDs = storyIDs[0:5]
	for i := 0; i < 5; i++ {
		one := fetchOne(storyIDs[i])
		fmt.Printf("%v\n", one)

	}
}
