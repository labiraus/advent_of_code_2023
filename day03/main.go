package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const (
	empty = iota
	gear
	other
)

type part struct {
	x, y          int
	partType      int
	value         int
	contributions int
}

func main() {
	f, err := os.Open("day03/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	parts := [][]part{}
	x := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			parts = append(parts, []part{})
			for y := 0; y < len(text); y++ {
				if unicode.IsNumber(rune(text[y])) || text[y] == '.' {
					parts[x] = append(parts[x], part{x: x, y: y, partType: empty})
				} else if text[y] == '*' {
					parts[x] = append(parts[x], part{x: x, y: y, partType: gear})
				} else {
					parts[x] = append(parts[x], part{x: x, y: y, partType: other})
				}
			}
			x++
		}
	}
	output := 0
	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	x = 0
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(fmt.Sprintf("%03d", x), text)
		if text == "" {
			continue
		} else {
			for y := 0; y < len(text); y++ {
				if !unicode.IsNumber(rune(text[y])) {
					continue
				} else {
					ystart := y
					number := ""
					for ; y < len(text) && unicode.IsNumber(rune(text[y])); y++ {
						number += string(text[y])
					}
					num, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}
					foundPart := checkParts(parts, x, ystart, len(number))
					if foundPart.partType == gear {
						if foundPart.contributions == 0 {
							foundPart.value = num
						} else {
							foundPart.value *= num
						}
						foundPart.contributions++
						parts[foundPart.x][foundPart.y] = foundPart
					}
				}
			}
			x++
		}
	}
	for _, row := range parts {
		for _, part := range row {
			if part.partType == gear && part.contributions == 2 {
				output += part.value
			}
		}
	}

	fmt.Println(output)
}

func checkParts(parts [][]part, i, j, count int) part {
	for k := -1; k <= 1; k++ {
		if i+k >= 140 || i+k < 0 {
			continue
		}
		for l := -1; l < count+1; l++ {
			if j+l >= 140 || j+l < 0 {
				continue
			}
			if parts[i+k][j+l].partType == gear {
				return parts[i+k][j+l]
			}
		}
	}
	return part{partType: empty}
}
