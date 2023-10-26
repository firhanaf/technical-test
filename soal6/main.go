package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
			res -= roman[s[i]]
		} else {
			res += roman[s[i]]
		}
	}
	return res
}

func main() {
	file, err := os.Open("roman.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outputFile, err := os.Create("numeric_values.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		romanNumeral := strings.TrimSpace(scanner.Text())
		intValue := romanToInt(romanNumeral)
		_, err := fmt.Fprintln(outputFile, intValue)
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
