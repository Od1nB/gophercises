package handle

import (
	"encoding/json"
	"os"
)

type Story map[string]Arc

type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func ParseStory() (Story, error) {
	var story Story
	f, err := os.Open("gopher.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	parser := json.NewDecoder(f)
	jsonerr := parser.Decode(&story)
	if jsonerr != nil {
		return nil, jsonerr
	}
	return story, nil

}
