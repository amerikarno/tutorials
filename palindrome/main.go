package main

import "fmt"

func main() {
	a := "zsdt"
	// b := palindrome(a)
	c := isPanlindrome(a)
	if c {
		fmt.Printf("This %v is palindrome",a)
	} else {
		fmt.Printf("This %v isn't palindrome",a)
	}
}

func palindrome(word string) []byte {
	var names []byte
	names = []byte(word)

	for i, j := 0, len(names)-1; i < j; i, j = i+1, j-1 {
		names[i], names[j] = names[j], names[i]
	}

	return names
}

func isPanlindrome(word string) bool {
	p := palindrome(word)
	if word == string(p) {
		return true
	}
	return false
}
