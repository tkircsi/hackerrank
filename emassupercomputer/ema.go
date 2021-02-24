package emassupercomputer

import (
	"fmt"
)

func Test() {
	// grid := []string{
	// 	"GGGGGG",
	// 	"GBBBGB",
	// 	"GGGGGG",
	// 	"GGBBGB",
	// 	"GGGGGG",
	// }
	// result: 5

	grid := []string{
		"BGBBGB",
		"GGGGGG",
		"BGBBGB",
		"GGGGGG",
		"BGBBGB",
		"BGBBGB",
	}
	// result: 25

	// grid := []string{
	// 	"GGGGGGGGGGGG",
	// 	"GBGGBBBBBBBG",
	// 	"GBGGBBBBBBBG",
	// 	"GGGGGGGGGGGG",
	// 	"GGGGGGGGGGGG",
	// 	"GGGGGGGGGGGG",
	// 	"GGGGGGGGGGGG",
	// 	"GBGGBBBBBBBG",
	// 	"GBGGBBBBBBBG",
	// 	"GBGGBBBBBBBG",
	// 	"GGGGGGGGGGGG",
	// 	"GBGGBBBBBBBG",
	// }
	// Result: 81

	r := twoPluses(grid)
	fmt.Println(r)

}

// Complete the twoPluses function below.
func twoPluses(grid []string) int32 {
	// find G-s and build plus's set
	var pluses = []map[int]struct{}{}
	for i, r := range grid {
		for j, c := range r {
			if c == 'G' {
				m := plusvalue(grid, i, j)
				pluses = append(pluses, m...)
			}
		}
	}

	// calculate pluses product and check
	// overlapping
	var res = 0
	for i := 0; i < len(pluses); i++ {
		for j := i + 1; j < len(pluses); j++ {
			p1 := pluses[i]
			p2 := pluses[j]
			if intersection(p1, p2) {
				continue
			}
			if len(p1)*len(p2) > res {
				res = len(p1) * len(p2)
			}
		}
	}
	return int32(res)
}

// create a set(map) from a plus. The grid's cell numbered from
// 1, 2, 3, ...... n*m
func plusvalue(grid []string, i, j int) []map[int]struct{} {
	val := 1
	l := len(grid[0])
	ret := make([]map[int]struct{}, 0)
	m := map[int]struct{}{}
	m[i*l+j] = struct{}{}
	nm := copy(m)
	ret = append(ret, nm)
	for {
		if i+val >= len(grid) || grid[i+val][j] == 'B' {
			break
		}
		if i-val < 0 || grid[i-val][j] == 'B' {
			break
		}
		if j+val >= len(grid[0]) || grid[i][j+val] == 'B' {
			break
		}
		if j-val < 0 || grid[i][j-val] == 'B' {
			break
		}
		m[(i+val)*l+j] = struct{}{}
		m[(i-val)*l+j] = struct{}{}
		m[i*l+j-val] = struct{}{}
		m[i*l+j+val] = struct{}{}
		// copy map and add to array
		nm := copy(m)
		ret = append(ret, nm)
		val++
	}
	return ret
}

// intersection returns True if the 2 map has at least one
// common key
func intersection(m1, m2 map[int]struct{}) bool {
	for k := range m1 {
		if _, ok := m2[k]; ok {
			return true
		}
	}
	return false
}

func copy(m map[int]struct{}) map[int]struct{} {
	nm := map[int]struct{}{}
	for k, v := range m {
		nm[k] = v
	}
	return nm
}
