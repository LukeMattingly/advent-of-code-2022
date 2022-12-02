package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	// open file
	f, err := os.Open("day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	highest := 0
	current := 0
	var top3 [3]int

	for scanner.Scan() {
		// do something with a line
		var cal int
		fmt.Sscanf(scanner.Text(), "%d", &cal)
		current += cal
		fmt.Println("current", current)
		if scanner.Text() == "" {
			fmt.Println("current, highest comparison", current, highest)
			if current > highest {
				highest = current
				fmt.Println("highest overtaken", highest)
			}
			sort.Ints(top3[:])
			if current > top3[0] {
				top3[0] = current
			}
			current = 0
		}
		fmt.Printf("line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	for _, c := range top3 {
		sum += c
	}
	fmt.Printf("highest:", highest)
	fmt.Printf("top3 sum:", sum)
}
