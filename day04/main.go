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
	f, err := os.Open("day04/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	cardNumber := 0
	cardCount := [213]int{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			cardCount[cardNumber]++
			points := 0
			sections := strings.Split(text, "|")

			winnerMap := map[int]bool{}
			for _, winnerString := range strings.Split(sections[0][10:], " ") {
				if winnerString == "" {
					continue
				}

				winner, err := strconv.Atoi(winnerString)
				if err != nil {
					panic(err)
				}
				winnerMap[winner] = false
			}

			for _, attemptString := range strings.Split(sections[1], " ") {
				if attemptString == "" {
					continue
				}
				attempt, err := strconv.Atoi(attemptString)
				if err != nil {
					panic(err)
				}
				if _, ok := winnerMap[attempt]; ok {
					winnerMap[attempt] = true
				}
			}

			for _, value := range winnerMap {
				if value {
					points++
				}
			}

			for j := 1; j <= points && j+cardNumber < 213; j++ {
				cardCount[cardNumber+j] += cardCount[cardNumber]
			}
		}
		cardNumber++
	}

	for _, count := range cardCount {
		output += count
	}

	fmt.Println(output)
}
