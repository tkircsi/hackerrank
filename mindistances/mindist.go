package mindistances

import "fmt"

func Test() {
	//arr := []int32{7, 1, 3, 4, 1, 7}
	//arr := []int32{3, 2, 1, 2, 3}
	arr := []int32{1, 2, 3, 4, 5, 7}
	r := minimumDistances(arr)
	fmt.Println(r)
}

// Complete the minimumDistances function below.
func minimumDistances(a []int32) int32 {
	m := make(map[int32]int)
	min := len(a)

	for i := 0; i < len(a); i++ {
		if v, ok := m[a[i]]; ok {
			// already in map
			if i-v < min {
				min = i - v
			}
			m[a[i]] = i
		} else {
			// not in map
			m[a[i]] = i
		}
	}

	if min == len(a) {
		min = -1
	}
	return int32(min)
}
