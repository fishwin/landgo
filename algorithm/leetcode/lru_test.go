package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	lru := Constructor(2)
	lru.Put(2,1)
	lru.Put(2,2)
	fmt.Println(lru.Get(1))
	lru.Put(3,3)
	lru.Put(4,4)
	fmt.Println(lru.Get(2))
}