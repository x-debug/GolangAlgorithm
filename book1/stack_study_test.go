package book1

import (
	"fmt"
	"testing"
)

func TestComputeOperateNum(t *testing.T) {
	scaned := ScanOperateNum("12+21")
	lookNodes(scaned, t)
}

func TestComputeNum(t *testing.T) {
	ll := ScanOperateNum("12+21*2+5*6+8/2-33")
	num := ComputeNum(ll)

	fmt.Println(num)
}
