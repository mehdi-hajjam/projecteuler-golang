//euler22 deals with the 22nd euler project problem. It gives insight into file reading, interface io.reader, sorting and concurrency.
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

//alphabet returns the alphabetical value of a string ie the sum of the rank of the letters in a string
func alphabet(a string) int {
	var sum int
	sum = 0
	for _, runeValue := range a {
		switch string(runeValue) {
		case "A":
			sum = sum + 1
		case "B":
			sum = sum + 2
		case "C":
			sum = sum + 3
		case "D":
			sum = sum + 4
		case "E":
			sum = sum + 5
		case "F":
			sum = sum + 6
		case "G":
			sum = sum + 7
		case "H":
			sum = sum + 8
		case "I":
			sum = sum + 9
		case "J":
			sum = sum + 10
		case "K":
			sum = sum + 11
		case "L":
			sum = sum + 12
		case "M":
			sum = sum + 13
		case "N":
			sum = sum + 14
		case "O":
			sum = sum + 15
		case "P":
			sum = sum + 16
		case "Q":
			sum = sum + 17
		case "R":
			sum = sum + 18
		case "S":
			sum = sum + 19
		case "T":
			sum = sum + 20
		case "U":
			sum = sum + 21
		case "V":
			sum = sum + 22
		case "W":
			sum = sum + 23
		case "X":
			sum = sum + 24
		case "Y":
			sum = sum + 25
		case "Z":
			sum = sum + 26
		}
	}
	return sum
}

func main() {

	//Opens the file
	file, err := os.Open("src/github.com/user/Euler22/data_euler22.txt")
	if err != nil {
		fmt.Println("There is a problem with reading the file!")
		fmt.Print(err)
	}
	//Creates a csv reader out of it
	list := csv.NewReader(bufio.NewReader(file))
	//Reads from the reader
	record, err := list.Read()
	//Sorts in increasing order
	sort.Strings(record)
	//Prints it on screen
	fmt.Print(record)

	var result []int

	//For each record, compute and store the alphabetical value of the name
	for j, elem := range record {

		result = append(result, alphabet(elem)*(j+1))
	}

	//then sum all the values in the result struct
	s := 0
	for _, v := range result {
		s = s + v
	}
	//print the result
	fmt.Print(s)

}
