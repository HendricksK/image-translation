package imageimporter

func Hello() string {
	return "Hello"
}

func GetImagesBasedOnTopic(topic string) []string {

	if topic == "" {
		return []string{}
	}

	return []string{
		"https://raw.githubusercontent.com/HendricksK/sacos_images/master/1/basil_williams_top_senior_1948.jpg",
		"https://raw.githubusercontent.com/HendricksK/sacos_images/master/1/basil_williams_top_senior_1948.jpg",
	}
}
