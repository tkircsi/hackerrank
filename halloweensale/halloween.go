package halloweensale

import "fmt"

func Test() {
	// var p, d, m, s int32 = 20, 3, 6, 80
	// var p, d, m, s int32 = 81, 7, 11, 3339
	var p, d, m, s int32 = 20, 3, 6, 85
	r := howManyGames(p, d, m, s)
	fmt.Println(r)
}

// Complete the howManyGames function below.
func howManyGames(p int32, d int32, m int32, s int32) int32 {
	var i int32 = 0
	for {
		s = s - p
		if s < 0 {
			break
		}
		i++
		if p-d > m {
			p = p - d
		} else {
			p = m
		}
	}
	return i
}
