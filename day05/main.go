package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var almanacs = map[string][]almanacRange{}

type almanacRange struct {
	origin      int
	destination int
	size        int
}

func main() {
	f, err := os.Open("day05/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0

	seeds := []almanacRange{}
	scanner.Scan()
	text := scanner.Text()
	parts := strings.Split(text, " ")
	seed := almanacRange{}
	for i, part := range parts {
		if i == 0 {
			continue
		}
		partNumber, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		if i%2 == 1 {
			seed.origin = partNumber
			seed.destination = partNumber
		} else {
			seed.size = partNumber
			seeds = append(seeds, seed)
		}
	}
	// skip the first blank line
	scanner.Scan()
	valKey := "seed"
	printCurrent(seeds, valKey)
	almanac := []almanacRange{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			appendAlmanac(almanac, valKey)
			seeds = calculateRange(seeds, valKey)
			// fmt.Println()
			// fmt.Println(valKey)
			// printCurrent(almanacs[valKey], "almanac")
			// printCurrent(seeds, "seed")
			almanac = []almanacRange{}
			continue
		}

		if !unicode.IsDigit(rune(text[0])) {
			re := regexp.MustCompile(`to-(\w+)`)
			match := re.FindStringSubmatch(text)
			valKey = match[1]
			continue
		}

		parts := strings.Split(text, " ")
		destination, _ := strconv.Atoi(parts[0])
		origin, _ := strconv.Atoi(parts[1])
		size, _ := strconv.Atoi(parts[2])
		almanac = append(almanac, almanacRange{
			origin:      origin,
			destination: destination,
			size:        size,
		})
	}
	appendAlmanac(almanac, valKey)
	seeds = calculateRange(seeds, valKey)
	// fmt.Println()
	// fmt.Println(valKey)
	// printCurrent(almanacs[valKey], "almanac")
	// printCurrent(seeds, "seed")

	for _, seedMap := range seeds {
		// fmt.Println(seedMap)
		loc := seedMap.destination
		if loc < output || output == 0 {
			output = loc
		}
	}

	fmt.Println(output)
}

func appendAlmanac(almanac []almanacRange, key string) {
	// Sort the almanacs by xStart
	for i := 0; i < len(almanac); i++ {
		for j := i + 1; j < len(almanac); j++ {
			if almanac[i].origin > almanac[j].origin {
				almanac[i], almanac[j] = almanac[j], almanac[i]
			}
		}
	}
	almanacs[key] = almanac
}

func calculateRange(seeds []almanacRange, key string) []almanacRange {
	output := []almanacRange{}
	for _, seed := range seeds {
		found := false
		for _, almanac := range almanacs[key] {
			if seed.destination+seed.size < almanac.origin || seed.destination > almanac.origin+almanac.size {
				// Seed range either fully before or fully after the almanac range
				continue
			}
			if seed.destination < almanac.origin {
				// yStart falls between the previous almanac and this one, so is a 1:1 mapping
				newSize := almanac.origin - seed.destination
				newSeed := almanacRange{
					origin:      seed.origin,
					destination: seed.destination,
					size:        newSize,
				}
				output = append(output, newSeed)
				seed = almanacRange{
					origin:      seed.origin + newSize,
					destination: seed.destination + newSize,
					size:        seed.size - newSize,
				}
			}

			if seed.destination+seed.size > almanac.origin+almanac.size {
				// seed needs to split, first half is saved recorded
				newSize := almanac.origin + almanac.size - seed.destination
				newSeed := almanacRange{
					origin:      seed.origin,
					destination: seed.destination - almanac.origin + almanac.destination,
					size:        newSize,
				}
				if newSeed.size < 0 || newSeed.destination < 0 {
					fmt.Printf("%v, %+v, %+v, %+v\n", newSize, newSeed, seed, almanac)
					panic("oof split min")
				}
				output = append(output, newSeed)
				seed = almanacRange{
					origin:      seed.origin + newSize,
					destination: seed.destination + newSize,
					size:        seed.size - newSize,
				}
			} else {
				found = true
				output = append(output, almanacRange{
					origin:      seed.origin,
					destination: seed.destination - almanac.origin + almanac.destination,
					size:        seed.size,
				})
				break
			}
		}
		if !found {
			output = append(output, almanacRange{
				origin:      seed.origin,
				destination: seed.destination,
				size:        seed.size,
			})
		}
	}
	return output
}

func printCurrent(seeds []almanacRange, key string) {
	fmt.Printf("Current %v\n", key)
	for _, seed := range seeds {
		fmt.Printf("%+v\n", seed)
	}

}

//318026523 too high
