package main

import (
	"flag"
	"fmt"
	"regexp"
	"os"
)

func main() {
	providerUrl := os.Getenv("IPTV_URL")

	userHome, err := GetCurrentUserHomeDir()
	if err != nil {
		panic("failed to findout user home.")
	}
	filePath := fmt.Sprintf("%s/.local/iptv/channels.m3u", userHome)

	searchPattern := flag.String("s", "", "channel search pattern")
	updateChannels := flag.Bool("d", false, "update channel file")

	flag.Parse()

	if *updateChannels {
		fmt.Println("updating channels...")
		err := download(providerUrl, filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	channels, err := parseChannels(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	regexPattern := regexp.MustCompile(fmt.Sprintf("(?i)%s", *searchPattern))
	searchResults := make([]Channel, 0)

	for _, channel := range channels {
		if regexPattern.MatchString(channel.Name) {
			searchResults = append(searchResults, channel)
		}
	}

	for i, channel := range searchResults {
		fmt.Println(i, "->", channel.Name)
	}

	var channelNumber int
	fmt.Print("Channel Number: ")
	_, err = fmt.Scanf("%d", &channelNumber)

	fmt.Printf("Selected Channel: %s\n", searchResults[channelNumber].Name)
	copyToClipboard(searchResults[channelNumber].Url)
}
