package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Create the queues.//try stacks
	queues := [][]string{
		[]string{"R", "P", "C", "D", "B", "G"},
		[]string{"H", "V", "G"},
		[]string{"N", "S", "Q", "D", "J", "P", "M"},
		[]string{"P", "S", "L", "G", "D", "C", "N", "M"},
		[]string{"J", "B", "N", "C", "P", "F", "L", "S"},
		[]string{"Q", "B", "D", "Z", "V", "G", "T", "S"},
		[]string{"B", "Z", "M", "H", "F", "T", "Q"},
		[]string{"C", "M", "D", "B", "F"},
		[]string{"F", "C", "Q", "G"},
	}

	input, err := readLines("day5input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//defer input.Close()

	// Loop over the strings and parse them.
	for _, s := range input {
		// Split the string into substrings.
		parts := strings.Split(s, " ")

		// Convert the first substring to an integer.
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		fmt.Println("number of items to move", n)

		// Convert the second and third substrings to integers.
		from, to := 0, 0
		if from, err = strconv.Atoi(parts[3]); err != nil {
			// Handle the error.
		}
		if to, err = strconv.Atoi(parts[5]); err != nil {
			// Handle the error.
		}
		if len(queues[from-1]) < n {
			fmt.Println("Error: Cannot move more items than there are in the queue.")
			continue
		}
		fmt.Println("from ", from)
		fmt.Println("to ", to)

		// Move the specified number of items from one queue to another.
		for i := 0; i < n; i++ {
			//STacks the
			if len(queues[from-1]) == 0 {
				break
			}
			item := queues[from-1][len(queues[from-1])-1]
			queues[from-1] = queues[from-1][:len(queues[from-1])-1]
			queues[to-1] = append(queues[to-1], item)

		}

		// Print the updated queues.
		fmt.Println(queues)

		// Print the values of each stack.
		for i, s := range queues {
			fmt.Printf("Stack %d: ", i+1)
			for _, e := range s {
				fmt.Printf("%s ", e)
			}
			fmt.Println()
		}

	}

	for i, s := range queues {
		fmt.Printf("Top of stack %d: %s\n", i+1, s[len(s)-1])
	}

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
