package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// StoryIDs are an array of integers
type StoryIDs []int

// Story structure from the API call
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// GetStory fetches a story based on its ID
func GetStory(id int) Story {
	baseURL := "https://hacker-news.firebaseio.com/v0/item/"
	url := baseURL + strconv.Itoa(id) + ".json"

	client := http.Client{Timeout: time.Second * 2}
	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Fatal(bodyErr)
	}

	var story Story
	jsonErr := json.Unmarshal(body, &story)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return story
}

// GetStoryIDs returns a list of the top story IDs
func GetStoryIDs() StoryIDs {
	url := "https://hacker-news.firebaseio.com/v0/topstories.json"

	client := http.Client{Timeout: time.Second * 2}
	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var storyIDs StoryIDs
	err = decoder.Decode(&storyIDs)

	return storyIDs
}

func main() {
	IDs := GetStoryIDs()
	fmt.Println(IDs)
	// story := GetStory(8863)
	// fmt.Println(stories)
}
