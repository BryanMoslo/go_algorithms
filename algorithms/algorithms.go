package algorithms

import (
	"fmt"
	"sort"
	"strings"
)

func FindSimgleElement(list []int) int {
	result := 0
	for _, current := range list {
		result ^= current // XOR Bitwise Operator 👀 - This is doing an XOR on every bits for each number. In other words, this is doing the "magic trick" of guessing the number by adding and subtracting numbers.
	}
	return result
}

func MoveZeroes(list []int) []int {
	nonZeroes := []int{}
	zeroes := []int{}

	for _, current := range list {
		if current == 0 {
			zeroes = append(zeroes, current)
			continue
		}

		nonZeroes = append(nonZeroes, current)
	}

	return append(nonZeroes, zeroes...)
}

func MajorityElement(list []int) (int, int) {
	var result int
	var coincidences int

	for _, current := range list {
		count := 0
		for _, aux := range list { // Could be inefficient with really big lists, TODO: Boyer-Moore Voting Algorithm can fix it.
			if current == aux {
				count++
			}
		}

		if coincidences < count {
			result = current
			coincidences = count
		}
	}

	if len(list) > 1 && coincidences == 1 {
		return 0, 0 // There is not a Majority element
	}

	return result, coincidences
}

func GroupAnagrams(list []string) [][]string {
	anagramMap := map[string][]string{}

	for _, current := range list {
		aux := strings.Split(current, "")
		sort.Strings(aux)
		key := strings.Join(aux, "")

		anagramMap[key] = append(anagramMap[key], current)
	}

	var result [][]string
	for _, anagrams := range anagramMap {
		result = append(result, anagrams)
	}

	return result
}

func ReverseWords(value string) string {
	var result string
	s := strings.TrimSpace(value)
	list := strings.Fields(s)

	for i := len(list) - 1; i >= 0; i-- {
		result = fmt.Sprintf("%s %s", result, list[i])
	}

	return strings.TrimSpace(result)
}

func ReverseWordsWithBuilder(value string) string {
	var result strings.Builder
	s := strings.TrimSpace(value)
	list := strings.Fields(s)

	for i := len(list) - 1; i >= 0; i-- {
		if result.Len() > 0 {
			result.WriteString(" ")
		}

		result.WriteString(list[i])
	}

	return result.String()
}

func RotateList(list []int, steps int) []int {
	limit := len(list) - steps
	return append(list[limit:], list[0:limit]...)
}
