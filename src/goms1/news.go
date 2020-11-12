package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Data Struct
type Data struct {
	Source struct {
		ID   int    `json: "id"`
		Name string `json: "name"`
	}
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	UrltoImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

// Result struct
type Result struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults`
	Articles     []Data
}

func news() {
	url := "https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=68abed74b6aa468aa1c2bcdd04ccaca9"

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close() // after everything is done close the connection. if sucess then close.
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			// fmt.Println(string(body))
			var result Result
			json.Unmarshal(body, &result)
			for i := 0; i < len(result.Articles); i++ {
				fmt.Println("=============================")
				fmt.Println(result.Articles[i].Source.Name)
				fmt.Println("=============================")
				fmt.Println(result.Articles[i].Author)
				fmt.Println(result.Articles[i].Title)
				fmt.Println(result.Articles[i].Description)
				fmt.Println("=============================")
			}

		}
	}
}
