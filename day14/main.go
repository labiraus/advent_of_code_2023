package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var gridSize int

func main() {
	f, err := os.Open("day14/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	grid := [][]rune{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			gridSize++
			grid = append(grid, []rune(text))
		}
	}
	max := 1000000000
	breakMap := map[string]int{}
	for x := 0; x <= max; x++ {
		for dir := 0; dir < 4; dir++ {
			grid = roll(grid)
			grid = rotate(grid)
		}
		// fmt.Println(x, calculate(grid))

		if line, ok := breakMap[toString(grid)]; ok {
			// fmt.Println(x, calculate(grid))
			if (max-x-1)%(x-line) == 0 {
				mult := (max - x) / (x - line)
				fmt.Println(x-line, max-x-1, mult, mult*(x-line)+x)
				break
			}
		} else {
			breakMap[toString(grid)] = x
		}
	}

	fmt.Println(calculate(grid))
}

func calculate(grid [][]rune) int {
	output := 0
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if grid[row][col] == 'O' {
				output += gridSize - row
			}
		}
	}
	return output
}

func printGrid(grid [][]rune) {
	for i := range grid {
		fmt.Println(string(grid[i]))
	}
}

func toString(grid [][]rune) string {
	output := ""
	for i := range grid {
		output += string(grid[i])
	}
	return output
}

func rotate(grid [][]rune) [][]rune {
	size := len(grid)
	rotated := make([][]rune, size)
	for i := range rotated {
		rotated[i] = make([]rune, size)
	}

	for i := range grid {
		for j := range grid[i] {
			rotated[j][size-i-1] = grid[i][j]
		}
	}

	return rotated
}

func roll(grid [][]rune) [][]rune {
	outputs := make([][]rune, gridSize)
	for i := 0; i < gridSize; i++ {
		outputs[i] = []rune(strings.Repeat(".", gridSize))
	}

	for col := 0; col < gridSize; col++ {
		end := 0
		for row := 0; row < len(grid); row++ {
			switch grid[row][col] {
			case 'O':
				outputs[end][col] = 'O'
				end++
			case '#':
				outputs[row][col] = '#'
				end = row + 1
			}
		}
	}
	return outputs
}
