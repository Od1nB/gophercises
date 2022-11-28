package main

import (
	handler "cyoa"
	"fmt"
)

func main() {

	fmt.Println("new story")
	_, err := handler.ParseStory()
	if err != nil {
		fmt.Println(err)
	}

}
