package longmath

import (
	"strconv"
)

// LongAdd performs long addition on the numbers given in string form
func LongAdd(n1, n2 string) string {
	var res []byte     // overall result of long addition
	var sum byte       // temporary sum of digits
	carry := byte('0') // carried digit resulting from a sum of digits

	// convert the numbers to []byte equivalents
	b1, b2 := paddedBytes(n1, n2)
	size := len(b1)

	// perform long addition digit-by-digit
	for i := size - 1; i >= 0; i-- {
		sum, carry = longAddDigit(b1[i], b2[i], carry)
		res = append([]byte{sum}, res...)
	}
	// if the last long addition resulted in a carry, ensure it is added to the final sum
	if carry != byte('0') {
		res = append([]byte{carry}, res...)
	}

	return string(res)
}

// longAddDigit performs long addition between 2 primary digits and a carry
func longAddDigit(d1, d2, d3 byte) (byte, byte) {
	// perform addition on the arguments and
	n1, _ := strconv.Atoi(string(d1))
	n2, _ := strconv.Atoi(string(d2))
	n3, _ := strconv.Atoi(string(d3)) // the carry
	sum := strconv.Itoa(n1 + n2 + n3)

	// determine the value of the carry
	carry := byte('0')
	if len(sum) > 1 {
		carry = sum[0]
		sum = sum[1:]
	}

	return sum[0], carry
}

// LongMult performs long multiplication on the numbers given in string form
func LongMult(n1, n2 string) string {
	res := "0"         // overall result of long multiplication
	var prod byte      // intermediate product of digits
	carry := byte('0') // carried digit resulting from a sum of digits

	var partials [][]byte // all partial products to sum

	cand, lier := []byte(n1), []byte(n2)
	// ensure the longer number is the multiplicand
	if len(lier) > len(cand) {
		cand, lier = lier, cand
	}

	for i := len(lier) - 1; i >= 0; i-- {
		// if the multiplier's current digit is 0, don't generate a partial
		if lier[i] == byte('0') {
			continue
		}

		partial := make([]byte, 0, len(cand)) // the partial product for this digit
		// align the partial with the corresponding multiplier digit
		for ralign := (len(lier) - 1) - i; ralign > 0; ralign-- {
			partial = append(partial, byte('0'))
		}

		// multiply the multiplier's current digit with each digit of the multiplicand
		for j := len(cand) - 1; j >= 0; j-- {
			prod, carry = longMultDigit(lier[i], cand[j], carry)
			partial = append([]byte{prod}, partial...)
		}
		partials = append(partials, partial) // save the partial product for summing
	}

	// sum all partial products
	for _, partial := range partials {
		res = LongAdd(res, string(partial))
	}

	return res
}

// longMultDigit performs long multiplication between 2 primary digits and adds a carry
func longMultDigit(d1, d2, d3 byte) (byte, byte) {
	// perform addition on the arguments and
	n1, _ := strconv.Atoi(string(d1))
	n2, _ := strconv.Atoi(string(d2))
	n3, _ := strconv.Atoi(string(d3)) // the carry
	prod := strconv.Itoa(n1*n2 + n3)

	// determine the value of the carry
	carry := byte('0')
	if len(prod) > 1 {
		carry = prod[0]
		prod = prod[1:]
	}

	return prod[0], carry
}

// paddedBytes returns the stringified numbers as byte arrays,
//   with the shorter one front-padded to match lengths
func paddedBytes(n1, n2 string) ([]byte, []byte) {
	b1, b2 := []byte(n1), []byte(n2)
	size := len(b1) - len(b2)
	if size < 0 {
		// since n1 is shorter, add padding
		pad := make([]byte, 0, 0-size)
		for i := 0; i < 0-size; i++ {
			pad = append(pad, byte('0'))
		}
		b1 = append(pad, b1...)
	} else {
		// since n2 is shorter, add padding
		pad := make([]byte, 0, size)
		for i := 0; i < size; i++ {
			pad = append(pad, byte('0'))
		}
		b2 = append(pad, b2...)
	}
	return b1, b2
}
