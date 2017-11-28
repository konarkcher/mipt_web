package main

import "unicode"
import "strings"

func RemoveEven(input []int) ([]int) {
	odds := make([]int, 0, len(input))

	for _, val := range input {
		if val%2 == 1 {
			odds = append(odds, val)
		}
	}

	return odds
}

func PowerGenerator(number int) (func() int) {
	currentPower := number

	return func() (ret int) {
		ret = currentPower
		currentPower *= number
		return
	}
}

func DifferentWordsCount(s string) int {
	words := make(map[string]bool)
	s += "_"

	wordStart := 0
	for i := 1; i < len(s); i++ {
		if !unicode.IsLetter(rune(s[i-1])) && unicode.IsLetter(rune(s[i])) {
			wordStart = i
		}
		if unicode.IsLetter(rune(s[i-1])) && !unicode.IsLetter(rune(s[i])) {
			words[strings.ToLower(s[wordStart:i])] = true
		}
	}

	return len(words)
}

/*
func main() {
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result) // Должно напечататься [3 5]

	gen := PowerGenerator(2)
	fmt.Println(gen()) // Должно напечатать 2
	fmt.Println(gen()) // Должно напечатать 4
	fmt.Println(gen()) // Должно напечатать 8

	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
}
*/
