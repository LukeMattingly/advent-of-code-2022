package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	// open file
	f, err := os.Open("day3input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()


	scanner := bufio.NewScanner(f)
	sum := 0

	var firstString string =""
	var secondString string =""
	var thirdString string =""
	line:=0
	for line < 3{
		for scanner.Scan(){
			fmt.Printf("line: %s\n", scanner.Text())
			firstSet := mapset.NewSet[string]()
			secondSet := mapset.NewSet[string]()
			thirdSet := mapset.NewSet[string]()
			if line ==0{
				firstString = scanner.Text()
				fmt.Println("First",firstString)
				for _, a := range strings.Split(firstString, "") {
					firstSet.Add(a)
					fmt.Println(firstSet)
				}
			}
			if line ==1{
				secondString = scanner.Text()
				fmt.Println("SEcond",secondString)
				for _, b := range strings.Split(secondString, "") {
					secondSet.Add(b)
					fmt.Println(secondSet)
				}
			}
			if line ==2{
				thirdString = scanner.Text()
				fmt.Println("thirdString",thirdString)
				for _, c := range strings.Split(thirdString, "") {
					thirdSet.Add(c)
					fmt.Println(thirdSet)
				}
			}
			line++
			if line==3{
				line=0
				fmt.Println("space")
				overlap := firstSet.Intersect(secondSet).Intersect(thirdSet)
				fmt.Println("overlap ", overlap)
		
				overlapValue := ""
				for i := range overlap.Iter() {
					overlapValue = i
				}
				var itemPriority = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}
				sum += itemPriority[overlapValue]
				fmt.Println(sum)
			}
		}
	}
}
