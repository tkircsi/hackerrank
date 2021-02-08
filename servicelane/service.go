package servicelane

import "fmt"

var width []int32 = []int32{2, 3, 1, 2, 3, 2, 3, 3}

func Test() {
	cases := [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}
	r := serviceLane(int32(len(width)), cases)
	fmt.Println(r)
}

// Complete the serviceLane function below.
func serviceLane(n int32, cases [][]int32) []int32 {
	res := []int32{}
	for _, segm := range cases {
		res = append(res, min(width[segm[0]:segm[1]+1]))
	}
	return res
}

func min(a []int32) int32 {
	var min int32 = 4
	for _, n := range a {
		if n < min {
			min = n
		}
	}
	return min
}
