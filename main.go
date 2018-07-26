package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	topStoriesURL = "https://hacker-news.firebaseio.com/v0/topstories.json"
	oneItemURL    = "https://hacker-news.firebaseio.com/v0/item/%d.json"
)

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

func printOne(item map[string]interface{}) {
	switch itype := item["type"]; itype {
	case "story":
		score, ok := item["score"].(float64)
		intScore := int(score)
		if ok {
			o := color.New(color.FgHiRed)
			o.Printf("\n(%d) %s\n", intScore, item["title"])
			d := color.New(color.FgCyan, color.Bold)
			d.Printf(" > %s\n", item["url"])
		}
	case "comment":
		fmt.Println("It's a comment! Skipping...")
	case "job":
		fmt.Println("It's a job! Skipping...")
	case "poll":
		fmt.Println("It's a poll! Skipping...")
	case "pollopt":
		fmt.Println("It's a pollopt! Skipping...")
	}
}

func main() {
	storyIDs := fetchStories(10)
	storyIDs = storyIDs[0:5]
	for i := 0; i < 5; i++ {
		one := fetchOne(storyIDs[i])
		printOne(one)

	}
}
