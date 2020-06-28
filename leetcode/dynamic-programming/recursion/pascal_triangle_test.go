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
