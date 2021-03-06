package numwords

import (
	"strconv"
)

// Letters takes in the word representation of a number and returns the number of letters used
func Letters(word string) int {
	count := 0
	for _, b := range word {
		if b == rune(' ') || b == rune('-') {
			continue
		}
		count++
	}
	return count
}

// Word returns the given number's word form
func Word(n int) string {
	word := ""
	nstr := strconv.Itoa(n)

	for len(nstr) > 0 {
		if len(nstr) > 3 {
			word += thousands(nstr[0])
			nstr = "" // thousands will not work above 1000
		} else if len(nstr) == 3 {
			word += hundreds(nstr)
			nstr = ""
		} else if len(nstr) > 1 {
			word += tens(nstr)
			nstr = ""
		} else {
			word += ones(nstr[0])
			nstr = ""
		}
	}

	return word
}

func hundreds(s string) string {
	word := ones(s[0]) + " hundred"
	if s[1:] != "00" {
		word += " and " + tens(s[1:])
	}
	return word
}

func ones(s byte) string {
	if s == byte('1') {
		return "one"
	} else if s == byte('2') {
		return "two"
	} else if s == byte('3') {
		return "three"
	} else if s == byte('4') {
		return "four"
	} else if s == byte('5') {
		return "five"
	} else if s == byte('6') {
		return "six"
	} else if s == byte('7') {
		return "seven"
	} else if s == byte('8') {
		return "eight"
	} else if s == byte('9') {
		return "nine"
	}
	return ""
}

func tens(s string) string {
	snum := ""
	if s[0] == byte('0') {
		return ones(s[1])
	} else if s == "10" {
		return "ten"
	} else if s == "11" {
		return "eleven"
	} else if s == "12" {
		return "twelve"
	} else if s == "13" {
		return "thirteen"
	} else if s == "15" {
		return "fifteen"
	} else if s == "18" {
		return "eighteen"
	} else if s[0] == byte('1') {
		return ones(s[1]) + "teen"
	} else if s[0] == byte('2') {
		snum = "twenty"
	} else if s[0] == byte('3') {
		snum = "thirty"
	} else if s[0] == byte('4') {
		snum = "forty"
	} else if s[0] == byte('5') {
		snum = "fifty"
	} else if s[0] == byte('6') {
		snum = "sixty"
	} else if s[0] == byte('7') {
		snum = "seventy"
	} else if s[0] == byte('8') {
		snum = "eighty"
	} else if s[0] == byte('9') {
		snum = "ninety"
	}

	if s[1] != byte('0') {
		snum += "-" + ones(s[1])
	}

	return snum
}

func thousands(s byte) string {
	return string(ones(s)) + " thousand"
}
