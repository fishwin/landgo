package leetcode

import "testing"

func Test_oddEvenArray(t *testing.T) {
	if r := oddEvenArray([]int{1,3,5,7,9,10}); r != 10 {
		t.Error("fail")
	} else {
		t.Log(r)
	}

	if r:=oddEvenArray([]int{1,2,4,6,8,10}); r != 1 {
		t.Error("fail")
	} else {
		t.Log(r)
	}

	if r := oddEvenArray([]int{1,3,5,6,7,9}); r != 6 {
		t.Error("fail")
	} else {
		t.Log(r)
	}
}
