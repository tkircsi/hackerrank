package kaprekar

import (
	"fmt"
	"strconv"
)

func Test() {
	kaprekarNumbers(1, 100)
}

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p int32, q int32) {
	inv := true
	for i := p; i <= q; i++ {
		if kaprekar(i) {
			fmt.Print(i, " ")
			inv = false
		}
	}
	if inv {
		fmt.Print("INVALID RANGE")
	}
}

func kaprekar(nn int32) bool {
	n := int(nn)
	n2 := n * n
	s := fmt.Sprintf("%d", n2)
	le := len(s)
	r := le / 2
	l := le / 2
	if r+l < le {
		r++
	}
	//fmt.Printf("n: %d, r: %d, l: %d\n", n, r, l)
	sn2 := fmt.Sprintf("%d", n2)
	sl := sn2[:l]
	sr := sn2[l:]

	nl, err := strconv.Atoi(sl)
	if err != nil {
		nl = 0
	}
	nr, err := strconv.Atoi(sr)
	if err != nil {
		nr = 0
	}
	if nl+nr != n {
		return false
	}
	return true
}
