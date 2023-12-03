package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func total(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func getNumbersFromString(input string) (int, int, int, int, error) {
	var first, last, first_2, last_2 int
	for pos, t := range input {
		if unicode.IsNumber(t) {
			if first == 0 {
				first = int(t - '0')
				if first_2 == 0 {
					first_2 = first
				}
			}
			last = int(t - '0')
			last_2 = last
		}
		p2 := convertToNumber(input[pos:])
		if p2 != 0 {
			if first_2 == 0 {
				first_2 = p2
			}
			last_2 = p2
		}
	}
	return first, last, first_2, last_2, nil
}

func convertToNumber(input string) int {
	num_map := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	var res int
	for key, val := range num_map {
		if strings.HasPrefix(input, key) {
			res = val
			break
		}
	}
	return res
}

func readFromFile(filename string) ([]int, []int, error) {
	var result_p1 []int
	var result_p2 []int
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		value := scanner.Text()
		first, last, first_2, last_2, err := getNumbersFromString(value)
		if err != nil {
			return nil, nil, err
		}
		result_p1 = append(result_p1, first*10+last)
		result_p2 = append(result_p2, first_2*10+last_2)
	}
	return result_p1, result_p2, nil
}

func main() {
	slice_p1 := []int{}
	slice_p2 := []int{}
	slice_p1, slice_p2, err := readFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 total =", total(slice_p1...))
	fmt.Println("Part 2 total =", total(slice_p2...))
}
