package main

import (
	"fmt"
	"strings"
)

// create a function that returns dictionary
func wordFreq(s string) map[string]int {

	// if s is none
	if len(s) == 0 {
		return nil
	}

	// create dictionary
	my_dict := make(map[string]int)

	//handle case-sensetivity
	s = strings.ToLower(s)

	// keep only letters and spaces
	var sentence string
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || ch == ' ' {
			sentence += string(ch)
		}
	}

	// Add the words into the dictionary
	words := strings.Fields(sentence)
	for _, word := range words {
		my_dict[word] += 1
	}

	return my_dict
}

func palindrom(s string) bool {

	s = strings.ToLower(s)

	var word string
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			word += string(ch)
		}
	}

	// check palindrom
	left := 0
	right := len(word) - 1

	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true

}

func main() {
	text := "Hello, hello how are you"
	result := wordFreq(text)

	for k, v := range result {
		fmt.Println(k, v)
	}

	p_word := "madam"
	p2_word := "hell0"

	fmt.Println(palindrom(p_word))
	fmt.Println(palindrom(p2_word))

}
