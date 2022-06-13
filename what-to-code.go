package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davipatricio/colors/colors"
	"github.com/davipatricio/colors/styles"
)

const API_URL string = "https://what-to-code.com/api/ideas/random"

type Tag struct {
	Value string `json:"value"`
}

type Idea struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        []Tag  `json:"tags"`
	Likes       uint64 `json:"likes"`
}

func main() {
	resp, err := http.Get(API_URL)
	if err != nil {
		panic(err)
	}

	var idea Idea
	err = json.NewDecoder(resp.Body).Decode(&idea)
	if err != nil {
		panic(err)
	}

	fmt.Printf("    %v\n", styles.Bold(idea.Title))
	if len(idea.Description) > 0 {
		fmt.Printf("%v\n", idea.Description)
	}
	fmt.Printf("%v %v\n", colors.Red("Likes:"), idea.Likes)

	if len(idea.Tags) > 0 {
		finalTags := ""
		for i, tag := range idea.Tags {
			if i > 0 {
				finalTags += ", "
			}
			finalTags += tag.Value
		}
		fmt.Printf("%v", colors.Green("Tags: ")+finalTags)
	}
}
