package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	filePath     string
	timeout      int
	correct      int
	totQuestions int
)

func main() {
	initFlags()
	s, q := parseAndReadFile(filePath), 0
	inpChn := make(chan string)
	newTime, timeOutChan := make(chan struct{}), make(chan struct{})
	go startTimer(&newTime, timeOutChan, timeout)
	go getAnswers(&inpChn)
	fmt.Printf("question %v: %v = ? \n", q, s[q][0])
	func() {
		for {
			select {
			case <-timeOutChan:
				return
			case str := <-inpChn:
				newTime <- struct{}{}
				if str == s[q][1] {
					correct += 1
				}
				if q+1 >= totQuestions {
					return
				}
				q += 1
				fmt.Printf("question %v: %v = ? \n", q, s[q][0])
			}
		}
	}()
	fmt.Println("You got ", correct, "out of ", totQuestions, "questions")
}

func initFlags() {
	flag.IntVar(&timeout, "timeout", 15, "time to answer a question")
	flag.StringVar(&filePath, "filePath", "problems.csv", "path to CSV file with questions")
	flag.Parse()

}

func parseAndReadFile(fileStringName string) map[int][]string {
	file, fileErr := os.Open(fileStringName)
	if fileErr != nil {
		log.Fatal("Fatal error \n Cant find file \n Remember to add .csv file to folder")
		os.Exit(1)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, _ := csvReader.ReadAll()
	outMap := map[int][]string{}
	for q, ans := range data {
		totQuestions += 1
		outMap[q] = ans
	}
	fmt.Println("tot Questions:", totQuestions)
	return outMap
}

func getAnswers(chn *chan string) {
	var input string
	for {
		fmt.Scanln(&input)
		*chn <- input
	}

}

func startTimer(newTime *chan struct{}, timeOut chan struct{}, timeout int) {
	tim := time.NewTimer(time.Duration(timeout) * time.Second)
	for {
		select {
		case <-tim.C:
			timeOut <- struct{}{}
			fmt.Println("Time is up, too late")
		case <-*newTime:
			tim.Reset(time.Duration(timeout) * time.Second)
		}
	}
}
