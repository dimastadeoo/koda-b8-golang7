package main

import (
	"encoding/json"
	"fetching-data/feature"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Character []struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
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

func getData(data *string) {
	*data = fetch("https://rickandmortyapi.com/api/character")

}

func main() {
	data := ""
	go getData(&data)
	spinner := []string{"|", "/", "-", "\\"}
	i := 0

	for data == "" {
		fmt.Printf("\rMengunduh data... %s", spinner[i])
		i = (i + 1) % len(spinner)
		time.Sleep(100 * time.Millisecond)
	}

	dataCharacter := parseData(data).Result

	// fmt.Println("--------------------------------------------------------")
	// fmt.Println("Semua Data: ")

	// for i, character := range dataCharacter{
	// 	fmt.Printf("%d. ", i+1)
	// 	fmt.Printf("%s\n",character.Name)
	// }
	// fmt.Println("--------------------------------------------------------")

	var choice string

	for {
		feature.CallClear()
		fmt.Println("Data Load")
		fmt.Print("Cari Nama Karakter Ketik 0 jika ingin Exit: ")
		fmt.Scan(&choice)
		choice = strings.ToLower(choice)

		if choice == "0" {
			break
		}

		dataSearchs := Character{}

		for _, character := range dataCharacter {
			if strings.Contains(strings.ToLower(character.Name), choice) {
				dataSearchs = append(dataSearchs, character)
			}
		}
		fmt.Println("--------------------------------------------------------")
		fmt.Println("Hasil Pencarian:")
		if dataSearchs != nil {
			fmt.Println("Not Found!")
		}

		for i, dataSearch := range dataSearchs {
			fmt.Printf("%d. ", i+1)
			fmt.Printf("%s\n", dataSearch.Name)
		}
		fmt.Println("--------------------------------------------------------")

		feature.WaitForKey("Tekan Enter untuk Ulangi lagi")

	}

}
