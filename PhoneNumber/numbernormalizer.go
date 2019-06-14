package main

import "strings"

func normalizer(input []string) map[string]int {
	numbers := map[string]int{}
	for _, element := range input {
		element = strings.ReplaceAll(element, " ", "")
		element = strings.ReplaceAll(element, "-", "")
		element = strings.ReplaceAll(element, "(", "")
		element = strings.ReplaceAll(element, ")", "")
		numbers[element]++
	}
	return numbers
}

func main() {

}
