package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var intConvert1 = map[string]rune{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var intConvert2 = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func main() {
	f, err := os.Open("day01/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			intArray := []int{}
			pos := 0
			for pos <= len(text)-1 {
				found := false
				for key, value := range intConvert2 {
					if strings.HasPrefix(text[pos:], key) {
						found = true
						intArray = append(intArray, value)
						pos++
						break
					}
				}
				if !found {
					pos++
				}
			}
			val := intArray[0]*10 + intArray[len(intArray)-1]
			output += val
			fmt.Println(text, val, intArray, output)
		}
	}

	fmt.Println(output)
}
