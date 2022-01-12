package checkifarraypairsaredivisiblebyk

func canArrange(arr []int, k int) bool {
	freq := make([]int, k)

	for _, i := range arr {

		rmd := i % k
		// convert negative remainders to positive
		if rmd < 0 {
			rmd += k
		}
		freq[rmd]++

	}
	for j := 1; j < k; j++ {
		// we can match all pairings with remaineder==j to remainder==k-j
		// eg. k=5 [1, 2,3,4, 12, 13, 22, 23]
		//     freq array looks like [0, 1, 3, 3, 1] => true

		if freq[j] != freq[k-j] {
			return false
		}

	}
	// for remainder==0, we must have an even number of pairings
	// eg. k=5 [5,10,15,20]
	//     freq array looks like [4, 0, 0, 0, 0] => true
	return freq[0]%2 == 0
}
