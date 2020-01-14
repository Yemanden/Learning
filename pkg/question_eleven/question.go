package questioneleven

type depthFirstSearch interface {
	dfs([][]byte, coordination)
}

type breadthFirstSearch interface {
	bfs([][]byte, coordination)
}

// IslandCounter ...
type IslandCounter interface {
	GetIslandCountDFS([][]byte) int
	GetIslandCountBFS([][]byte) int
	depthFirstSearch
	breadthFirstSearch
}

type islandCounter struct{}

type coordination struct {
	x int
	y int
}

// GetIslandCountDFS ...
func (i *islandCounter) GetIslandCountDFS(grid [][]byte) int {
	var count int
	for k := 0; k < len(grid); k++ {
		for j := 0; j < len(grid[k]); j++ {
			if grid[k][j] == 1 {
				i.dfs(grid, coordination{k, j})
				count++
			}
		}
	}
	return count
}

// GetIslandCountBFS ...
func (i *islandCounter) GetIslandCountBFS(grid [][]byte) int {
	var count int
	for k := 0; k < len(grid); k++ {
		for j := 0; j < len(grid[k]); j++ {
			if grid[k][j] == 1 {
				i.bfs(grid, coordination{k, j})
				count++
			}
		}
	}
	return count
}

func (i *islandCounter) dfs(grid [][]byte, coord coordination) {
	if coord.x < 0 || coord.x >= len(grid) || // if coordination not contains in range on horizontal
		coord.y < 0 || coord.y >= len(grid[coord.x]) || // if coordination not contains in range on vertical
		grid[coord.x][coord.y] == 0 { // if coordination not is earth
		return
	}
	if grid[coord.x][coord.y] == 1 {
		grid[coord.x][coord.y] = 0
		i.dfs(grid, coordination{coord.x + 1, coord.y})
		i.dfs(grid, coordination{coord.x, coord.y + 1})
		i.dfs(grid, coordination{coord.x - 1, coord.y})
		i.dfs(grid, coordination{coord.x, coord.y - 1})
	}
}

func (i *islandCounter) bfs(grid [][]byte, coord coordination) {
	var queue = []coordination{coord}
	for len(queue) > 0 {
		if queue[0].x < len(grid)-1 && grid[queue[0].x+1][queue[0].y] == 1 {
			queue = append(queue, coordination{queue[0].x + 1, queue[0].y})
		}
		if queue[0].y < len(grid[queue[0].x])-1 && grid[queue[0].x][queue[0].y+1] == 1 {
			queue = append(queue, coordination{queue[0].x, queue[0].y + 1})
		}
		if queue[0].x > 0 && grid[queue[0].x-1][queue[0].y] == 1 {
			queue = append(queue, coordination{queue[0].x - 1, queue[0].y})
		}
		if queue[0].y > 0 && grid[queue[0].x][queue[0].y-1] == 1 {
			queue = append(queue, coordination{queue[0].x, queue[0].y - 1})
		}
		grid[queue[0].x][queue[0].y] = 0
		queue = queue[1:]
	}
}

// NewIslandCounter ...
func NewIslandCounter() IslandCounter {
	return &islandCounter{}
}
