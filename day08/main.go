package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	name        string
	leftString  string
	rightString string
	left        *node
	right       *node
}

func main() {
	f, err := os.Open("day08/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	scanner.Scan()
	text := scanner.Text()
	input := text
	nodes := map[string]*node{}
	for scanner.Scan() {
		text = scanner.Text()
		if text == "" {
			continue
		} else {
			newNode := node{
				name:        text[0:3],
				leftString:  text[7:10],
				rightString: text[12:15],
			}
			nodes[newNode.name] = &newNode
		}
	}

	currentNodes := []*node{}
	for nodeName, node := range nodes {
		node.left = nodes[node.leftString]
		node.right = nodes[node.rightString]
		if nodeName[2] == 'A' {
			currentNodes = append(currentNodes, node)
		}
	}
	paths := map[int]int{}
	for _, node := range currentNodes {
		path := getPath(node, input)
		paths[path]++
		output = path
	}
	for path := range paths {
		fmt.Println(path)
		output = lcm(output, path)
	}

	fmt.Println(output)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func getPath(node *node, input string) int {
	output := 0
	for {
		for _, direction := range input {
			if direction == 'R' {
				node = node.right
			} else if direction == 'L' {
				node = node.left
			}
			output++
			if node.name[2] == 'Z' {
				return output
			}
		}
	}
}
