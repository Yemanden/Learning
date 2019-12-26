package questioneleven

func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(&grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid *[][]byte, x, y int) {
	if x < 0 || x >= len(*grid) || y < 0 || y >= len((*grid)[x]) || (*grid)[x][y] == 0 {
		return
	}
	if (*grid)[x][y] == 1 {
		(*grid)[x][y] = 0
		dfs(grid, x+1, y)
		dfs(grid, x, y+1)
		dfs(grid, x-1, y)
		dfs(grid, x, y-1)
	}
}
