package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type Channel struct {
	Name string
	Url  string
	Id   string
}

func NewChannel(name string, url string) *Channel {
	return &Channel{
		Name: name,
		Url:  url,
		Id:   uuid.New().String(),
	}
}

func clearShizzelFromName(name string) string {
	name = strings.ReplaceAll(name, "'", "")
	return name
}

func parseChannels(filepath string) ([]Channel, error) {
	pattern := regexp.MustCompile(`^([^\s]+)\s#Name:\s(.*)$`)
	var channels []Channel

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := pattern.FindStringSubmatch(line)
		if matches != nil {
			clearedUrl := matches[1]
			clearedName := clearShizzelFromName(matches[2])

			channels = append(channels, *NewChannel(clearedName, clearedUrl))
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}
