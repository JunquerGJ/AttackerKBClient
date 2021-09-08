package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/JunquerGJ/AttackerKBClient/attackerkbclient"
	"github.com/fatih/color"
)

var vulnTitle = color.New(color.FgWhite).Add(color.Bold).Add(color.Underline)
var vulnDescription = color.New(color.FgWhite)
var vulnScore = color.New(color.FgRed).Add(color.Bold)
var header = color.New(color.FgGreen).Add(color.Bold)
var detail = color.New(color.FgYellow)

func word_wrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped

}

func printTopic(i int, topic attackerkbclient.Topic) {
	vulnTitle.Print(fmt.Sprintf("[%d] %s", i, topic.Name))
	header.Print(" Attacker Value ")
	vulnScore.Print(topic.Score.AttackerValue)
	header.Print(" Exploitability ")
	vulnScore.Println(topic.Score.Exploitability)
	fmt.Printf("%s\n", topic.Metadata.BaseMetricV3.CVSSV3.VectorString)
	header.Println("Description")
	fmt.Println(word_wrap(topic.Document, 100))

	//	vulnDescription.Printf("    %20s\n", topic.Document)
	fmt.Println()

}

func printTopicWithDetails(i int, topic attackerkbclient.Topic) {
	printTopic(i, topic)
	header.Println("Affected versions")
	detail.Print("- ")
	detail.Println(strings.Join((topic.Metadata.VulnerableVersions), "\n- "))
	fmt.Println()
}

func main() {
	var state = "list"
	var currentTopic = 0
	ch := make(chan string)
	go func(ch chan string) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		var b []byte = make([]byte, 1)
		for {
			os.Stdin.Read(b)
			var a = string(b)
			if a != "\n" {
				ch <- string(a)
			}
		}
	}(ch)

	//	colorRed := color.New(color.FgRed)
	//	colorGreen := color.New(color.FgGreen)

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:8080")
	apiKey := os.Getenv("API_KEY")
	s := attackerkbclient.New(apiKey)

	topics, err := s.TopicSearch(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	for i, topic := range topics.Data {
		printTopic(i, topic)
		/*		vulnTitle.Print(fmt.Sprintf("[%d] %s", i, topic.Name))
				//		vulnTitle.Println(topic.Name)
				header.Print("Score ")
				vulnScore.Println(topic.Score.AttackerValue)
				header.Println("Description")
				vulnDescription.Println(topic.Document)
				fmt.Println()*/
	}

	for {
		select {
		case stdin, _ := <-ch:
			i, err := strconv.Atoi(stdin)
			if err == nil {
				fmt.Print("\033[H\033[2J")
				state = "detail"
				currentTopic = i
				printTopic(i, topics.Data[i])
			} else {
				fmt.Println(stdin)
				switch stdin {
				case "q":
					exec.Command("reset").Run()
					os.Exit(0)
				case "l":
					state = "list"
					for i, topic := range topics.Data {
						printTopic(i, topic)
					}
				case "a":
					if state == "detail" {
						fmt.Print("\033[H\033[2J")
						currentTopic = currentTopic - 1
						if currentTopic < 0 {
							currentTopic = 0
						}
						printTopic(currentTopic, topics.Data[currentTopic])
					}
				case "d":
					if state == "detail" {
						fmt.Print("\033[H\033[2J")
						currentTopic = currentTopic + 1
						if currentTopic > 9 {
							currentTopic = 9
						}
						printTopic(currentTopic, topics.Data[currentTopic])
					}
				case "s":
					if state == "detail" {
						fmt.Print("\033[H\033[2J")
						printTopicWithDetails(currentTopic, topics.Data[currentTopic])
					}
				}
			}

		}
		time.Sleep(time.Millisecond * 100)
	}
}
