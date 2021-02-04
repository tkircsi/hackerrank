package beautiful3s

import "fmt"

func Test() {
	var d int32 = 3
	// arr := []int32{
	// 	1, 2, 4, 5, 7, 8, 10,
	// }
	// arr := []int32{
	// 	1, 6, 7, 7, 8, 10, 12, 13, 14, 19,
	// }
	r := beautifulTriplets(d, TestData)
	fmt.Println(r)
}

// Complete the beautifulTriplets function below.
func beautifulTriplets(d int32, arr []int32) int32 {
	var count int32

	for j := 0; j < len(arr); j++ {
		// left until arr[j] - d >= arr[i]
		ri := 0
		for i := j - 1; i >= 0; i-- {
			if arr[i] < arr[j]-d {
				break
			}
			if arr[j]-d == arr[i] {
				ri++
			}
		}

		// right until arr[j] + d <= arr[i]
		rk := 0
		for k := j + 1; k < len(arr); k++ {
			if arr[k] > arr[j]+d {
				break
			}
			if arr[j]+d == arr[k] {
				rk++
			}
		}

		if ri > 0 && rk > 0 {
			count = count + int32(ri*rk)
		}
	}
	return count
}
