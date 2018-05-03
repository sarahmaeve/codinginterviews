// https://gophercises.com/exercises/quiz
// Gophercises #1: create a quiz that reads a csv file and asks the user questions
// return a score
// version 2: add a time limit; when at limit return score and exit

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var score, seconds, questions int

func init() {
	flag.IntVar(&seconds, "timer", 30, "time for quiz in seconds")
	flag.IntVar(&questions, "questions", -1, "number of questions")
	flag.Parse()
}

func main() {
	quizfile := "quiz1.csv"
	f, err := os.OpenFile(quizfile, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// do ReadAll() instead of step by step read so we can report # of questions
	qlist, _ := csv.NewReader(f).ReadAll()
	var target int
	target = len(qlist)
	if questions > 0 && questions < target {
		target = questions
	}
	defer f.Close()
	c := make(chan int)
	go quiz(qlist, questions, c)
	go timer(seconds, c)
	// add when receiving
	for i := range c {
		if i > 0 {
			score += i
		}
	}
	fmt.Printf("\nYou answered %v questions correctly out of %v.\n", score, target)
}

func timer(s int, c chan int) {
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Println("\nTime's up!")
	close(c)
	return
}

func shuffle(qlist [][]string) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range qlist {
		swap := rand.Intn(len(qlist) - 1)
		qlist[i], qlist[swap] = qlist[swap], qlist[i]
	}
}

func quiz(qlist [][]string, questions int, c chan int) {
	var counter int
	shuffle(qlist)
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
	close(c)
	return
}
