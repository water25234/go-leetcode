package main

func solution(R []string) int {
	n := len(R)
	m := len(R[0])

	// 0:right, 1:down, 2:left, 3:up
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	// seenState[r][c][dir] = 這個狀態是否出現過
	seenState := make([][][]bool, n)
	for i := 0; i < n; i++ {
		seenState[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			seenState[i][j] = make([]bool, 4)
		}
	}

	// cleaned[r][c] = 這格有沒有清過
	cleaned := make([][]bool, n)
	for i := 0; i < n; i++ {
		cleaned[i] = make([]bool, m)
	}

	r, c, dir := 0, 0, 0
	cleaned[r][c] = true
	ans := 1

	for !seenState[r][c][dir] {
		seenState[r][c][dir] = true

		nr := r + dr[dir]
		nc := c + dc[dir]

		// 前方不能走: 轉向
		if nr < 0 || nr >= n || nc < 0 || nc >= m || R[nr][nc] == 'X' || R[nr][nc] == 'x' {
			dir = (dir + 1) % 4
			continue
		}

		// 前進
		r, c = nr, nc
		if !cleaned[r][c] {
			cleaned[r][c] = true
			ans++
		}
	}

	return ans
}
