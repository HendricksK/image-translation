package imageimporter

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const filePath = "./config"

func Hello() string {
	return "Hello"
}

func GetAllUploadedImageTopics() []string {
	return retrieveRequestedDataFromFile("topics")
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

func GetImagesBasedOnTopic(topic string) []string {
	return retrieveRequestedDataFromFile(topic)
}

// https://gobyexample.com/reading-files
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
// Not going to pretend that I can write a file reder off by heart.
func retrieveRequestedDataFromFile(file_name string) []string {

	// we want to check if the file exists first
	topicFile, err := os.Open(filePath + "/" + file_name + ".txt")
	if err != nil {
		fmt.Println(err)
		// need to write to a log file at some point
		return []string{}
	}

	defer topicFile.Close()

	scanner := bufio.NewScanner(topicFile)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	var returnData []string

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		returnData = append(returnData, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return returnData
}
