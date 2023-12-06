package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time     int
	distance int
}

func main() {
	f, err := os.Open("day06/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	scanner.Scan()
	text := scanner.Text()
	text = strings.ReplaceAll(text, " ", "")
	time, err := strconv.Atoi(strings.Split(text, ":")[1])
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	text = scanner.Text()
	text = strings.ReplaceAll(text, " ", "")
	distance, err := strconv.Atoi(strings.Split(text, ":")[1])
	if err != nil {
		panic(err)
	}
	for i := 0; i < time; i++ {
		if i*(time-i) > distance {
			output++
		}
	}

	fmt.Println(output)
}
