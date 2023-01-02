 
var directions = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
	[]int{-1, 0},
}

func inBounds(grid [][]int, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}
	return true
}



func orangesRotting(grid [][]int) int {
    
	cnt := 0

	// pop from queue append queue[1:]
	queue := [][3]int{}
	//1. Iterate through grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			curr := grid[i][j]

			if curr == 1 {
				cnt += 1
			}
			if curr == 2 {
				thing := [3]int{i, j, 0}
				queue = append(queue, thing)
			}
		}
	}
	if cnt == 0 {
		return 0
	}
	output := 0

	for len(queue) > 0 {
		x, y, steps := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		output = steps
		for i := range directions {
			di, dj := directions[i][0], directions[i][1]
			dx := di + x
			dy := dj + y
			if inBounds(grid, dx, dy) && grid[dx][dy] == 1 {
				grid[dx][dy] = 2
				cnt -= 1
				queue = append(queue, [3]int{dx, dy, steps + 1})
			}
		}
	}
    if cnt != 0 {
        return -1
    }

	return output
}