package leetcode

//剑指 Offer 12. 矩阵中的路径
//请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
//
//[["a","b","c","e"],
//["s","f","c","s"],
//["a","d","e","e"]]
//
//但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
//
//
//
//示例 1：
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
//示例 2：
//
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
//提示：
//
//1 <= board.length <= 200
//1 <= board[i].length <= 200

var m map[int]bool

func search(board [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}

	var res bool
	var x, y int
	// left
	x, y = i, j-1
	if y >= 0 && !m[mapKey(x, y)] && board[x][y] == word[0] {
		m[mapKey(x, y)] = true
		res = res || search(board, x, y, word[1:])
		m[mapKey(x, y)] = false
	}

	// right
	x, y = i, j+1
	if y < len(board[i]) && !m[mapKey(x, y)] && board[x][y] == word[0] {
		m[mapKey(x, y)] = true
		res = res || search(board, x, y, word[1:])
		m[mapKey(x, y)] = false
	}

	// up
	x, y = i-1, j
	if x >= 0 && !m[mapKey(x, y)] && board[x][y] == word[0] {
		m[mapKey(x, y)] = true
		res = res || search(board, x, y, word[1:])
		m[mapKey(x, y)] = false
	}

	// below
	x, y = i+1, j
	if x < len(board) && !m[mapKey(x, y)] && board[x][y] == word[0] {
		m[mapKey(x, y)] = true
		res = res || search(board, x, y, word[1:])
		m[mapKey(x, y)] = false
	}
	return res
}

func mapKey(i, j int) int {
	return i*1000 + j
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	// 首先查找第一个字符在矩阵中的位置，可能有多个
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				m = make(map[int]bool)
				m[mapKey(i, j)] = true
				// 找到第一个之后，进行上下左右扫描接下来的字符
				if search(board, i, j, word[1:]) {
					return true
				}
			}
		}
	}

	return false
}
