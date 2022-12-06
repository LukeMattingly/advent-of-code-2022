package main

import (
	"fmt"
)

func main() {

	// Define the commands.
	commands := []struct {
		n        int
		from, to int
	}{
		{1, 7, 4},
		{3, 4, 7},
		{4, 3, 4},
		{5, 6, 9},
		{1, 8, 1},
		{2, 3, 2},
		{3, 4, 6},
		{1, 3, 6},
		{9, 7, 1},
	}

	// Initialize the queues.
	q1 := make([]string, 0)
	q2 := make([]string, 0)
	q3 := make([]string, 0)
	q4 := make([]string, 0)
	q5 := make([]string, 0)
	q6 := make([]string, 0)
	q7 := make([]string, 0)
	q8 := make([]string, 0)
	q9 := make([]string, 0)

	// Loop over the commands and execute them.
	for _, command := range commands {
		// Determine the source and destination queues.
		var src, dst *[]string
		switch command.from {
		case 1:
			src = &q1
		case 2:
			src = &q2
		case 3:
			src = &q3
		case 4:
			src = &q4
		case 5:
			src = &q5
		case 6:
			src = &q6
		case 7:
			src = &q7
		case 8:
			src = &q8
		case 9:
			src = &q9
		}
		switch command.to {
		case 1:
			dst = &q1
		case 2:
			dst = &q2
		case 3:
			dst = &q3
		case 4:
			dst = &q4
		case 5:
			dst = &q5
		case 6:
			dst = &q6
		case 7:
			dst = &q7
		case 8:
			dst = &q8
		case 9:
			dst = &q9
		}

		// Move the specified number of items from the source queue to the
		// destination queue.
		for i := 0; i < command.n; i++ {
			item := (*src)[0]
			*src = (*src)[1:]
			*dst = append(*dst, item)
		}
	}

	// Print the final state of the queues to verify the result.
	fmt.Println("q1 ", q1)
}
