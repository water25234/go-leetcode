package checkifarraypairsaredivisiblebyk

func canArrange(arr []int, k int) bool {
	length := len(arr)
	visit := make([]bool, len(arr))

	for i := 0; i < length-1; i++ {
		if true == visit[i] {
			continue
		}

		for j := i + 1; j < length; j++ {
			if true == visit[j] {
				continue
			}

			if (arr[i]+arr[j])%k == 0 {
				visit[i] = true
				visit[j] = true
				break
			}
		}

		if false == visit[i] {
			return false
		}
	}

	return true
}
