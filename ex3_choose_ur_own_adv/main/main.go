package main

import (
	cyoa "cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	filename := flag.String("filename", "gopher.json", "a json file for your cyoa")
	flag.Parse()

	fmt.Println("new story")
	story, err := cyoa.ParseStory(*filename)
	if err != nil {
		fmt.Println(err)
	}
	handle := cyoa.NewHandler(story)

	fmt.Println("Starting new server")
	log.Fatal(http.ListenAndServe(":3000", handle))
}
