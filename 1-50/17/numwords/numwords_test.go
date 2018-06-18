package numwords

import (
	"testing"
)

func TestWord(t *testing.T) {
	var num int
	var word, exp string

	num = 2
	word = Word(num)
	exp = "two"
	if word != exp {
		t.Errorf("%d's word is %s, not %s", num, exp, word)
	}

	num = 11
	word = Word(num)
	exp = "eleven"
	if word != exp {
		t.Errorf("%d's word is %s, not %s", num, exp, word)
	}

	num = 15
	word = Word(num)
	exp = "fifteen"
	if word != exp {
		t.Errorf("%d's word is %s, not %s", num, exp, word)
	}

	num = 27
	word = Word(num)
	exp = "twenty-seven"
	if word != exp {
		t.Errorf("%d's word is %s, not %s", num, exp, word)
	}

	num = 179
	word = Word(num)
	exp = "one hundred and seventy-nine"
	if word != exp {
		t.Errorf("%d's word is %s, not %s", num, exp, word)
	}

	num = 1582
	word = Word(num)
	exp = "one thousand five hundred and eighty-two"
	if word != exp {
		t.Errorf("%d's word is %s, not %s", num, exp, word)
	}
}
