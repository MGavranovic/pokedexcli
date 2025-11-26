package main

import "strings"

func cleanInput(text string) []string {
	cleaned := strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")

	return cleaned
}
