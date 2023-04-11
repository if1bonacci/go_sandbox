package main

import (
	"fmt"
)

const (
	Fizz     string = "FIZZ"
	Buzz     string = "BUZZ"
	FizzBuzz string = "FIZZBUZZ"
)

type CallNumber int64

type Numbers []CallNumber

type GameVal struct {
	numbers Numbers
}

var game GameVal

func (gameVal *GameVal) SetNumbers(data Numbers) {
	(*gameVal).numbers = data
}

type Item interface{}

type ResultMap map[CallNumber]Item

func playGame() {
	resMap := make(ResultMap, len(game.numbers))

	for _, number := range game.numbers {
		if (number%3 == 0) && (number%5 == 0) {
			resMap[number] = FizzBuzz
		} else if number%3 == 0 {
			resMap[number] = Fizz
		} else if number%5 == 0 {
			resMap[number] = Buzz
		} else {
			resMap[number] = number
		}
	}

	fmt.Println(resMap)
}

func main() {
	var numbers Numbers
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 15)
	game.SetNumbers(numbers)

	playGame()
}
