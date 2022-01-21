package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1
	}

	currnetGas, start := 0, 0

	for k := 0; k < len(gas); k++ {
		currnetGas = currnetGas - cost[k] + gas[k]
		if currnetGas < 0 {
			currnetGas = 0
			start = k + 1
		}
	}

	return start
}
