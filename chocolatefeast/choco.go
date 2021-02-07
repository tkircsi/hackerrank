package chocolatefeast

import (
	"fmt"
)

func Test() {
	var test [][]int32 = [][]int32{
		{15, 3, 2},
		{10, 2, 5},
		{12, 4, 4},
		{6, 2, 2},
		{7, 2, 2},
	}

	for _, t := range test {
		r := chocolateFeast(t[0], t[1], t[2])
		fmt.Println(r)
	}
}

// Complete the chocolateFeast function below.
func chocolateFeast(n int32, c int32, m int32) int32 {
	var cc, w, e int32
	cc = n / c
	w = 0
	e = 0

	for {
		e = e + cc
		w = cc + w
		cc = w / m
		w = w % m
		//fmt.Printf("e: %d, cc:%d, w: %d\n", e, cc, w)
		if cc == 0 {
			break
		}
	}

	return e
}
