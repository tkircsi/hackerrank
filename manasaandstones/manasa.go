package manasaandstones

import "fmt"

func Test() {

	tc := [][]int32{
		{58, 69, 24},
		{83, 86, 81},
		{73, 25, 25},
		{12, 73, 82},
		{5, 3, 23},
	}

	for _, t := range tc {
		r := stones(t[0], t[1], t[2])
		fmt.Println(r)
		fmt.Println()
	}
}

// Complete the stones function below.
func stones(n int32, a int32, b int32) []int32 {
	if a == b {
		return []int32{(n - 1) * a}
	}

	if a > b {
		a, b = b, a
	}

	var i int32
	var arr []int32
	for i = 0; i < n; i++ {
		arr = append(arr, (n-i-1)*a+i*b)
	}

	return arr
}
