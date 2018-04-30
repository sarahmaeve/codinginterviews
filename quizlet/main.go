// https://gophercises.com/exercises/quiz
// Gophercises #1: create a quiz that reads a csv file and asks the user questions
// return a score
// version 2: add a time limit; when at limit return score and exit
// TODO: randomize question order

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var score int

func main() {
	quizfile := "quiz1.csv"
	f, err := os.OpenFile(quizfile, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	pSeconds := flag.Int("timer", 30, "time for quiz in seconds")
	pQuestions := flag.Int("questions", -1, "number of questions")
	flag.Parse()

	// do ReadAll() instead of step by step read so we can report # of questions
	qlist, _ := csv.NewReader(f).ReadAll()
	var target int
	target = len(qlist)
	if *pQuestions > 0 && *pQuestions < target {
		target = *pQuestions
	}
	defer f.Close()
	c := make(chan int)
	go quiz(qlist, *pQuestions, c)
	go timer(*pSeconds, c)
	// add when receiving
	for i := range c {
		if i > 0 {
			score += i
		} else {
			fmt.Printf("\nYou answered %v questions correctly out of %v.\n", score, target)
			break
		}
	}
}

func timer(s int, c chan int) {
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Println("\nTime's up!")
	c <- 0
	return
}

func quiz(qlist [][]string, questions int, c chan int) {
	var counter int
	for _, line := range qlist {
		if questions > 0 && counter >= questions {
			break
		} else {

			fmt.Print(line[0], "? ")
			answer := line[1]
			var response string
			fmt.Scanln(&response)
			counter++
			if strings.ToLower(answer) == strings.ToLower(response) {
				c <- 1
			}
		}
	}
	c <- 0
	return
}
