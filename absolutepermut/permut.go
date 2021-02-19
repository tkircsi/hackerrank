package absolutepermut

import "fmt"

func Test() {
	// tcs := [][]int32{
	// 	// {4, 2},
	// 	// {2, 1},
	// 	// {3, 0},
	// 	// {3, 2},
	// 	{10, 0},
	// 	{10, 1},
	// 	{7, 0},
	// 	{10, 9},
	// 	{9, 0},
	// 	{10, 3},
	// 	{8, 2},
	// 	{8, 0},
	// }

	// tcs := [][]int32{
	// 	{94, 70},
	// 	{95, 49},
	// 	{92, 14},
	// 	{96, 2},
	// 	{98, 7},
	// 	{92, 85},
	// 	{90, 15},
	// 	{92, 10},
	// 	{94, 0},
	// 	{92, 40},
	// }

	tcs := [][]int32{
		{69187, 0},
		{55596, 42041},
		{49056, 0},
		{93559, 72338},
		{1394, 0},
		{68546, 34273},
		{4979, 3186},
		{89291, 0},
		{86542, 1},
		{69652, 0},
	}

	for _, tc := range tcs {
		r := absolutePermutation(tc[0], tc[1])
		fmt.Println(r)
	}
}

// Complete the absolutePermutation function below.
func absolutePermutation(n int32, k int32) []int32 {
	res := make([]int32, n)
	m := map[int32]struct{}{}
	var i int32
	// r := [][]int32{}
	for i = 1; i <= n; i++ {
		a := i + k
		b := i - k
		// r = append(r, []int32{a, b})
		//fmt.Println(r)

		if a < 0 || a > n {
			a = 0
		}
		if b < 0 || b > n {
			b = 0
		}

		switch {
		case a == b:
			if _, ok := m[a]; a == 0 || ok {
				return []int32{-1}
			}
			res[i-1] = a
		case a < b:
			if _, ok := m[a]; !ok && a > 0 {
				res[i-1] = a
			} else if _, ok := m[b]; !ok && b > 0 {
				res[i-1] = b
			} else {
				return []int32{-1}
			}
		case a > b:
			if _, ok := m[b]; !ok && b > 0 {
				res[i-1] = b
			} else if _, ok := m[a]; !ok && a > 0 {
				res[i-1] = a
			} else {
				return []int32{-1}
			}
		}
		m[res[i-1]] = struct{}{}
	}
	return res
}

// func contains(arr []int32, i int32) bool {
// 	for _, v := range arr {
// 		if v == i {
// 			//fmt.Printf("v: %d, i: %d\n", v, i)
// 			return true
// 		}
// 	}
// 	return false
// }

// 5
// 2

// 5-2 = 3
// -1-2 = -3

// p = i + k
// p = i - k
