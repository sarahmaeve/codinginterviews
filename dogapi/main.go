// process breed information from https://dog.ceo/dog-api/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// fields must be capitalized!!
type ApiMessage struct {
	//	Status  string              `json:"status"`
	Breeds map[string][]string `json:"message"`
}

func getBreedList() *ApiMessage {
	breedUrl := "https://dog.ceo/api/breeds/list/all"
	resp, err := http.Get(breedUrl)
	if err != nil {
		fmt.Println("Unable to connect to page", breedUrl, " : ", err)
		os.Exit(1)
	}

	var m ApiMessage
	errJ := json.NewDecoder(resp.Body).Decode(&m)
	if errJ != nil {
		fmt.Println("GAAAH", err)
		os.Exit(1)
	}
	//	fmt.Println(m.Breeds["spaniel"])
	return &m
}

func checkBreedList(m *ApiMessage, breed string) bool {
	breed = strings.ToLower(breed)
	for k, v := range m.Breeds {
		if breed == k {
			return true
		} else {
			for _, sub := range v {
				if breed == sub+" "+k {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	m := getBreedList()
	myBreed := "herp derp"
	if checkBreedList(m, myBreed) == true {
		fmt.Println(myBreed, "is a dog breed.")
	} else {
		fmt.Println(myBreed, "is not a dog breed.")
	}

}
