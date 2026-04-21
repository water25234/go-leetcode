package twosum

func TwoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, value := range nums {

		balance := target - value

		if j, ok := m[balance]; ok {
			return []int{i, j}
		}

		m[value] = i

	}

	return nil
}

func twoSumTowPointer(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	var left int
	var right int

	left = 0
	right = len(nums) - 1

	for left < right {
		if nums[left]+nums[right] == target {
			return []int{left, right}
		}

		right--
		if left >= right {
			left++
			right = len(nums) - 1
		}
	}

	return nil
}

// head 是固定的錨點，current 是移動的工作指標
// 想像你在串珠子：

// head = 你左手抓著繩子的起點，永遠不動
// current = 你右手正在串下一顆珠子的位置，一直往後移動

// 初始狀態：
// head → [0]
// current → [0]  （兩者指向同一個節點）

// 第一輪迴圈後：
// head → [0] → [7]
// current ────────→ [7]  （current 移到新節點了）

// 第二輪迴圈後：
// head → [0] → [7] → [0]
// current ──────────────→ [0]

// 最後 return head.Next → [7] → [0] → ...（跳過那個假的 [0] 開頭）
