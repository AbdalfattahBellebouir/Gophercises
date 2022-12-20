package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Printf("The program running that has the name: %v\n", os.Args[0])

	var path string
	var timelimit int
	if len(os.Args[1:]) > 0 {
		path = os.Args[1]
		timelimit, _ = strconv.Atoi(os.Args[2])
	} else {
		path = "problems.csv"
		timelimit = 30
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var i int
	var sum int
	totalQuestions := len(data)
	done := make(chan int)
	t := time.NewTimer(time.Second * time.Duration(timelimit))

	go func(data [][]string, d chan int, t *time.Timer) {
		for _, v := range data {
			fmt.Printf("what is %v, sir?\n", v[0])
			fmt.Scan(&i)

			s, erro := strconv.Atoi(v[1])
			if erro != nil {
				log.Fatal(erro)
			}
			if s == i {
				sum++
			}
		}
		t.Stop()
		d <- 1
	}(data, done, t)

	select {
	case <-done:
		fmt.Printf("You finished the test.\n")
	case <-t.C:
		fmt.Printf("Time is up. %d seconds has passed\n", timelimit)
	}
	fmt.Printf("you have answered %v out of %v\n", sum, totalQuestions)

}
