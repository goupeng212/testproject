package main

//https://segmentfault.com/a/1190000003817711
func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return nil
	}
	var result []int
	row := len(matrix)
	col := len(matrix[0])
	min := 0
	if row < col {
		min = row
	} else {
		min = col
	}
	circle := (min + 1) / 2
	for i := 0; i < circle; i++ {
		lastrow := row - 1 - i
		lastcol := col - 1 - i
		if lastrow == i {
			//第一行等于最后一行，遍历列
			for j := i; j <= lastcol; j++ {
				result = append(result, matrix[i][j])
			}
		} else if lastcol == i {
			//第一列等于最后一列 遍历行
			for j := i; j <= lastrow; j++ {
				result = append(result, matrix[j][i])
			}
		} else {
			//遍历第一行 最后一个值给到下一个遍历
			for j := i; j < lastcol; j++ {
				result = append(result, matrix[i][j])
			}
			//遍历最后一列 最后一个值给到下一个遍历
			for j := i; j < lastrow; j++ {
				result = append(result, matrix[j][lastcol])
			}
			//遍历最后一行，倒序 最后一个值给到下一个遍历
			for j := lastcol; j > i; j-- {
				result = append(result, matrix[lastrow][j])
			}
			//遍历第一列 倒叙 最后一个值给到下一个遍历，相当于开始已遍历
			for j := lastrow; j > i; j-- {
				result = append(result, matrix[j][i])
			}

		}

	}

	return result

}
