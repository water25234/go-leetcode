package numberofislands

// Depth-First Search
// Input: grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// Output: 3
func numIslands(grid [][]byte) int {
	count := 0

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if grid[row][column] == '1' {
				count++
				checkGrid(grid, row, column)
			}
		}
	}

	return count
}

func checkGrid(grid [][]byte, row, column int) {
	if true == inNotArea(grid, row, column) {
		return
	}

	grid[row][column] = '2'

	checkGrid(grid, row-1, column) // 上
	checkGrid(grid, row+1, column) // 下
	checkGrid(grid, row, column-1) // 左
	checkGrid(grid, row, column+1) // 右
}

func inNotArea(grid [][]byte, row, column int) bool {
	return row < 0 ||
		row >= len(grid) ||
		column < 0 ||
		column >= len(grid[0]) ||
		grid[row][column] == '0' ||
		grid[row][column] == '2'
}
