package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("day10/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	i := 0
	vertical := 0
	horizontal := 0
	graph := [][]rune{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			graph = append(graph, []rune{})
			for j, c := range text {
				if c == 'S' {
					vertical = i
					horizontal = j
				}
				graph[i] = append(graph[i], c)
			}
			i++
		}
	}
	output = traversGraph(vertical+1, horizontal, "S", graph)

	fmt.Println(output)
}

// 13379 too high

func traversGraph(vertical, horizontal int, direction string, graph [][]rune) int {
	//fmt.Println("------------------")
	output := 0
	maxHorizontal := len(graph[0])
	maxVertical := len(graph)
	fmt.Println(maxHorizontal, maxVertical)
	followedPath := [][]rune{}
	for i := 0; i < len(graph); i++ {
		followedPath = append(followedPath, []rune{})
		for j := 0; j < len(graph[i]); j++ {
			followedPath[i] = append(followedPath[i], '.')
		}
	}

	for {
		exit := true
		followedPath[vertical][horizontal] = graph[vertical][horizontal]
		// fmt.Println(string(graph[vertical][horizontal]), " ", string(direction), " ")
		switch graph[vertical][horizontal] {
		case '|':
			if direction == "N" {
				if horizontal+1 < maxHorizontal && followedPath[vertical][horizontal+1] == '.' {
					followedPath[vertical][horizontal+1] = 'O'
				}
				if horizontal-1 >= 0 && followedPath[vertical][horizontal-1] == '.' {
					followedPath[vertical][horizontal-1] = 'I'
				}
				vertical--
				if vertical >= 0 {
					exit = false
				}
			} else if direction == "S" {
				if horizontal+1 < maxHorizontal && followedPath[vertical][horizontal+1] == '.' {
					followedPath[vertical][horizontal+1] = 'I'
				}
				if horizontal-1 >= 0 && followedPath[vertical][horizontal-1] == '.' {
					followedPath[vertical][horizontal-1] = 'O'
				}
				vertical++
				if vertical < maxVertical {
					exit = false
				}
			}
		case '-':
			if direction == "W" {
				if vertical+1 < maxVertical && followedPath[vertical+1][horizontal] == '.' {
					followedPath[vertical+1][horizontal] = 'I'
				}
				if vertical-1 >= 0 && followedPath[vertical-1][horizontal] == '.' {
					followedPath[vertical-1][horizontal] = 'O'
				}
				horizontal--
				if horizontal >= 0 {
					exit = false
				}
			} else if direction == "E" {
				if vertical+1 < maxVertical && followedPath[vertical+1][horizontal] == '.' {
					followedPath[vertical+1][horizontal] = 'O'
				}
				if vertical-1 >= 0 && followedPath[vertical-1][horizontal] == '.' {
					followedPath[vertical-1][horizontal] = 'I'
				}
				horizontal++
				if horizontal < maxHorizontal {
					exit = false
				}
			}
		case 'L':
			if direction == "W" {
				if vertical+1 < maxVertical && followedPath[vertical+1][horizontal] == '.' {
					followedPath[vertical+1][horizontal] = 'I'
				}
				if horizontal-1 >= 0 && followedPath[vertical][horizontal-1] == '.' {
					followedPath[vertical][horizontal-1] = 'I'
				}
				vertical--
				direction = "N"
				if vertical >= 0 {
					exit = false
				}
			} else if direction == "S" {
				if vertical+1 < maxVertical && followedPath[vertical+1][horizontal] == '.' {
					followedPath[vertical+1][horizontal] = 'O'
				}
				if horizontal-1 >= 0 && followedPath[vertical][horizontal-1] == '.' {
					followedPath[vertical][horizontal-1] = 'O'
				}
				horizontal++
				direction = "E"
				if horizontal < maxHorizontal {
					exit = false
				}
			}
		case 'J':
			if direction == "S" {
				if vertical+1 < maxVertical && followedPath[vertical+1][horizontal] == '.' {
					followedPath[vertical+1][horizontal] = 'I'
				}
				if horizontal+1 < maxHorizontal && followedPath[vertical][horizontal+1] == '.' {
					followedPath[vertical][horizontal+1] = 'I'
				}
				horizontal--
				direction = "W"
				if horizontal >= 0 {
					exit = false
				}
			} else if direction == "E" {
				if vertical+1 < maxVertical && followedPath[vertical+1][horizontal] == '.' {
					followedPath[vertical+1][horizontal] = 'O'
				}
				if horizontal+1 < maxHorizontal && followedPath[vertical][horizontal+1] == '.' {
					followedPath[vertical][horizontal+1] = 'O'
				}
				vertical--
				direction = "N"
				if vertical < maxVertical {
					exit = false
				}
			}
		case '7':
			if direction == "N" {
				if vertical-1 >= 0 && followedPath[vertical-1][horizontal] == '.' {
					followedPath[vertical-1][horizontal] = 'O'
				}
				if horizontal+1 < maxHorizontal && followedPath[vertical][horizontal+1] == '.' {
					followedPath[vertical][horizontal+1] = 'O'
				}
				horizontal--
				direction = "W"
				if horizontal >= 0 {
					exit = false
				}
			} else if direction == "E" {
				if vertical-1 >= 0 && followedPath[vertical-1][horizontal] == '.' {
					followedPath[vertical-1][horizontal] = 'I'
				}
				if horizontal+1 < maxHorizontal && followedPath[vertical][horizontal+1] == '.' {
					followedPath[vertical][horizontal+1] = 'I'
				}
				vertical++
				direction = "S"
				if vertical < maxVertical {
					exit = false
				}
			}
		case 'F':
			if direction == "W" {
				if vertical-1 >= 0 && followedPath[vertical-1][horizontal] == '.' {
					followedPath[vertical-1][horizontal] = 'O'
				}
				if horizontal-1 >= 0 && followedPath[vertical][horizontal-1] == '.' {
					followedPath[vertical][horizontal-1] = 'O'
				}
				vertical++
				direction = "S"
				if vertical < maxVertical {
					exit = false
				}
			} else if direction == "N" {
				if vertical-1 >= 0 && followedPath[vertical-1][horizontal] == '.' {
					followedPath[vertical-1][horizontal] = 'I'
				}
				if horizontal-1 >= 0 && followedPath[vertical][horizontal-1] == '.' {
					followedPath[vertical][horizontal-1] = 'I'
				}
				horizontal++
				direction = "E"
				if horizontal < maxHorizontal {
					exit = false
				}
			}
		case 'S':
			//fmt.Println("looped!")
			output++
			output /= 2
		case '.':
		default:
		}

		if exit {
			for _, row := range followedPath {
				inside := false
				for _, c := range row {
					switch c {
					case 'I':
						inside = false
					case 'O':
						inside = true
						output++
					case '.':
						if inside {
							output++
						}
					}
				}
			}
			for _, path := range followedPath {
				fmt.Println(string(path))
			}
			return output
		} else {
			//output++
			//fmt.Println("I", output)
		}
	}
}
