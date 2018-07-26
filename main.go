package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	storiesFeedURL = "https://hacker-news.firebaseio.com/v0/%sstories.json"
	oneItemURL     = "https://hacker-news.firebaseio.com/v0/item/%d.json"
	threadURL      = "https://news.ycombinator.com/item?id=%d"
	defaultCount   = 10
	defaultFeed    = "top"
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

func printOne(item map[string]interface{}) {
	score, ok := item["score"].(float64)
	intScore := int(score)
	if ok {
		o := color.New(color.FgHiRed)
		o.Printf("\n(%d) %s\n", intScore, item["title"])
		d := color.New(color.FgCyan, color.Bold)
		url, ok := item["url"]
		if ok == false {
			// This conversion seems shitty, can I do better?
			floatID := item["id"].(float64)
			id := int(floatID)
			url = fmt.Sprintf(threadURL, id)
		}
		d.Printf(" > %s\n", url)
	}
}

func printUsage() {
	fmt.Println("Usage: ./hn <best|top> <count||int>")
}

func main() {
	args := os.Args
	var count int
	var feedType string
	if len(args) > 1 && (args[1] == "-h" || args[1] == "-help") {
		printUsage()
	} else {
		if len(args) == 1 {
			feedType = defaultFeed
			count = defaultCount
		} else if len(args) == 3 {
			feedType = args[1]
			countStr := args[2]
			countInt, err := strconv.Atoi(countStr)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			count = countInt
		} else {
			printUsage()
		}
	}
	storyIDs := fetchStories(feedType)
	for i := 0; i < count; i++ {
		one := fetchOne(storyIDs[i])
		printOne(one)
	}
}
