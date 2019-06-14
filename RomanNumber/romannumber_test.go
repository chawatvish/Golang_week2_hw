package main

import (
	"fmt"
	"testing"
)

type testValue struct {
	input  int
	output string
}

func TestRomanNumber(t *testing.T) {
	data := []testValue{{1, "I"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{15, "XV"},
		{49, "XLIX"},
		{50, "L"},
		{70, "LXX"},
	}

	for _, element := range data {
		t.Run(fmt.Sprintf("number = %d", element.input), func(t *testing.T) {
			if romanNumber(element.input) != element.output {
				t.Errorf("romanNumber(%d) != %s", element.input, element.output)
			}
		})
	}
}
