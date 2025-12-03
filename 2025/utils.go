package main

import (
	"strconv"
	"strings"
)

func ParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func SplitLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
