📖 第一部分：什麼是 Sliding Window？
生活化比喻
想像你在看電影院的 連續座位排，你要找出「連續 3 個座位中，加總票價最高的組合」。
你不會每次都從頭數，而是：

看完座位 1, 2, 3 後
下次只要 減掉座位 1，加上座位 4，就變成 2, 3, 4
繼續 減掉座位 2，加上座位 5，就變成 3, 4, 5

這個「視窗」就像一個固定大小的框框在陣列上滑動，每次只更新進出的元素，不重新計算整個視窗。

為什麼要用 Sliding Window？
核心目的：把 O(n²) 暴力解 → 優化成 O(n)
暴力解通常是「兩層 for 迴圈」去檢查所有子陣列／子字串，但 Sliding Window 利用「視窗滑動時，重複的部分不用重算」這個特性來省時間。

Sliding Window 的兩種類型
類型視窗大小用途固定視窗固定 k找連續 k 個元素的最大和、平均值等動態視窗會伸縮找最長/最短滿足某條件的子陣列
固定視窗比較簡單，我們先從這個開始。

Sliding Window 的通用模板（固定視窗）
1. 先計算前 k 個元素的「初始視窗值」
2. 從第 k 個位置開始往右滑動：
   - 減掉視窗左邊離開的元素
   - 加上視窗右邊新進來的元素
   - 更新答案
3. 回傳結果

🎯 第二部分：例題實戰
經典例題：找連續 k 個數字的最大總和
題目：給你一個陣列 nums = [2, 1, 5, 1, 3, 2] 和數字 k = 3，找出「連續 3 個數字」加起來的最大值。
答案：5 + 1 + 3 = 9

🐢 暴力解法（先看為什麼慢）
pythondef maxSum_brute(nums, k):
    max_sum = 0
    for i in range(len(nums) - k + 1):  # 外層：每個起點
        current_sum = 0
        for j in range(i, i + k):       # 內層：累加 k 個元素
            current_sum += nums[j]
        max_sum = max(max_sum, current_sum)
    return max_sum
問題：每次都重新加 k 個元素，時間複雜度 O(n × k)。

🚀 Sliding Window 解法
pythondef maxSum_sliding(nums, k):
    # Step 1: 計算第一個視窗的總和
    window_sum = sum(nums[:k])
    max_sum = window_sum
    
    # Step 2: 從第 k 個位置開始滑動
    for i in range(k, len(nums)):
        window_sum += nums[i]       # 加入新元素
        window_sum -= nums[i - k]   # 移除舊元素
        max_sum = max(max_sum, window_sum)
    
    return max_sum
時間複雜度：O(n) ✨

📊 圖解步驟（用 nums = [2, 1, 5, 1, 3, 2], k = 3）
初始視窗：計算前 3 個元素
索引:    0   1   2   3   4   5
陣列:  [ 2,  1,  5,  1,  3,  2 ]
        ↑       ↑
       [█████████]              ← 視窗
       
window_sum = 2 + 1 + 5 = 8
max_sum = 8

Step 1：視窗右移一格（i = 3）
索引:    0   1   2   3   4   5
陣列:  [ 2,  1,  5,  1,  3,  2 ]
            ↑       ↑
           [█████████]          ← 視窗右移
        ✗                       ← 移除 nums[0] = 2
                    ✓           ← 加入 nums[3] = 1

window_sum = 8 - 2 + 1 = 7
max_sum = max(8, 7) = 8

Step 2：視窗再右移（i = 4）
索引:    0   1   2   3   4   5
陣列:  [ 2,  1,  5,  1,  3,  2 ]
                ↑       ↑
               [█████████]      ← 視窗右移
            ✗                   ← 移除 nums[1] = 1
                        ✓       ← 加入 nums[4] = 3

window_sum = 7 - 1 + 3 = 9
max_sum = max(8, 9) = 9  ⭐

Step 3：視窗最後一次右移（i = 5）
索引:    0   1   2   3   4   5
陣列:  [ 2,  1,  5,  1,  3,  2 ]
                    ↑       ↑
                   [█████████]  ← 視窗右移
                ✗               ← 移除 nums[2] = 5
                            ✓   ← 加入 nums[5] = 2

window_sum = 9 - 5 + 2 = 6
max_sum = max(9, 6) = 9

最終答案：max_sum = 9 ✅
🧠 第三部分：心法整理
1. 辨識題型的關鍵字
看到這些字眼，先想 Sliding Window：

「連續子陣列／子字串」
「長度為 k 的最大／最小值」
「最長/最短滿足某條件的子字串」

2. 核心思想：避免重複計算
舊視窗 → 新視窗，只變動進出的兩個元素，中間部分不要動。
3. 不要重置變數
新手常犯錯：每次迴圈都重新 current_sum = 0，那就變回暴力解了。視窗值要累積更新。
4. 固定視窗的實作公式
window_sum += nums[i]        # 右邊進
window_sum -= nums[i - k]    # 左邊出
記住這兩行，固定視窗的題目八成都這樣寫。
5. 動手畫圖
寫不出來時，就在紙上畫陣列 + 視窗框框，一步步推。LeetCode 的滑動視窗題目大多不複雜，卡住就是因為腦中沒有那個視覺畫面。

🎁 給你的下一步建議
學完固定視窗後，下一個重點是 動態視窗（變長視窗），那才是 LeetCode 真正常考的型態，例如：

LeetCode 3: Longest Substring Without Repeating Characters
LeetCode 209: Minimum Size Subarray Sum
LeetCode 76: Minimum Window Substring（這題是大魔王）