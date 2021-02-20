package bomberman

import (
	"fmt"
	"strings"
)

const sec = 4

func Test() {

	// var n int32 = 3
	// grid := []string{
	// 	".......",
	// 	"...O...",
	// 	"....O..",
	// 	".......",
	// 	"OO.....",
	// 	"OO.....",
	// }

	var n int32 = 5
	grid := []string{
		".......",
		"...O.O.",
		"....O..",
		"..O....",
		"OO...OO",
		"OO.O...",
	}

	// var n int32 = 181054341
	// grid := []string{
	// 	"O..OO........O..O........OO.O.OO.OO...O.....OOO...OO.O..OOOOO...O.O..O..O.O..OOO..O..O..O....O...O....O...O..O..O....O.O.O.O.....O.....OOOO..O......O.O.....OOO....OO....OO....O.O...O..OO....OO..O...O",
	// }

	r := bomberMan(n, grid)
	for _, v := range r {
		fmt.Println(v)
	}
}

// Complete the bomberMan function below.
func bomberMan(n int32, grid []string) []string {
	m := strToIntArr(grid)

	// TODO: pattern is repeat every 4 iteration. P[1] == P[5], P[2] == P[6]
	// 4 is 'sec'?
	if n > 4 {
		n = n%4 + 4
	}
	// fmt.Printf("%d. sec\n", 0)
	// printM(m)
	for i := 1; i < int(n); i++ {
		//fmt.Printf("\r%d", i)
		m = incrM(m)
		m = detonate(m)
		// fmt.Printf("%d. sec\n", i)
		// printM(m)
	}
	return intTostrArr(m)
}

// Incerment every value in the array by 1
func incrM(m [][]int) [][]int {
	for r, row := range m {
		for c := range row {
			m[r][c] += 1
		}
	}
	return m
}

// Find a value with 'sec' and set it to 0.
// Check the value of left, right, up and down cells and
// set it's value to 0 if it's not 'sec'
func detonate(m [][]int) [][]int {
	rl := len(m)
	cl := len(m[0])
	nc := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for r, row := range m {
		for c := range row {
			if m[r][c] != sec {
				continue
			}
			m[r][c] = 0
			for _, v := range nc {
				rr := r + v[0]
				cc := c + v[1]
				if rr >= 0 && rr < rl && cc >= 0 && cc < cl && m[rr][cc] != sec {
					m[rr][cc] = 0
				}
			}

		}
	}
	return m
}

// Convert a string array to a 2D int array
func strToIntArr(grid []string) [][]int {
	res := make([][]int, len(grid))
	for row, s := range grid {
		res[row] = make([]int, len(s))
		for col, c := range s {
			res[row][col] = 0
			if c == 'O' {
				res[row][col] = 2
			}
		}
	}
	return res
}

// Convert a 2D int array to a string array
func intTostrArr(m [][]int) []string {
	res := make([]string, 0)
	for _, row := range m {
		var s strings.Builder
		for _, v := range row {
			if v == 0 {
				s.WriteRune('.')
			} else {
				s.WriteRune('O')
			}
		}
		res = append(res, s.String())
	}
	return res
}

// Prints a 2D int array in pretty format
// func printM(m [][]int) {
// 	for _, r := range m {
// 		for _, e := range r {
// 			fmt.Printf("%d", e)
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }
