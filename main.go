package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	
)

type Character []struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Image string `json:"image"`
}

type Results struct {
	Result Character `json:"results"`
}

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error Message : %s\n", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body)

}

func parseData(data string) Results {
	dataCharacter := Results{}

	err := json.Unmarshal([]byte(data), &dataCharacter)
	if err != nil {
		fmt.Printf("Error Message : %s\n", err)
	}
	return dataCharacter
}



func main() {

	

}
