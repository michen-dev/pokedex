package main

import (
	"strings"
)

func cleanInput(text string) []string {
	res := []string{}
	start := -1
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			if start != -1 {
				res = append(res, strings.ToLower(text[start:i]))
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}
	if start != -1 {
		res = append(res, strings.ToLower(text[start:]))
	}
	return res
}