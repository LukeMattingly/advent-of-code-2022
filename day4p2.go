//check top one is inside bottom on
//check if bottom is inside top on

//if left is greater than top left, check

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
	// open file
	f, err := os.Open("day4input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	insideTotal := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//var line string
		//fmt.Sscanf(scanner.Text(), "%d", &line)
		fmt.Printf("line: %s\n", scanner.Text())
		line := scanner.Text()
		split := strings.Split(line, ",")
		first := split[0]
		second := split[1]
		f := strings.Split(first, "-")
		s := strings.Split(second, "-")

		fmt.Println("first", first)
		fmt.Println("second", second)
		fmt.Println(f)
		fmt.Println(s)

		f0, _ := strconv.Atoi(f[0])
		f1, _ := strconv.Atoi(f[1])
		s0, _ := strconv.Atoi(s[0])
		s1, _ := strconv.Atoi(s[1])

		//if the right one on either top of bottom is smaller than left one cant be in there
		if f1 < s0 || s1 < f0 {
			fmt.Println("outside")
		} else {
			insideTotal++
			fmt.Println("overlap")
		}

	}
	fmt.Println(insideTotal)
}
