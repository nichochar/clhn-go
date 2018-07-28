package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

const (
	threadURL    = "https://news.ycombinator.com/item?id=%d"
	defaultCount = 10
	defaultFeed  = "top"
)

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
	fmt.Println("Usage:\n  $ ./hn [best|top] [count]")
	fmt.Println("Examples:\n  $ ./hn\n  $ ./hn top 5")
	os.Exit(0)
}

// Parse incoming arguments, and return (count, feedType)
func parseArgs(args []string) (int, string) {
	var count int
	var feedType string
	switch len(args) {
	case 1:
		feedType = defaultFeed
		count = defaultCount
	case 2:
		// if err is nil this is castable into an int
		potentialCount, err := strconv.Atoi(args[1])
		if args[1] == "best" || args[1] == "top" {
			count = defaultCount
			feedType = args[1]
		} else if err == nil {
			count = potentialCount
			feedType = defaultFeed
		} else {
			printUsage()
		}
	case 3:
		feedType = args[1]
		countStr := args[2]
		countInt, err := strconv.Atoi(countStr)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		count = countInt
	}
	return count, feedType
}

func main() {
	args := os.Args
	var count int
	var feedType string
	if len(args) > 1 && (args[1] == "-h" || args[1] == "-help") {
		printUsage()
		os.Exit(0)
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
