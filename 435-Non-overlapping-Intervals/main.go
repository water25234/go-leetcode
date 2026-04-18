package nonoverlappingintervals

import "sort"

// Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

// Note that intervals which only touch at a point are non-overlapping. For example, [1, 2] and [2, 3] are non-overlapping.

// Example 1:

// Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
// Output: 1
// Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
// Example 2:

// Input: intervals = [[1,2],[1,2],[1,2]]
// Output: 2
// Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
// Example 3:

// Input: intervals = [[1,2],[2,3]]
// Output: 0
// Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

// Constraints:

// 1 <= intervals.length <= 105
// intervals[i].length == 2
// -5 * 104 <= starti < endi <= 5 * 104

// Junior Array Dynamic Programming Greedy Sorting
// 複習
func eraseOverlapIntervals(intervals [][]int) int {
	// 1. 按結束時間排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prevEnd {
			// 重疊 → 移除（不更新 prevEnd，因為留的是結束更早的那個）
			count++
		} else {
			// 不重疊 → 保留，更新 prevEnd
			prevEnd = intervals[i][1]
		}
	}

	return count
}

// 先按結束時間排序：

// 排序前：[[1,100],[11,22],[1,11],[2,12]]
// 排序後：[[1,11],[2,12],[11,22],[1,100]]
// 畫在時間軸上：

// [1,--------11]
//   [2,----------12]
//               [11,--------22]
// [1,--------------------------------------100]

// 時間：1    11   22                         100
// 起始：選 [1,11]，prevEnd = 11

// i=1 看 [2,12]：起點 2 < prevEnd 11 → 重疊！移除。count = 1

// i=2 看 [11,22]：起點 11 >= prevEnd 11 → 不重疊，保留。prevEnd = 22

// i=3 看 [1,100]：起點 1 < prevEnd 22 → 重疊！移除。count = 2

// 結果
// 保留：[1,11]  [11,22]   ← 兩個，剛好首尾接上
// 移除：[2,12]  [1,100]   ← 兩個

// 答案 = 2
// 為什麼這樣最好？
// [1,--------11][11,--------22]   ← 完美接上，不重疊

//   [2,----------12]              ← 跟 [1,11] 重疊，留它沒好處
// [1,--------------------------------------100] ← 這個太長，跟所有人重疊，一定要移除
// [1,100] 佔了整個時間軸，留它的話其他三個都要移除，移除 3 個更虧。所以 greedy 的精神就是：結束越早的越優先留，因為它佔的空間最小，留給後面的機會最多。
