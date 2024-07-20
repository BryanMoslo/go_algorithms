package algorithms

import (
	"testing"
)

func TestFindSimgleElement(t *testing.T) {
	tcases := []struct {
		Input    []int
		Expected int
	}{
		{
			Input:    []int{3, 3, 5, 6, 7, 6, 7, 9, 8, 8, 9, 3, 3},
			Expected: 5,
		},
		{
			Input:    []int{1, 1, 2, 2, 3},
			Expected: 3,
		},
		{
			Input:    []int{4, 5, 6, 7, 7, 6, 5},
			Expected: 4,
		},
		{
			Input:    []int{10, 10, 20, 20, 30, 30, 40},
			Expected: 40,
		},
		{
			Input:    []int{100, 200, 100, 300, 200},
			Expected: 300,
		},
		{
			Input:    []int{100, 200, 100, 200},
			Expected: 0,
		},
	}

	for caseIndex, tcase := range tcases {
		result := FindSimgleElement(tcase.Input)
		if result != tcase.Expected {
			t.Fatalf(`Case: %v - got: %v, expected: %v`, caseIndex, result, tcase.Expected)
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	tcases := []struct {
		Input    []int
		Expected []int
	}{
		{
			Input:    []int{0, 1, 0, 3, 12},
			Expected: []int{1, 3, 12, 0, 0},
		},
		{
			Input:    []int{0, 0, 1, 2, 3},
			Expected: []int{1, 2, 3, 0, 0},
		},
		{
			Input:    []int{4, 2, 0, 1, 0, 3},
			Expected: []int{4, 2, 1, 3, 0, 0},
		},
		{
			Input:    []int{1, 2, 3, 4, 5},
			Expected: []int{1, 2, 3, 4, 5},
		},
		{
			Input:    []int{0, 0, 0, 0, 0},
			Expected: []int{0, 0, 0, 0, 0},
		},
	}

	for caseIndex, tcase := range tcases {
		result := MoveZeroes(tcase.Input)

		if len(result) != len(tcase.Expected) {
			t.Fatalf(`Length is not the same. got: %v, expected: %v`, len(result), len(tcase.Expected))
		}

		for i := range result {
			if result[i] != tcase.Expected[i] {
				t.Fatalf(`Case: %v - Wrong value in index: %v. Got: %v, expected: %v`, caseIndex, i, result, tcase.Expected)
			}
		}
	}
}

func TestMajorityElement(t *testing.T) {
	tcases := []struct {
		Input         []int
		Expected      int
		ExpectedTimes int
	}{
		{
			Input:         []int{3, 3, 4, 2, 4, 4, 2, 4, 4},
			Expected:      4,
			ExpectedTimes: 5,
		},
		{
			Input:         []int{1, 2, 3, 4, 5},
			Expected:      0, // No majority element
			ExpectedTimes: 0,
		},
		{
			Input:         []int{1, 1, 1, 2, 2},
			Expected:      1,
			ExpectedTimes: 3,
		},
		{
			Input:         []int{5, 5, 5, 2, 2, 2, 5},
			Expected:      5,
			ExpectedTimes: 4,
		},
		{
			Input:         []int{7, 7, 7, 7, 7, 7, 7, 6, 6},
			Expected:      7,
			ExpectedTimes: 7,
		},
	}

	for caseIndex, tcase := range tcases {
		result, times := MajorityElement(tcase.Input)
		if result != tcase.Expected {
			t.Fatalf(`Case: %v - got: %v, expected: %v`, caseIndex, result, tcase.Expected)
		}

		if times != tcase.ExpectedTimes {
			t.Fatalf(`Case: %v - got: %v times, expected: %v times`, caseIndex, result, tcase.Expected)
		}
	}
}

// func TestGroupAnagrams(t *testing.T) {
// 	tcases := []struct {
// 		Input    []string
// 		Expected [][]string
// 	}{
// 		{
// 			Input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
// 			Expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
// 		},
// 		{
// 			Input:    []string{"", ""},
// 			Expected: [][]string{{"", ""}},
// 		},
// 		{
// 			Input:    []string{"a"},
// 			Expected: [][]string{{"a"}},
// 		},
// 		{
// 			Input:    []string{"rat", "tar", "art"},
// 			Expected: [][]string{{"rat", "tar", "art"}},
// 		},
// 		{
// 			Input:    []string{"abc", "bca", "acb", "xyz", "zyx"},
// 			Expected: [][]string{{"abc", "bca", "acb"}, {"xyz", "zyx"}},
// 		},
// 	}

// 	for _, tcase := range tcases {
// 		Requires manual check
// 	}
// }

func TestReverseWords(t *testing.T) {
	tcases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "the sky is blue",
			Expected: "blue is sky the",
		},
		{
			Input:    "hello world",
			Expected: "world hello",
		},
		{
			Input:    "  a good   example ",
			Expected: "example good a",
		},
		{
			Input:    "  Alice  Bob  Charlie ",
			Expected: "Charlie Bob Alice",
		},
		{
			Input:    "singleword",
			Expected: "singleword",
		},
		{
			Input:    "  ",
			Expected: "",
		},
	}

	for caseIndex, tcase := range tcases {
		result := ReverseWords(tcase.Input)
		resultWithBuilder := ReverseWordsWithBuilder(tcase.Input)
		if result != tcase.Expected {
			t.Fatalf(`Case: %v - got: %v, expected: %v`, caseIndex, result, tcase.Expected)
		}

		if resultWithBuilder != tcase.Expected {
			t.Fatalf(`Case: %v - got: %v, expected: %v`, caseIndex, result, tcase.Expected)
		}
	}
}

// func TestRotateList(t *testing.T) {
// 	tcases := []struct {
// 		Input    []int
// 		Steps    int
// 		Expected []int
// 	}{
// 		{
// 			Input:    []int{1, 2, 3, 4, 5, 6, 7},
// 			Steps:    3,
// 			Expected: []int{5, 6, 7, 1, 2, 3, 4},
// 		},
// 		{
// 			Input:    []int{1, 2, 3, 4, 5},
// 			Steps:    2,
// 			Expected: []int{4, 5, 1, 2, 3},
// 		},
// 		{
// 			Input:    []int{10, 20, 30, 40, 50},
// 			Steps:    1,
// 			Expected: []int{50, 10, 20, 30, 40},
// 		},
// 		{
// 			Input:    []int{1, 2, 3, 4},
// 			Steps:    4,
// 			Expected: []int{1, 2, 3, 4},
// 		},
// 		{
// 			Input:    []int{1, 2, 3},
// 			Steps:    0,
// 			Expected: []int{1, 2, 3},
// 		},
// 		{
// 			Input:    []int{1, 2, 3, 4, 5, 6},
// 			Steps:    7,
// 			Expected: []int{6, 1, 2, 3, 4, 5},
// 		},
// 	}

// 	for caseIndex, tcase := range tcases {
// 		result := RotateList(tcase.Input, tcase.Steps)
// 		Requires manual check
// 	}
// }
