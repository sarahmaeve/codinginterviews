// process breed information from https://dog.ceo/dog-api/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
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

	return &m
}

func checkBreedList(m *ApiMessage, breed string) bool {
	breed = strings.ToLower(breed)
	for k, v := range m.Breeds {
		if breed == k || strings.Replace(breed, " ", "", -1) == k {
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

type BreedCount struct {
	Breed string
	Value int
}

func topSubbreeds(m *ApiMessage, count int) []BreedCount {
	// given the list of breeds, return a string with the breeds with
	// the most subbreeds, ordered by number
	// breed1 : 5, breed2: 3, breed3: 3 ....

	if count < 0 {
		count = 1
	}

	var ss []BreedCount
	// does this make sense to avoid a lot of copies during the append()?
	// (in general, as this is a small frequency set)
	ss = make([]BreedCount, 0, len(m.Breeds))
	for k, v := range m.Breeds {
		ss = append(ss, BreedCount{k, len(v)})
	}

	sort.SliceStable(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss[:count]
}

func main() {
	m := getBreedList()
	fmt.Println(topSubbreeds(m, 5))
}
