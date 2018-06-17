package longmath

import (
	"bytes"
	"strconv"
)

func digit(s string, i int) byte {
	b := byte('0')
	if i < len(s) {
		b = s[i]
	}
	return b
}

// LongAdd performs long addition on the numbers given in string form
func LongAdd(n1, n2 string) string {
	var res bytes.Buffer // overall result of long addition
	var sum byte         // temporary sum of digits
	carry := byte('0')   // any carried digit resulting from a sum of digits

	size := len(n1)
	if len(n2) > size {
		size = len(n2)
	}

	n1 = reverse(n1)
	n2 = reverse(n2)

	// perform long addition digit-by-digit
	for i := 0; i < size; i++ {
		sum, carry = longAddDigit(digit(n1, i), digit(n2, i), carry)
		res.WriteByte(sum)
	}
	if carry != byte('0') {
		res.WriteByte(carry)
	}

	return reverse(res.String())
}

func longAddDigit(d1, d2, d3 byte) (byte, byte) {
	// perform addition on the arguments and
	n1, _ := strconv.Atoi(string(d1))
	n2, _ := strconv.Atoi(string(d2))
	n3, _ := strconv.Atoi(string(d3))
	sum := strconv.Itoa(n1 + n2 + n3)

	// determine the value of the carry
	carry := byte('0')
	if len(sum) > 1 {
		carry = sum[0]
		sum = sum[1:]
	}

	return sum[0], carry
}

func reverse(s string) string {
	var rev bytes.Buffer

	for i := len(s) - 1; i > -1; i-- {
		rev.WriteByte(s[i])
	}

	return rev.String()
}
