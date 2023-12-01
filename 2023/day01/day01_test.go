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
