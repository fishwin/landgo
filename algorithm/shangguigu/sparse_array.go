package shangguigu

import "fmt"

// 稀疏数组

func SparseArrayCompress(src [][]int) [][3]int {
	if len(src) == 0 || len(src[0]) == 0 {
		return [][3]int{{0, 0, 0}}
	}

	ret := make([][3]int, 0)
	ret = append(ret, [3]int{len(src), len(src[0]), 0})

	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src[0]); j++ {
			if src[i][j] != 0 {
				ret = append(ret, [3]int{i, j, src[i][j]})
				ret[0][2]++
			}
		}
	}

	fmt.Println("test111     ")

	return ret
}

func SparseArrayDecompress(src [][3]int) [][]int {
	if len(src) == 0 {
		return nil
	}

	ret := make([][]int, src[0][0])
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, src[0][1])
	}

	for i := 1; i < len(src); i++ {
		ret[src[i][0]][src[i][1]] = src[i][2]
	}

	return ret
}
