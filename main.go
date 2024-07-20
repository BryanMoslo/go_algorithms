package main

import (
	"academic_algorithms/algorithms"
	"fmt"
)

func main() {
	fmt.Println("1. Single Number:", algorithms.FindSimgleElement([]int{5, 3, 5, 3, 7, 9, 6, 9, 6}))
	fmt.Println("2. Moving Zeroes:", algorithms.MoveZeroes([]int{0, 0, 1, 2, 3}))

	n, times := algorithms.MajorityElement([]int{4, 1, 4, 3, 4, 12})
	fmt.Printf("3. Majority Element %v appears %v times \n", n, times)

	fmt.Println("4. Group Anagrams:", algorithms.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println("5. Reverse words:", algorithms.ReverseWords("   Go   is   Awesome"))
	fmt.Println("5. Reverse words with Builder:", algorithms.ReverseWordsWithBuilder("   Go   is   Awesome"))
	fmt.Println("6. Rotate Array n steps", algorithms.RotateList([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}
