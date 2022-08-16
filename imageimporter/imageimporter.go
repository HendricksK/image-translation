package imageimporter

import (
	"bufio"
	"log"
	"os"
)

const filePath = "./config"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Hello() string {
	return "Hello"
}

func GetImagesBasedOnTopicMock(topic string) []string {

	if topic == "" {
		return []string{}
	}

	return []string{
		"https://raw.githubusercontent.com/HendricksK/sacos_images/master/1/basil_williams_top_senior_1948.jpg",
		"https://raw.githubusercontent.com/HendricksK/sacos_images/master/1/basil_williams_top_senior_1948.jpg",
	}
}

// https://gobyexample.com/reading-files
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
// Not going to pretend that I can write a file reder off by heart.
func GetImagesBasedOnTopic(topic string) []string {

	topicFile, err := os.Open(filePath + "/test.txt")
	check(err)

	defer topicFile.Close()

	scanner := bufio.NewScanner(topicFile)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	var imageArr []string

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		imageArr = append(imageArr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return imageArr
}
