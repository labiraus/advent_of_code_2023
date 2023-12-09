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
	f, err := os.Open("day09/input.txt")

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
			row := []int{}
			for _, val := range strings.Split(text, " ") {
				num, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				row = append(row, num)
			}
			output += determineNext(row)
		}
	}

	fmt.Println(output)
}

func determineNext(row []int) int {
	newRow := []int{}
	isZero := true

	for i := 1; i < len(row); i++ {
		diff := row[i] - row[i-1]
		newRow = append(newRow, diff)
		isZero = isZero && diff == 0
	}
	if isZero {
		return row[0]
	} else {
		return row[0] - determineNext(newRow)
	}
}
