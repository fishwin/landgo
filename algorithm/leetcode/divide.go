package leetcode

import "math"

func checkOverflow(v int) int {
	if v > math.MaxInt32 || v < math.MinInt32 {
		return math.MaxInt32
	}
	return v
}

func divide(a int, b int) int {
	if a == 0 {
		return 0
	}
	if b == 1 {
		return checkOverflow(a)
	}
	if b == -1 {
		return checkOverflow(-a)
	}

	same := true
	if a > 0 && b < 0 {
		same = false
		b = -b
	} else if a < 0 && b > 0 {
		same = false
		a = -a
	} else if a < 0 && b < 0 {
		same = true
		a, b = -a, -b
	}

	ret := 0
	sum := 0
	for {
		sum += b
		if sum == a {
			ret++
			break
		} else if sum > a {
			break
		} else {
			ret++
		}
	}

	if !same {
		ret = -ret
	}

	return checkOverflow(ret)
}
