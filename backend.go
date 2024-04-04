package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main2() {
	response, error := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if error != nil {
		log.Fatalln(error)
	}
	defer response.Body.Close()

	var post Post
	err := json.NewDecoder(response.Body).Decode(&post)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(post.Body)

}
