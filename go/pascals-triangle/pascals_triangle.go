package pascal

// Triangle builds pascal's triangle up to level n.
func Triangle(n int) [][]int {
	var triangle [][]int

	for rn := 0; rn < n; rn++ {
		row := make([]int, rn+1)

		for cn := 0; cn < rn+1; cn++ {
			if rn == 0 || cn == 0 || cn == rn {
				row[cn] = 1
				continue
			}
			row[cn] = triangle[rn-1][cn-1] + triangle[rn-1][cn]
		}
		triangle = append(triangle, row)
	}

	return triangle
}
