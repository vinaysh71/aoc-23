package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func total(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func getNumbersFromString(input string) (int, int, error) {
	var first, last int
	for _, t := range input {
		if unicode.IsNumber(t) {
			if first == 0 {
				first = int(t - '0')
			}
			last = int(t - '0')
		}
	}
	return first, last, nil
}

func readFromFile(filename string) ([]int, error) {
	var result []int
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		value := scanner.Text()
		first, last, err := getNumbersFromString(value)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, first*10+last)
	}
	return result, nil
}

func main() {
	slice := []int{}
	slice, err := readFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total is =", total(slice...))
}
