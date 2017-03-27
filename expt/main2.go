package main

import (
	"fmt"
	"gc8/sudoku"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

func main() {
	data, err := ioutil.ReadFile("/Users/kelvinyong/Downloads/sudoku17.txt")
	if err != nil {
		log.Fatal(err)
	}
	puzzles := strings.Split(string(data), "\n")

	workers := 4
	n := len(puzzles)
	batch := n / workers

	var wg sync.WaitGroup
	wg.Add(workers)

	start := time.Now()
	for i := 0; i < workers; i++ {
		var end int
		if i == workers-1 {
			end = n
		} else {
			end = (i + 1) * batch
		}

		go func(inputs []string) {
			for _, input := range inputs {
				s, _ := sudoku.NewSudoku(input)
				s.Solve(1)
			}
			wg.Done()
		}(puzzles[i*batch : end])
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(n, elapsed)
}
