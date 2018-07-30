package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aybabtme/rgbterm"
)

const (
	threadURL    = "https://news.ycombinator.com/item?id=%d"
	defaultCount = 10
	defaultFeed  = "top"
)

// Helper function to get orange and light white
func getColor(colorName string) (uint8, uint8, uint8) {
	switch colorName {
	case "orange":
		return 254, 152, 36
	case "white":
		return 247, 233, 217
	default:
		return 247, 233, 217
	}
}

func colorWord(word string, colorName string) string {
	r, g, b := getColor(colorName)
	return rgbterm.FgString(word, r, g, b)
}

func printOne(item map[string]interface{}) {
	score, ok := item["score"].(float64)
	intScore := int(score)
	if ok {
		firstLine := fmt.Sprintf("\n(%d) %s\n", intScore, item["title"])
		url, ok := item["url"]
		if ok == false {
			// This conversion seems shitty, can I do better?
			floatID := item["id"].(float64)
			id := int(floatID)
			url = fmt.Sprintf(threadURL, id)
		}
		secondLine := fmt.Sprintf(" > %s\n", url)

		fmt.Printf(colorWord(firstLine, "orange"))
		fmt.Printf(colorWord(secondLine, "white"))
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
			printUsage()
		}
		count = countInt
	}
	return count, feedType
}

func main() {
	count, feedType := parseArgs(os.Args)
	storyIDs := fetchStories(feedType)
	for i := 0; i < count; i++ {
		one := fetchOne(storyIDs[i])
		printOne(one)
	}
}
