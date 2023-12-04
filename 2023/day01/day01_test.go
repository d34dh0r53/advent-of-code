package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1(`1abc2 
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)
	want := 142
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)
	want := 281
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}
