package leetcode

//剑指 Offer 29. 顺时针打印矩阵
//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
//
//
//示例 1：
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//示例 2：
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//限制：
//
//0 <= matrix.length <= 100
//0 <= matrix[i].length <= 100

const (
	right = 0
	down  = 1
	left  = 2
	up    = 3
)

func switchDirection(cur *int) {
	*cur = (*cur + 1) % 4
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return nil
	}

	var ret []int
	direction := right
	i, j := 0, 0
	u, d, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1 // 上下左右边界

	for {
		ret = append(ret, matrix[i][j])

		// 判断下一个方格是否是有效的，如果有效则继续沿当前方向走，否则变向
		if direction == right {
			j++
			if j > r { // 到达右边界,则转向下走
				switchDirection(&direction)
				u++ // 上边界收缩
				j--
				i++
			}
		} else if direction == down {
			i++
			if i > d { // 如果越界
				switchDirection(&direction)
				r-- // 右边界收缩
				i--
				j--
			}
		} else if direction == left {
			j--
			if j < l { // 如果越界
				switchDirection(&direction)
				d-- // 下边界收缩
				j++
				i--
			}
		} else {
			i--
			if i < u { // 如果越界
				switchDirection(&direction)
				l++ // 左边界收缩
				i++
				j++
			}
		}

		if l > r || u > d { // 上下边界相遇或左右边界相遇
			break
		}
	}

	return ret
}
