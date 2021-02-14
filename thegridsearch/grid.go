package thegridsearch

import (
	"fmt"
	"strings"
)

func Test() {
	tG := []string{ // YES
		"7283455864",
		"6731158619",
		"8988242643",
		"3830589324",
		"2229505813",
		"5633845374",
		"6473530293",
		"7053106601",
		"0834282956",
		"4607924137",
	}
	tP := []string{
		"9505",
		"3845",
		"3530",
	}

	// tG := []string{ // NO
	// 	"400453592126560",
	// 	"114213133098692",
	// 	"474386082879648",
	// 	"522356951189169",
	// 	"887109450487496",
	// 	"252802633388782",
	// 	"502771484966748",
	// 	"075975207693780",
	// 	"511799789562806",
	// 	"404007454272504",
	// 	"549043809916080",
	// 	"962410809534811",
	// 	"445893523733475",
	// 	"768705303214174",
	// 	"650629270887160",
	// }
	// tP := []string{
	// 	"99",
	// 	"99",
	// }

	// tG := []string{ // NO
	// 	"1234",
	// 	"4321",
	// 	"9999",
	// 	"9999",
	// }
	// tP := []string{
	// 	"12",
	// 	"21",
	// }

	// tG := []string{ // YES
	// 	"123412",
	// 	"561212",
	// 	"123634",
	// 	"781288",
	// }

	// tP := []string{
	// 	"12",
	// 	"34",
	// }

	r := gridSearch(tG, tP)
	fmt.Println(r)
}

// Complete the gridSearch function below.
func gridSearch(G []string, P []string) string {
	for i := 0; i < len(G)-len(P)+1; i++ {
		r := substrings(G[i], P[0])
		for _, v := range r {
			r := checkPattern(i, v, G, P)
			if r {
				return "YES"
			}
		}
	}
	return "NO"
}

func substrings(s string, su string) []int {
	var res []int
	start := 0
	for {
		ix := strings.Index(s, su)
		if ix == -1 {
			break
		}
		start += ix
		res = append(res, start)
		start++
		s = s[ix+1:]
	}
	return res
}

func checkPattern(r int, c int, G []string, P []string) bool {
	var ret bool = true
	for i := 1; i < len(P); i++ {
		end := c + len(P[i])
		if end > len(G[r+i]) {
			end = len(G[r+i])
		}
		if G[r+i][c:end] != P[i] {
			ret = false
			break
		}
	}
	return ret
}
