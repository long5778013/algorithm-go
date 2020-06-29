package recursion

import (
	"fmt"
	"testing"
)

func Test_PascalTriangle(t *testing.T) {
	arr := generate(6)
	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v \t \n", i, j, v2)
		}
	}

	t.Log("PascalTriangle success")
}
func Test_PascalTriangleIIGetRow(t *testing.T) {
	for i, v := range getRow(4) {
		println(i, v)
	}

	t.Log("PascalTriangleIIGetRow success")
}
