package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type csvData struct {
	Keyword string
	Value1  int
	Value2  float32
}

type fileData []csvData

func (d fileData) Len() int           { return len(d) }
func (d fileData) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d fileData) Less(i, j int) bool { return d[i].Value1 < d[j].Value1 }

// given a filename read the contents and return a data structure
func readCSV(filename string) (*[]csvData, error) {
	data := []csvData{}
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("could not open file %s %p", filename, err)
	}

	r := csv.NewReader(csvFile)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		keyword := line[0]
		val1, err := strconv.Atoi(line[1])
		if err != nil {
			log.Output(1, "atoi conversion error")
			break
		}
		parsed, _ := strconv.ParseFloat(line[2], 32)
		val2 := float32(parsed)
		data = append(data, csvData{Keyword: keyword, Value1: val1, Value2: val2})
	}

	return &data, nil
}

func main() {
	myData, _ := readCSV("data.txt")
	sort.Sort(fileData(*myData))

	for _, record := range *myData {
		fmt.Printf("%15s\t%4d\t%4.2f\n", record.Keyword, record.Value1, record.Value2)
	}
}
