package leetcode

import "math"

//1512. 好数对的数目
//给你一个整数数组 nums 。
//
//如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。
//
//返回好数对的数目。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3,1,1,3]
//输出：4
//解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始
//示例 2：
//
//输入：nums = [1,1,1,1]
//输出：6
//解释：数组中的每组数字都是好数对
//示例 3：
//
//输入：nums = [1,2,3]
//输出：0
//
//
//提示：
//
//1 <= nums.length <= 100
//1 <= nums[i] <= 100

func numIdenticalPairs(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

//1389. 按既定顺序创建目标数组
//给你两个整数数组 nums 和 index。你需要按照以下规则创建目标数组：
//
//目标数组 target 最初为空。
//按从左到右的顺序依次读取 nums[i] 和 index[i]，在 target 数组中的下标 index[i] 处插入值 nums[i] 。
//重复上一步，直到在 nums 和 index 中都没有要读取的元素。
//请你返回目标数组。
//
//题目保证数字插入位置总是存在。
//
//
//
//示例 1：
//
//输入：nums = [0,1,2,3,4], index = [0,1,2,2,1]
//输出：[0,4,1,3,2]
//解释：
//nums       index     target
//0            0        [0]
//1            1        [0,1]
//2            2        [0,1,2]
//3            2        [0,1,3,2]
//4            1        [0,4,1,3,2]
//示例 2：
//
//输入：nums = [1,2,3,4,0], index = [0,1,2,3,0]
//输出：[0,1,2,3,4]
//解释：
//nums       index     target
//1            0        [1]
//2            1        [1,2]
//3            2        [1,2,3]
//4            3        [1,2,3,4]
//0            0        [0,1,2,3,4]
//示例 3：
//
//输入：nums = [1], index = [0]
//输出：[1]
//
//
//提示：
//
//1 <= nums.length, index.length <= 100
//nums.length == index.length
//0 <= nums[i] <= 100
//0 <= index[i] <= i

//[0,1,2,3,4]
//[0,1,2,2,1]
func createTargetArray(nums []int, index []int) []int {
	ret := make([]int, len(nums))

	for i := 0; i < len(ret); i++ {
		ret[i] = -1
	}

	for i := 0; i < len(ret); i++ {
		if ret[index[i]] == -1 {
			ret[index[i]] = nums[i]
		} else {
			for j := len(ret) - 1; j > index[i]; j-- {
				if ret[j] == -1 && ret[j-1] != -1 {
					for m := j; m > index[j]; m-- {
						ret[m], ret[m-1] = ret[m-1], ret[m]
					}
					break
				}
			}
			ret[index[i]] = nums[i]
		}
	}
	return ret
}

func diagonalSum(mat [][]int) int {
	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
		sum += mat[i][len(mat)-1-i]
	}
	if len(mat)%2 != 0 {
		sum -= mat[len(mat)/2][len(mat)/2]
	}
	return sum
}

//1450. 在既定时间做作业的学生人数
//给你两个整数数组 startTime（开始时间）和 endTime（结束时间），并指定一个整数 queryTime 作为查询时间。
//
//已知，第 i 名学生在 startTime[i] 时开始写作业并于 endTime[i] 时完成作业。
//
//请返回在查询时间 queryTime 时正在做作业的学生人数。形式上，返回能够使 queryTime 处于区间 [startTime[i], endTime[i]]（含）的学生人数。
//
//
//
//示例 1：
//
//输入：startTime = [1,2,3], endTime = [3,2,7], queryTime = 4
//输出：1
//解释：一共有 3 名学生。
//第一名学生在时间 1 开始写作业，并于时间 3 完成作业，在时间 4 没有处于做作业的状态。
//第二名学生在时间 2 开始写作业，并于时间 2 完成作业，在时间 4 没有处于做作业的状态。
//第三名学生在时间 3 开始写作业，预计于时间 7 完成作业，这是是唯一一名在时间 4 时正在做作业的学生。
//示例 2：
//
//输入：startTime = [4], endTime = [4], queryTime = 4
//输出：1
//解释：在查询时间只有一名学生在做作业。
//示例 3：
//
//输入：startTime = [4], endTime = [4], queryTime = 5
//输出：0
//示例 4：
//
//输入：startTime = [1,1,1,1], endTime = [1,3,2,4], queryTime = 7
//输出：0
//示例 5：
//
//输入：startTime = [9,8,7,6,5,4,3,2,1], endTime = [10,10,10,10,10,10,10,10,10], queryTime = 5
//输出：5
//
//
//提示：
//
//startTime.length == endTime.length
//1 <= startTime.length <= 100
//1 <= startTime[i] <= endTime[i] <= 1000
//1 <= queryTime <= 1000

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count++
		}
	}
	return count
}

//LCP 17. 速算机器人
//小扣在秋日市集发现了一款速算机器人。店家对机器人说出两个数字（记作 x 和 y），请小扣说出计算指令：
//
//"A" 运算：使 x = 2 * x + y；
//"B" 运算：使 y = 2 * y + x。
//在本次游戏中，店家说出的数字为 x = 1 和 y = 0，小扣说出的计算指令记作仅由大写字母 A、B 组成的字符串 s，字符串中字符的顺序表示计算顺序，请返回最终 x 与 y 的和为多少。
//
//示例 1：
//
//输入：s = "AB"
//
//输出：4
//
//解释：
//经过一次 A 运算后，x = 2, y = 0。
//再经过一次 B 运算，x = 2, y = 2。
//最终 x 与 y 之和为 4。
//
//提示：
//
//0 <= s.length <= 10
//s 由 'A' 和 'B' 组成

func calculate(s string) int {
	x, y := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			x = 2*x + y
		} else if s[i] == 'B' {
			y = 2*y + x
		}
	}
	return x + y
}

//1266. 访问所有点的最小时间
//平面上有 n 个点，点的位置用整数坐标表示 points[i] = [xi, yi]。请你计算访问所有这些点需要的最小时间（以秒为单位）。
//
//你可以按照下面的规则在平面上移动：
//
//每一秒沿水平或者竖直方向移动一个单位长度，或者跨过对角线（可以看作在一秒内向水平和竖直方向各移动一个单位长度）。
//必须按照数组中出现的顺序来访问这些点。
//
//
//示例 1：
//
//
//
//输入：points = [[1,1],[3,4],[-1,0]]
//输出：7
//解释：一条最佳的访问路径是： [1,1] -> [2,2] -> [3,3] -> [3,4] -> [2,3] -> [1,2] -> [0,1] -> [-1,0]
//从 [1,1] 到 [3,4] 需要 3 秒
//从 [3,4] 到 [-1,0] 需要 4 秒
//一共需要 7 秒
//示例 2：
//
//输入：points = [[3,2],[-2,2]]
//输出：5
//
//
//提示：
//
//points.length == n
//1 <= n <= 100
//points[i].length == 2
//-1000 <= points[i][0], points[i][1] <= 1000

func minTimeToVisitAllPoints(points [][]int) int {
	sum := 0
	for i := 0; i < len(points)-1; i++ {
		sum += int(math.Max(math.Abs(float64(points[i][0]-points[i+1][0])), math.Abs(float64(points[i][1]-points[i+1][1]))))
	}
	return sum
}

//1534. 统计好三元组
//给你一个整数数组 arr ，以及 a、b 、c 三个整数。请你统计其中好三元组的数量。
//
//如果三元组 (arr[i], arr[j], arr[k]) 满足下列全部条件，则认为它是一个 好三元组 。
//
//0 <= i < j < k < arr.length
//|arr[i] - arr[j]| <= a
//|arr[j] - arr[k]| <= b
//|arr[i] - arr[k]| <= c
//其中 |x| 表示 x 的绝对值。
//
//返回 好三元组的数量 。
//
//
//
//示例 1：
//
//输入：arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
//输出：4
//解释：一共有 4 个好三元组：[(3,0,1), (3,0,1), (3,1,1), (0,1,1)] 。
//示例 2：
//
//输入：arr = [1,1,2,2,3], a = 0, b = 0, c = 1
//输出：0
//解释：不存在满足所有条件的三元组。
//
//
//提示：
//
//3 <= arr.length <= 100
//0 <= arr[i] <= 1000
//0 <= a, b, c <= 1000

func countGoodTriplets(arr []int, a int, b int, c int) int {
	i, j, k, sum := 0, 0, 0, 0
	for i = 0; i < len(arr); i++ {
		for j = i + 1; j < len(arr); j++ {
			for k = j + 1; k < len(arr); k++ {
				if int(math.Abs(float64(arr[i]-arr[j]))) <= a && int(math.Abs(float64(arr[j]-arr[k]))) <= b && int(math.Abs(float64(arr[i]-arr[k]))) <= c {
					sum++
				}
			}
		}
	}
	return sum
}

func decrypt(code []int, k int) []int {
	ret := make([]int, len(code))
	t, p := 0, 0
	for i := 0; i < len(code); i++ {
		if k == 0 {
			ret[i] = 0
		} else if k > 0 {
			t = 1
			for t <= k {
				ret[i] += code[(i+t) % len(code)]
				t++
			}
		} else {
			t = -1
			for t >= k {
				p = i+t
				if p < 0 {
					p = len(code)+p
				}
				ret[i] += code[p]
				t--
			}
		}
	}
	return ret
}