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
	port := flag.Int("port", 3000, "specify the port the app should be served on")
	flag.Parse()
	story, err := cyoa.ParseStory(*filename)
	if err != nil {
		fmt.Println(err)
	}
	handle := cyoa.NewHandler(story)
	fmt.Printf("Starting new server on http://localhost:%d \n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handle))
}
