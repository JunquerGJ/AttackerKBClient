package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JunquerGJ/AttackerKBClient/attackerkbclient"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := attackerkbclient.New(apiKey)

	topics, err := s.TopicSearch("Confluence")
	if err != nil {
		log.Fatalln(err)
	}
	for _, topic := range topics.Data {
		fmt.Printf("%18s\n", topic.Name)
	}
}
