package day01

import (
	"strings"
	"unicode"
)

func getNumberStrings() []string {
	return []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
}

func Part1(input string) int {
	substrings := strings.Split(input, "\n")
	sum := 0
	for _, calibration_value := range substrings {
		calibration_value = strings.Trim(calibration_value, " ")
		reversed := Reverse(calibration_value)
		sum += FindInt(calibration_value, reversed)
	}
	return sum
}

func Part2(input string) int {
	substrings := strings.Split(input, "\n")
	sum := 0
	for _, calibration_value := range substrings {
		calibration_value = strings.Trim(calibration_value, " ")
		reversed := Reverse(calibration_value)
		sum += FindInt2(calibration_value, reversed)
	}
	return sum
}

func FindInt2(fwd string, rev string) int {
	first := 0
	last := 0
	var found bool

	var numStr strings.Builder

	//	fmt.Println("Forward: ", fwd, "Reverse: ", rev)

	// Find first digit
	found = false
	numStr.Reset()
	for _, f := range fwd {
		if unicode.IsDigit(f) {
			first = int(f) - 48
			//			fmt.Println("Got first digit: ", first)
			break
		} else {
			numStr.WriteString(string(f))
			for idx, num := range getNumberStrings() {
				if strings.Contains(numStr.String(), num) {
					//					fmt.Println("Got first digit (number): ", numStr.String(), "SubString: ", num, "Index: ", idx)
					first = idx
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	// Find last digit
	found = false
	numStr.Reset()
	for _, r := range rev {
		if unicode.IsDigit(r) {
			last = int(r) - 48
			//			fmt.Println("Got second digit: ", last)
			break
		} else {
			numStr.WriteString(string(r))
			for idx, num := range getNumberStrings() {
				if strings.Contains(numStr.String(), Reverse(num)) {
					//					fmt.Println("Got second digit (number): ", numStr.String(), "SubString: ", Reverse(num), "Index: ", idx)
					last = idx
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	//	fmt.Println("First: ", first, "Last: ", last)
	//	fmt.Println("----------------------------------")
	return first*10 + last
}

func FindInt(fwd string, rev string) int {
	first := 0
	last := 0

	// Find first digit
	for _, f := range fwd {
		if unicode.IsDigit(f) {
			first = int(f) - 48
			break
		}
	}

	// Find last digit
	for _, r := range rev {
		if unicode.IsDigit(r) {
			last = int(r) - 48
			break
		}
	}

	return first*10 + last
}

func Reverse(input string) string {
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	return output
}
