package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserId int    `json: "userId"`
	Id     int    `json: "id"`
	Title  string `json: "title"`
	Body   string `json: "body"`
}

func main() {
	var client = &http.Client{Timeout: 10 * time.Second}
	var url string = "https://jsonplaceholder.typicode.com/posts"

	res, err := client.Get(url)
	if err != nil {
		panic(err.Error())
	}

	var data []Post
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(data[0].Title)
}
