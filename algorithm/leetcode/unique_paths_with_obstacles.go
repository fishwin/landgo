package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	if len(obstacleGrid[0]) == 0 {
		return 0
	}

	ret := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		ret[i] = make([]int, len(obstacleGrid[0]))
	}

	for i := 0; i < len(ret); i++ {
		for j := 0; j < len(ret[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				ret[i][j] = 0
			} else {
				if i-1 < 0 && j-1 >= 0 { // 当方格处于左边缘时
					ret[i][j] = 0 + ret[i][j-1]
				} else if i-1 >= 0 && j-1 < 0 { // 当方格处于上边缘时
					ret[i][j] = ret[i-1][j]
				} else if i-1 >= 0 && j-1 >= 0 { // // 当方格处于非边缘时
					ret[i][j] = ret[i-1][j] + ret[i][j-1]
				} else { // 原点[0,0]
					ret[i][j] = 1
				}
			}
		}
	}

	return ret[len(ret)-1][len(ret[0])-1]
}
