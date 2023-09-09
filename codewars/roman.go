package codewars

import "fmt"

var mapping = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func ToRoman(in int) string {

	return ""
}

func FromRoman(in string) int {
	chars := make([]string, 0, len(in))
	for _, r := range in {
		chars = append(chars, string(r))
	}

	var out int
	for i := len(chars) - 1; i >= 0; i-- {
		fmt.Printf("chars[%d]: %v\n", i, chars[i])

		// if last elem
		if i == len(chars)-1 {
			out += mapping[chars[i]]
			continue
		}

		// if just add
		prev := mapping[chars[i+1]]
		curr := mapping[chars[i]]
		if curr >= prev {
			out += curr
			continue
		}

		// if curr larger than prev ->
		out = out - prev + (curr - prev)
	}

	return out
}
