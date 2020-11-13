package leetcode

import (
	"fmt"
	"testing"
)

func Test_generate(t *testing.T) {
	ret := generate(5)
	for i := 0; i < len(ret);i++ {
		fmt.Println(ret[i])
	}
}