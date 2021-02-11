package flatland

import (
	"fmt"
	"sort"
)

func Test() {
	//c := []int32{0, 3, 1, 7, 13, 22, 10}
	//c := []int32{0, 4}

	// n := 20
	// c := []int32{13, 1, 11, 10, 6}

	// n := 95
	// c := []int32{68, 81, 46, 54, 30, 11, 19, 23, 22, 12, 38, 91, 48, 75, 26, 86, 29, 83, 62}

	// n := 90
	// c := []int32{4, 76, 16, 71, 56, 7, 77, 31, 2, 66, 12, 32, 57, 11, 19, 14, 42}

	n := 100
	c := []int32{0}

	r := flatlandSpaceStations(int32(n), c)
	fmt.Println(r)
}

// Complete the flatlandSpaceStations function below.
func flatlandSpaceStations(n int32, c []int32) int32 {
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	var max int32 = 0
	// check first
	max = c[0]
	// check last
	if max < n-c[len(c)-1] {
		max = n - c[len(c)-1] - 1
	}

	for i := 0; i < len(c)-1; i++ {
		diff := Abs(c[i]-c[i+1]) / 2
		if diff > max {
			max = diff
		}
	}
	return max
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
