package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type galaxy struct {
	x      int
	xShift int
	y      int
	yShfit int
}

const shift = 999999

func main() {
	f, err := os.Open("day11/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	galaxies := []*galaxy{}
	blankColumns := []bool{}
	i := 0
	yShift := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			rowBlank := true
			for j, val := range text {
				if i == 0 {
					blankColumns = append(blankColumns, val == '.')
				} else {
					blankColumns[j] = blankColumns[j] && val == '.'
				}

				if val == '#' {
					rowBlank = false
					galaxies = append(galaxies, &galaxy{x: j, y: i, yShfit: yShift})
				}
			}
			i++
			if rowBlank {
				yShift += shift
			}
		}
	}
	for j, val := range blankColumns {
		if val {
			for _, galaxy := range galaxies {
				if galaxy.x > j {
					galaxy.xShift += shift
				}
			}
		}
	}

	for i, galaxy1 := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			output += galaxyDistance(*galaxy1, *galaxies[j])
		}
	}

	fmt.Println(output)
}

func galaxyDistance(galaxy1 galaxy, galaxy2 galaxy) int {
	x := galaxy1.x - galaxy2.x + galaxy1.xShift - galaxy2.xShift
	y := galaxy1.y - galaxy2.y + galaxy1.yShfit - galaxy2.yShfit
	dist := abs(x) + abs(y)
	return dist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//678626878094 too high
