package numwords

import (
	"testing"
)

func TestLetters(t *testing.T) {
	var word string
	var count, exp int

	word = Word(1)
	count = Letters(word)
	exp = 3
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word += Word(2)
	word += Word(3)
	word += Word(4)
	word += Word(5)
	count = Letters(word)

	exp = 19
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(127)
	count = Letters(word)
	exp = 24
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(610)
	count = Letters(word)
	exp = 16
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(812)
	count = Letters(word)
	exp = 5 + 7 + 3 + 6
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(13)
	count = Letters(word)
	exp = 8
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(38)
	count = Letters(word)
	exp = 6 + 5
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(38)
	count = Letters(word)
	exp = 6 + 5
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(41)
	count = Letters(word)
	exp = 5 + 3
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(51)
	count = Letters(word)
	exp = 5 + 3
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(61)
	count = Letters(word)
	exp = 5 + 3
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(991)
	count = Letters(word)
	exp = 4 + 7 + 3 + 6 + 3
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(30)
	count = Letters(word)
	exp = 6
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(342)
	count = Letters(word)
	exp = 23
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}

	word = Word(115)
	count = Letters(word)
	exp = 20
	if count != exp {
		t.Errorf("%s has %d letters, but got %d", word, exp, count)
	}
}

func TestWord(t *testing.T) {
	var num int
	var word, exp string

	num = 2
	word = Word(num)
	exp = "two"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	num = 11
	word = Word(num)
	exp = "eleven"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	num = 15
	word = Word(num)
	exp = "fifteen"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	num = 27
	word = Word(num)
	exp = "twenty-seven"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	num = 179
	word = Word(num)
	exp = "one hundred and seventy-nine"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	// num = 1582
	// word = Word(num)
	// exp = "one thousand five hundred and eighty-two"
	// if word != exp {
	// 	t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	// }

	num = 1000
	word = Word(num)
	exp = "one thousand"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	// test all the ones-matched teens
	num = 14
	word = Word(num)
	num = 16
	word += Word(num)
	num = 17
	word += Word(num)
	num = 18
	word += Word(num)
	num = 19
	word += Word(num)
	exp = "fourteensixteenseventeeneighteennineteen"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	// ensure no "and" follows a hundred if theres no ten or one
	num = 100
	word = Word(num)
	exp = "one hundred"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	// ensure there's no hyphen if a ten is not followed by a one
	num = 80
	word = Word(num)
	exp = "eighty"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}

	// ensure ones without a ten calculate properly
	num = 501
	word = Word(num)
	exp = "five hundred and one"
	if word != exp {
		t.Errorf("%d's word is \"%s\", not \"%s\"", num, exp, word)
	}
}
