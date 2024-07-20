## Algorithm Practice - Golang

Solution of simple algorithms I felt indebted to.

### 1. Single Number
Given a non-empty array of integers where every element appears twice except for one. Find that single one. 
```
Input: [4, 1, 2, 1, 2]
Output: 4
```

### 2. Move Zeroes
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements. 
```
Input: [0, 1, 0, 3, 12] 
Output: [1, 3, 12, 0, 0]
```

### 3. Majority Element 
Given an array of size n, find the majority element.  
```
Input: [3, 2, 3] 
Output: 3
```

### 4. Group Anagrams
Given an array of strings, group the anagrams together. An anagram is a word formed by repositioning the letters of another word.
```
Input: ["eat", "tea", "tan", "ate", "nat", "bat"]
Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
```

### 5. Reverse Words in a String
Given a string, reverse the order of the words.
```
Input: "  Go  is  awesome   "
Output: "awesome is Go"
```

### 6. Array Rotation
Rotate an array to the right by a given number of steps.
```
Input: [1, 2, 3, 4, 5, 6, 7], steps = 3
Output after rotations: [5, 6, 7, 1, 2, 3, 4]
```

## What I Learned Today

- **XOR Operator:** Useful for bitwise operations, including checking if two numbers are the same or finding unique elements in arrays of int's.
- **Updating Arrays During Iteration:** It's risky to modify an array's length while iterating over it. It can lead to unexpected behavior or bugs.
- **Boyer-Moore Voting Algorithm:** A more efficient algorithm for finding the majority element in an array.
- **`strings.Builder`:** A utility in Go for efficiently building strings through concatenation, which helps in performance and memory usage (it's better than the `fmt.Sprintf()`).

## What I Would Like to Explore Next

- **Implementing the Boyer-Moore Voting Algorithm:** Apply the algorithm to practice and understand its application in a more complex case.
- **Advanced Bitwise Operations:** Go deeper into real-world apps where the XOR operator and other bitwise operations can help us with.
- **Command-Line Tools:** Create command-line tools for running the algorithms to play with the Go's `flag` package.
- **Performance Optimization:** Explore techniques for optimizing Go code after you solve any problem.

- **Complete the Unit Tests for these algorithms:** Write the unit tests for all algorithms ensuring edge cases.