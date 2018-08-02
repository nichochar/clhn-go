package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aybabtme/rgbterm"
)

const (
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

func printOne(story Story) {
	firstLine := fmt.Sprintf("\n(%d) %s\n", story.votes, story.title)
	secondLine := fmt.Sprintf(" > %s\n", story.url)

	fmt.Printf(colorWord(firstLine, "orange"))
	fmt.Printf(colorWord(secondLine, "white"))
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
	var stories []Story
	stories = fetchStories(count, feedType)

	for i := 0; i < len(stories); i++ {
		printOne(stories[i])
	}
}
