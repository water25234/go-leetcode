package complementofbase10integer

func bitwiseComplement(n int) int {
	p := 1
	for p < n {
		p = 2*p + 1 // pow(2, N) - 1
	}

	// fmt.Println(p, strconv.FormatInt(int64(p), 2))
	// fmt.Println(n, strconv.FormatInt(int64(n), 2))
	return (p ^ n) // p (XOR) n
}
