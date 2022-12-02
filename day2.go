package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// open file
	f, err := os.Open("day2input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	m := make(map[string]int)

	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	m["X"] = 1
	m["Y"] = 2
	m["Z"] = 3

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	score := 0
	score2 := 0

	win := 6
	draw := 3
	loss := 0

	for scanner.Scan() {
		//var line string
		//fmt.Sscanf(scanner.Text(), "%d", &line)
		fmt.Printf("line: %s\n", scanner.Text())

		line := scanner.Text()
		opp := strings.Split(line, " ")

		if opp[0] == "A" { //They have Rock
			if opp[1] == "X" { //I have Rock
				score += m[opp[1]] + draw //draw =3 + my rock
				score2 += m["Z"] + loss   //need to lose: scizers + lose -
			}
			if opp[1] == "Y" { //I have Paper
				score += m[opp[1]] + win // Win + my paper
				score2 += m["X"] + draw  //need to draw: rock + draw -
			}
			if opp[1] == "Z" { //I have sciz
				score += m[opp[1]] + loss // lose + ssciz
				score2 += m["Y"] + win    //need to win: paper + win -
			}
		} else if opp[0] == "B" { //They have Paper
			if opp[1] == "X" { //I have rock
				score += m[opp[1]] + loss //lose + my rock
				score2 += m["X"] + loss   //need to lose: rock + lose -
			}
			if opp[1] == "Y" { //i have paper
				score += m[opp[1]] + draw //draw + my paper
				score2 += m["Y"] + draw   //need to draw: paper + draw -
			}
			if opp[1] == "Z" {
				score += m[opp[1]] + win //win + my sciz
				score2 += m["Z"] + win   //need to win: sciz + win -
			}
		} else if opp[0] == "C" { //they have sciz
			if opp[1] == "X" { //I have rock
				score += m[opp[1]] + win // win + rock
				score2 += m["Y"] + loss  //need to lose: paper + lose -
			}
			if opp[1] == "Y" { // paper
				score += m[opp[1]] + loss //lose + paper
				score2 += m["Z"] + draw   //need to draw: sciz + draw -
			}
			if opp[1] == "Z" {
				score += m[opp[1]] + draw //draw + sciz
				score2 += m["X"] + win    //need to win: rock + win -
			}
		}

		fmt.Println("score", score)
		fmt.Println("score2", score2)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("final", score)
	fmt.Println("score 2", score2)
}
