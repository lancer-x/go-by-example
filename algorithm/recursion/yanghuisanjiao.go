package main

//生成杨辉三角

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		//第i层
		res[i] = append(res[i], make([]int, i + 1)...)
		for j := 0; j <= i; j++ {
			res[i][j] = calc(res, i, j)
		}
	}
	return res
}


func calc(res [][]int, x int, y int) int {
	if x == 0 || y == 0 || x == y {
		return 1
	}
	if res[x][y] != 0 {
		return res[x][y]
	}
	return calc(res, x-1, y-1) + calc(res, x-1, y)
}