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
	f, err := os.Open("day02/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output1 := 0
	output2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			game := strings.Split(text, ":")
			gameID, err := strconv.Atoi(game[0][5:])
			if err != nil {
				panic(err)
			}
			sets := strings.Split(game[1], ";")
			colourMax := make(map[string]int)
			for _, set := range sets {
				colours := strings.Split(set, ",")
				for _, colour := range colours {
					tuple := strings.Split(strings.TrimSpace(colour), " ")
					count, err := strconv.Atoi(tuple[0])
					if err != nil {
						panic(err)
					}
					if count > colourMax[tuple[1]] {
						colourMax[tuple[1]] = count
					}
				}
			}
			if colourMax["red"] == 0 {
				colourMax["red"] = 1
			}
			if colourMax["blue"] == 0 {
				colourMax["blue"] = 1
			}
			if colourMax["green"] == 0 {
				colourMax["green"] = 1
			}

			if colourMax["red"] <= 12 && colourMax["green"] <= 13 && colourMax["blue"] <= 14 {
				output1 += gameID
			}
			power := colourMax["red"] * colourMax["green"] * colourMax["blue"]
			output2 += power
		}
	}

	fmt.Println(output2)
	// 1805 too low
}
