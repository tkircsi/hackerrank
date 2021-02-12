package cavitymap

import "fmt"

func Test() {
	grid := []string{"1112", "1912", "1892", "1234"}
	r := cavityMap(grid)
	fmt.Printf("%v\n", r)
}

// Complete the cavityMap function below.
func cavityMap(grid []string) []string {

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			// check up
			if grid[i-1][j] >= grid[i][j] {
				continue
			}
			// check down
			if grid[i+1][j] >= grid[i][j] {
				continue
			}
			// chek right
			if grid[i][j+1] >= grid[i][j] {
				continue
			}
			// check left
			if grid[i][j-1] >= grid[i][j] {
				continue
			}
			grid[i] = grid[i][:j] + "X" + grid[i][j+1:]
		}
	}

	return grid
}
