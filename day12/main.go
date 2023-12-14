package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type lineOut struct {
	output int64
	line   int
}
type hash struct {
	currentLen int
	attempt    string
}
type childHash struct {
	currentLen int
	position   int
}

type lineData struct {
	damaged       []string
	residuals     []int
	springMap     string
	springLength  int
	damagedLength int
	attemptChan   chan struct{}
	positionMap   map[hash]bool
	childMap      map[childHash]int64
}

var dots = []string{}

// 4609768071855 high

// 18716325559999
// 18769993049598
func main() {
	begin := time.Now()
	defer func() { fmt.Println(time.Since(begin)) }()

	main2()
	return
	f, err := os.Open("day12/input.txt")
	for i := 0; i < 100; i++ {
		dots = append(dots, strings.Repeat(".", i))
	}

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 0
	var output int64
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			line++
			start := time.Now()
			out := processLine(text, line)
			fmt.Println(line, out, time.Since(start))
			output += out
		}
	}

	fmt.Println(output)
}

func main2() {
	f, err := os.Open("day12/input.txt")
	for i := 0; i < 100; i++ {
		dots = append(dots, strings.Repeat(".", i))
	}

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 0
	lineOutputs := make(chan lineOut, 1000)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			line++
			go func(text string, line int) {
				lineOutput := lineOut{line: line}
				lineOutput.output = processLine(text, line)
				lineOutputs <- lineOutput
			}(text, line)
		}
	}
	var output int64
	for lineOutput := range lineOutputs {
		fmt.Println(time.Now().Format("17:06:04.000000"), line, lineOutput.line, lineOutput.output)
		output += lineOutput.output
		line--
		if line == 0 {
			close(lineOutputs)
		}
	}

	fmt.Println(output)
}

func processLine(text string, line int) int64 {
	parts := strings.Split(text, " ")
	s := &lineData{damaged: []string{}, attemptChan: make(chan struct{}, 1000000), positionMap: map[hash]bool{}, childMap: map[childHash]int64{}}
	s.damaged = []string{}
	s.springMap = ""
	for i := 0; i < 5; i++ {
		s.springMap += parts[0] + "?"
		for _, val := range strings.Split(parts[1], ",") {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			s.damaged = append(s.damaged, strings.Repeat("#", num))
		}
	}
	s.damagedLength = len(s.damaged)
	s.springLength = len(s.springMap) - 1
	s.springMap = s.springMap[:s.springLength]

	s.residuals = make([]int, s.damagedLength)
	residualSum := 0
	for i := s.damagedLength - 1; i >= 0; i-- {
		residualSum += len(s.damaged[i])
		s.residuals[i] = residualSum
		residualSum++
	}
	// fmt.Println(s.springMap)
	return s.generateArraysAsync(0, 0, "")
}

func (s *lineData) generateArraysAsync(currentLen int, position int, current string) int64 {
	child := childHash{currentLen, position}
	childcount, ok := s.childMap[child]
	if ok {
		return childcount
	}
	if position == s.damagedLength {
		remainder := s.springLength - currentLen
		attempt := dots[remainder]
		if remainder == 0 || s.valid(s.springMap[currentLen:], attempt, remainder, hash{currentLen, attempt}) {
			s.childMap[child] = 1
			// fmt.Println(current)
			return 1
		} else {
			s.childMap[child] = 0
			return 0
		}
	}
	start := 0
	if position != 0 {
		start = 1
	}
	var output int64
	for i := start; i <= s.springLength-s.residuals[position]-currentLen; i++ {
		attempt := dots[i] + s.damaged[position]
		key := hash{currentLen, attempt}
		val, ok := s.positionMap[key]
		if ok && !val {
			continue
		}
		attemptLength := len(attempt)
		if !ok && !s.valid(s.springMap[currentLen:attemptLength+currentLen], attempt, attemptLength, key) {
			continue
		}
		output += s.generateArraysAsync(currentLen+attemptLength, position+1, current+attempt)
	}
	s.childMap[child] = output
	return output
}

func (s *lineData) valid(springMap, attempt string, attemptLength int, key hash) bool {
	for i := 0; i < attemptLength; i++ {
		if springMap[i] != '?' && springMap[i] != attempt[i] {
			s.positionMap[key] = false
			return false
		}
	}
	s.positionMap[key] = true
	return true
}
