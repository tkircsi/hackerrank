package lisasworkbook

import "fmt"

func Test() {
	var n, k int32 = 5, 3
	var arr []int32 = []int32{4, 2, 6, 1, 10}
	r := workbook(n, k, arr)
	fmt.Println(r)
}

// Complete the workbook function below.
func workbook(n int32, k int32, arr []int32) int32 {
	var page, special int32 = 0, 0
	for i, c := range arr {
		var j int32
		for j = 1; j <= c; j = j + k {
			page++
			var s, e int32 = j, j + k - 1
			if j+k-1 > c {
				e = c
			}
			if page >= s && page <= e {
				special++
			}
			fmt.Printf("Page: %d, Chapter: %d, Chapter-page: %d, page start: %d, page end: %d\n", page, i+1, c, s, e)
		}
	}
	return special
}
