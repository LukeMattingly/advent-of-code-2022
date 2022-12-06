package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Create the queues.
	queues := []*ring.Ring{
		ring.New(6),
		ring.New(3),
		ring.New(8),
		ring.New(8),
		ring.New(8),
		ring.New(7),
		ring.New(7),
		ring.New(5),
		ring.New(4),
	}
	queues[0].Value = "G"
	queues[0].Next().Value = "B"
	queues[0].Next().Next().Value = "D"
	queues[0].Next().Next().Next().Value = "C"
	queues[0].Next().Next().Next().Next().Value = "P"
	queues[0].Next().Next().Next().Next().Next().Value = "R"
	queues[1].Value = "G"
	queues[1].Next().Value = "V"
	queues[1].Next().Next().Value = "H"
	queues[2].Value = "M"
	queues[2].Next().Value = "M"
	queues[2].Next().Next().Value = "P"
	queues[2].Next().Next().Next().Value = "J"
	queues[2].Next().Next().Next().Next().Value = "D"
	queues[2].Next().Next().Next().Next().Next().Value = "Q"
	queues[2].Next().Next().Next().Next().Next().Next().Value = "S"
	queues[2].Next().Next().Next().Next().Next().Next().Next().Value = "N"
	queues[3].Value = "S"
	queues[3].Next().Value = "N"
	queues[3].Next().Next().Value = "C"
	queues[3].Next().Next().Next().Value = "D"
	queues[3].Next().Next().Next().Next().Value = "G"
	queues[3].Next().Next().Next().Next().Next().Value = "L"
	queues[3].Next().Next().Next().Next().Next().Next().Value = "S"
	queues[3].Next().Next().Next().Next().Next().Next().Next().Value = "P"
	queues[4].Value = "S"
	queues[4].Next().Value = "L"
	queues[4].Next().Next().Value = "F"
	queues[4].Next().Next().Next().Value = "P"
	queues[4].Next().Next().Next().Next().Value = "C"
	queues[4].Next().Next().Next().Next().Next().Value = "N"
	queues[4].Next().Next().Next().Next().Next().Next().Value = "B"
	queues[4].Next().Next().Next().Next().Next().Next().Next().Value = "J"
	queues[5].Value = "T"
	queues[5].Next().Value = "G"
	queues[5].Next().Next().Value = "V"
	queues[5].Next().Next().Next().Value = "Z"
	queues[5].Next().Next().Next().Next().Value = "D"
	queues[5].Next().Next().Next().Next().Next().Value = "B"
	queues[5].Next().Next().Next().Next().Next().Next().Value = "Q"
	queues[6].Value = "Q"
	queues[6].Next().Value = "T"
	queues[6].Next().Next().Value = "F"
	queues[6].Next().Next().Next().Value = "H"
	queues[6].Next().Next().Next().Next().Value = "M"
	queues[6].Next().Next().Next().Next().Next().Value = "Z"
	queues[6].Next().Next().Next().Next().Next().Next().Value = "B"
	queues[7].Value = "F"
	queues[7].Next().Value = "B"
	queues[7].Next().Next().Value = "D"
	queues[7].Next().Next().Next().Value = "M"
	queues[7].Next().Next().Next().Next().Value = "C"
	queues[7].Value = "G"
	queues[8].Next().Value = "Q"
	queues[8].Next().Next().Value = "C"
	queues[8].Next().Next().Next().Value = "F"

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
		/*
			if len(queues[from-1]) < n {
				fmt.Println("Error: Cannot move more items than there are in the queue.")
				continue
			}
		*/
		fmt.Println("from ", from)
		fmt.Println("to ", to)

		// Move the specified number of items from one queue to another.
		//moved := queues[from-1][:n]
		//queues[from-1] = queues[from-1][n:]
		//queues[to-1] = append(queues[to-1], moved...)

		// Move the specified number of items from one queue to another.
		for i := 0; i < n; i++ {
			// Pop the first element from the source queue and push it onto the destination queue.
			item := queues[from-1].Next()
			queues[from-1].Link(item)
			queues[to-1].Link(item)
		}

	}
	// Print the values of each queue.
	for i, q := range queues {
		fmt.Printf("Queue %d: ", i+1)
		for e := q.Next(); e != nil; e = e.Next() {
			fmt.Printf("%s ", e.Value.(string))
		}
		fmt.Println()
	}

	// Print the top of each queue.
	for i, q := range queues {
		fmt.Printf("Top of queue %d: %s\n", i+1, q.Next().Value)
	}
	/*
		for _, q := range queues {
			fmt.Println("top of queue", q[len(q)-1])
		}*/
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
