package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardConvert = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type hand struct {
	cards    []int
	cardsMap map[int]int
	bid      int
	highCard int
	rank     int
}

func main() {
	f, err := os.Open("day07/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	hands := []hand{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			hand := hand{
				cards:    []int{},
				cardsMap: map[int]int{},
			}
			parts := strings.Split(text, " ")
			for _, cardRune := range parts[0] {
				card := cardConvert[cardRune]
				if card > hand.highCard {
					hand.highCard = card
				}
				hand.cards = append(hand.cards, card)
				hand.cardsMap[card]++
			}
			hand.bid, err = strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			switch {
			case fiveOfAKind(hand):
				hand.rank = 6
			case fourOfAKind(hand):
				hand.rank = 5
			case fullHouse(hand):
				hand.rank = 4
			case threeOfAKind(hand):
				hand.rank = 3
			case twoPair(hand):
				hand.rank = 2
			case onePair(hand):
				hand.rank = 1
			default:
				hand.rank = 0
			}
			hands = append(hands, hand)
		}
	}
	// sorting ascending
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j]) < 0
	})
	for i := 0; i < len(hands); i++ {
		//fmt.Println(hands[i])
		output += hands[i].bid * (i + 1)
	}
	fmt.Println(output)
}

func fiveOfAKind(hand hand) bool {
	if hand.cardsMap[1] == 5 {
		return true
	}
	for card, cardCount := range hand.cardsMap {
		if card == 1 {
			continue
		}
		if cardCount+hand.cardsMap[1] == 5 {
			return true
		}
	}
	return false
}

func fourOfAKind(hand hand) bool {
	if hand.cardsMap[1] == 4 {
		return true
	}
	for card, cardCount := range hand.cardsMap {
		if card == 1 {
			continue
		}
		if cardCount+hand.cardsMap[1] == 4 {
			return true
		}
	}
	return false
}

func fullHouse(hand hand) bool {
	allocateJokers(&hand, 3)
	for _, cardCount := range hand.cardsMap {
		if cardCount == 3 {
			for _, cardCount := range hand.cardsMap {
				if cardCount == 2 {
					return true
				}
			}
		}
	}
	return false
}

func twoPair(hand hand) bool {
	allocateJokers(&hand, 2)
	for val, cardCount := range hand.cardsMap {
		if cardCount == 2 {
			for val2, cardCount := range hand.cardsMap {
				if cardCount == 2 && val != val2 {
					return true
				}
			}
		}
	}
	return false
}

func threeOfAKind(hand hand) bool {
	allocateJokers(&hand, 2)
	for _, cardCount := range hand.cardsMap {
		if cardCount == 3 {
			return true
		}
	}
	return false
}

func onePair(hand hand) bool {
	allocateJokers(&hand, 2)
	for _, cardCount := range hand.cardsMap {
		if cardCount == 2 {
			return true
		}
	}
	return false
}

func allocateJokers(hand *hand, maxSize int) {
	for i := 0; i < hand.cardsMap[1]; i++ {
		biggestCard := 0
		biggestCardCount := 0
		for card, cardCount := range hand.cardsMap {
			if biggestCardCount < cardCount && card != 1 && cardCount <= maxSize {
				biggestCard = card
				biggestCardCount = cardCount
			}
		}
		hand.cardsMap[biggestCard]++
	}
	hand.cardsMap[1] = 0
}

func compareHands(hand1, hand2 hand) int {
	if hand1.rank > hand2.rank {
		return 1
	} else if hand1.rank < hand2.rank {
		return -1
	} else {
		for i := 0; i < 5; i++ {
			if hand1.cards[i] > hand2.cards[i] {
				return 1
			} else if hand1.cards[i] < hand2.cards[i] {
				return -1
			}
		}
	}
	return 0
}
