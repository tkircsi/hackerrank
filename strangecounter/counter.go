package strangecounter

import (
	"fmt"
)

func Test() {
	tcs := []int64{14, 2, 7, 4, 98744}

	for _, t := range tcs {
		r := strangeCounter(t)
		fmt.Printf("%d : %d\n", t, r)
	}
}

// Complete the strangeCounter function below.
func strangeCounter(t int64) int64 {
	i := int64(1)
	j := int64(1)
	for {
		// step counter counts the first number of the given part
		i = i + j*3
		// step size
		j = j * 2
		if int64(i) > t {
			break
		}

	}
	return (i - j/2*3) + j/2*3 - t
}
