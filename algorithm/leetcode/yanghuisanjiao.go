package leetcode

//118. 杨辉三角

//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
//
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 5
//输出:
//[
//[1],
//[1,1],
//[1,2,1],
//[1,3,3,1],
//[1,4,6,4,1]
//]

func generate(numRows int) [][]int {
	var ret [][]int
	for r := 0; r < numRows; r++ {
		temp := make([]int, r+1)
		for i := 0; i <= r; i++ {
			if r-1 < 0 {
				temp[i] = 1
			} else {
				if i-1 >= 0 {
					temp[i] += ret[r-1][i-1]
				}
				if i < len(ret[r-1]) {
					temp[i] += ret[r-1][i]
				}
			}
		}
		ret = append(ret, temp)
	}
	return ret
}
