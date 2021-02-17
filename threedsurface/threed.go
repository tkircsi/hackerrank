package threedsurface

import "fmt"

func Test() {
	A := [][]int32{
		{1, 3, 4},
		{2, 2, 3},
		{1, 2, 4},
	}

	r := surfaceArea(A)
	fmt.Println(r)
}

// Complete the surfaceArea function below.
func surfaceArea(A [][]int32) int32 {
	lx := len(A)
	ly := len(A[0])
	var surf int32

	for x := 0; x < lx; x++ {
		for y := 0; y < ly; y++ {
			// default
			var val int32
			val = A[x][y]

			var left, front, right, back int32
			// left
			if y > 0 {
				left = A[x][y-1]
				if left > val {
					left = val
				}
			}
			// front
			if x > 0 {
				front = A[x-1][y]
				if front > val {
					front = val
				}
			}
			// right
			if y < ly-1 {
				right = A[x][y+1]
				if right > val {
					right = val
				}
			}
			// back
			if x < lx-1 {
				back = A[x+1][y]
				if back > val {
					back = val
				}
			}

			// full front, left, right, back surface
			val *= 4

			// top and bottom if not zero
			if val > 0 {
				val = val + 2
			}

			val = val - left - right - front - back

			fmt.Printf("x: %d, y: %d, left: %d, right: %d, front: %d, back: %d, val: %d\n", x, y, left, right, front, back, val)

			surf += val

		}
	}

	return surf
}
