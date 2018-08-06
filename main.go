package main

import (
	"flag"
	"fmt"
	"os"

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

func main() {
	var count = flag.Int("count", defaultCount, "How many stories that will be fetched")
	var feedType = flag.String("Feedtype", defaultFeed, "Feedtype: top|best|new")
	flag.Parse()

	var stories []Story
	stories = fetchStories(count, feedType)

	for i := 0; i < len(stories); i++ {
		printOne(stories[i])
	}
}
