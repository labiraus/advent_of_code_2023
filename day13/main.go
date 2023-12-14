package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("day13/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	rows := []string{}
	columns := []string{}
	line := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			line++
			// if line == 13 || line == 6 {
			out := sumEvaluate(rows, columns)
			fmt.Println(line, out, "done")
			output += out
			// 	if line == 13 {
			// 		return
			// 	}
			// }
			rows = []string{}
			columns = []string{}
		} else {
			if len(columns) == 0 {
				for i := 0; i < len(text); i++ {
					columns = append(columns, "")
				}
			}
			rows = append(rows, text)
			for i, char := range text {
				columns[i] += string(char)
			}
		}
	}
	output += sumEvaluate(rows, columns)

	fmt.Println(output)
}

func sumEvaluate(rows []string, columns []string) int {
	out := evalute(rows) * 100
	if out == 0 {
		out = evalute(columns)
	}
	return out
}

func evalute(rows []string) int {
	for i := 0; i < len(rows)-1; i++ {
		if ok, _ := compare(rows[i], rows[len(rows)-1], 0); ok {
			j := (len(rows) + i) / 2
			// fmt.Println("checking")
			if check(rows, j) {
				return j
			}
		}
	}
	// fmt.Println("two")
	for i := 1; i < len(rows); i++ {
		// fmt.Println(rows[i], rows[0])
		if ok, _ := compare(rows[i], rows[0], 0); ok {
			// fmt.Println(rows[i], rows[0])
			j := i/2 + 1
			if check(rows, j) {
				return j
			}
		}
	}
	// fmt.Println("nothing")

	return 0
}

//.....##..#.

func compare(row1 string, row2 string, mistake int) (bool, int) {
	for i := 0; i < len(row1); i++ {
		if row1[i] != row2[i] {
			mistake++
			if mistake > 1 {
				return false, mistake
			}
		}
	}
	return true, mistake
}

func check(rows []string, mid int) bool {
	mistake := 0
	ok := false
	if mid < len(rows)/2 {
		// fmt.Println("here", (mid)*2-2)
		for i := 0; i < mid*2-i-1; i++ {
			// fmt.Println(rows[i], rows[mid*2-i-1], mistake)
			if ok, mistake = compare(rows[i], rows[mid*2-i-1], mistake); !ok {
				// fmt.Println(rows[i], rows[mid*2-i-1], mistake, "out")
				return false
			}
		}
	} else {
		for i := mid*2 - len(rows); i < mid; i++ {
			if i < 0 {
				continue
			}
			if ok, mistake = compare(rows[i], rows[2*mid-i-1], mistake); !ok {
				// fmt.Println(rows[i], rows[mid*2-i-1], mistake, "out")
				return false
			}
			// fmt.Println(rows[i], rows[mid*2-i-1], mistake)
		}
	}
	// fmt.Println(mid, mistake)
	return mistake == 1
}

// 51481 high
// 27131 low
// 16129 low
