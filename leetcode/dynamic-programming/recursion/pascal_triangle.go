package recursion

// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
// 在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 5
//输出:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//]

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return make([][]int, 0)
	}

	// slice：动态数组
	arr := make([][]int, numRows)
	// 创建二维Slice
	for i := range arr {
		subArray := make([]int, i+1)
		for j := range subArray {
			subArray[j] = j + 1
		}
		arr[i] = subArray
	}

	arr[0][0] = 1
	if numRows <= 1 {
		return arr
	}
	arr[1][0] = 1
	arr[1][1] = 1

	if numRows <= 2 {
		return arr
	}
	// 从三行开始
	for row := 2; row < numRows; row++ {
		for col := 0; col <= row; col++ {
			if col == 0 || col == row {
				arr[row][col] = 1
			} else {
				arr[row][col] = arr[row-1][col-1] + arr[row-1][col]
			}
		}
	}

	return arr
}
