package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	m := make(map[string]int)
	for _, el := range fields {
		m[el] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
