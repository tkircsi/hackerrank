package fairrations

import "fmt"

func Test() {
	B := []int32{2, 3, 4, 4, 5, 6, 7, 7, 6, 6, 4, 5, 4, 5}
	r := fairRations(B)
	if r == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println(r)
	}
}

// Complete the fairRations function below.
func fairRations(B []int32) int32 {
	lx := -1
	c := 0
	for i := 0; i < len(B); i++ {
		if B[i]%2 == 1 {
			if lx == -1 {
				lx = i
			} else {
				c = c + i - lx
				lx = -1
			}
		}
	}

	if lx != -1 {
		return -1
	}
	return int32(c * 2)
}

//Start:	0 1 0 1 0
//Step 1:	0 0 1 1 0
//Step 2:	0 0 0 0 0

// 0 1 0 0 1 0
// 0 0 1 0 1 0
// 0 0 0 1 1 0
// 0 0 0 0 0 0

// 0 1 0 0 1 1 0 0 1 0
// 0 0 1 0 1 1 0 0 1 0
// 0 0 0 1 1 1 0 0 1 0
// 0 0 0 0 0 1 0 0 1 0
// 0 0 0 0 0 0 1 0 1 0
// 0 0 0 0 0 0 0 1 1 0
// 0 0 0 0 0 0 0 0 0 0
