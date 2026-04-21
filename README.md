# Leetcode By Golang

## Easy
- 1 Two Sum By Array
- 21 Merge Two Sorted Lists
- 50 Pow(x, n)
- 104 Maximum Depth of Binary Tree
- 141 Linked List Cycle
- 206 Reverse Linked List
- 455 Assign Cookies
- 696 Count Binary Substrings
- 700 Search in a Binary Search Tree
- 704 Binary Search
- 872 Leaf-Similar Trees
- 1009 Complement of Base 10 Integer
- 3783 Mirror Distance of an Integer

## Medium
- 2 Add Two Numbers By Linked List
- 3 Longest Substring Without Repeating Characters
- 5 Longest Palindromic Substring
- 19 Remove Nth Node From End of List
- 33 Search in Rotated Sorted Array
- 45 Jump Game II
- 81 Search In Rotated Sorted Array II
- 128 Longest Consecutive Sequence
- 133 Clone Graph
- 134 Gas Station
- 150 Evaluate Reverse Polish Notation
- 152 Maximum Product Subarray
- 167 Two Sum II - Input Array Is Sorted
- 200 Number of Islands
- 437 Path Sum III
- 435 Non-overlapping Intervals
- 450 Delete Node in a BST
- 1448 Count Good Nodes in Binary Tree
- 1497 Check If Array Pairs Are Divisible by k
- 1589 Maximum Sum Obtained of Any Permutation
- 2461 Maximum Sum of Distinct Subarrays With Length K

## Hard



## Category
## Dynamic Programming（動態規劃）
- 62 Unique Paths
- 70 Climbing Stairs
- 509 Fibonacci Number
- 198 House Robber
- 2061 Number of Spaces Cleaning Robot Cleaned

## Greedy
- Calculate Maximum Profit

---

## 🗺️ 全部模式一覽

| # | 模式 | 英文 | 一句話說明 | 面試重要度 |
|---|------|------|-----------|-----------|
| 1 | 滑動視窗 | Sliding Window | 視窗在陣列上滑動，避免重複計算 | ⭐⭐⭐ 必考 |
| 2 | 雙指針 | Two Pointers | 兩個指針從不同方向逼近答案 | ⭐⭐⭐ 必考 |
| 3 | 動態規劃 | Dynamic Programming | 記住子問題的答案，避免重複計算 | ⭐⭐⭐ 必考 |
| 4 | 二分搜尋 | Binary Search | 在排序資料中用「砍半」快速找目標 | ⭐⭐⭐ 必考 |
| 5 | BFS（廣度優先搜尋） | Breadth-First Search | 一層一層往外探索，找最短路徑 | ⭐⭐⭐ 必考 |
| 6 | DFS（深度優先搜尋） | Depth-First Search | 一條路走到底再回頭，用在樹/圖/排列組合 | ⭐⭐⭐ 必考 |
| 7 | 雜湊表技巧 | Hash Map Patterns | 用 map 做計數、配對、去重 | ⭐⭐⭐ 必考 |
| 8 | 前綴和 | Prefix Sum | 預處理陣列，讓區間加總變 O(1) | ⭐⭐⭐ 必考 |
| 9 | 單調堆疊 | Monotonic Stack | 維護遞增/遞減的 stack，找「下一個更大/更小」 | ⭐⭐⭐ 必考 |
| 10 | 鏈結串列技巧 | Linked List Patterns | 快慢指針、反轉、合併鏈結串列 | ⭐⭐⭐ 必考 |
| 11 | 樹的遍歷 | Tree Traversal | 前序/中序/後序/層序遍歷 | ⭐⭐⭐ 必考 |
| 12 | 貪心演算法 | Greedy | 每一步都選當下最好的，不回頭 | ⭐⭐ 常考 |
| 13 | 回溯法 | Backtracking | DFS + 撤銷選擇，暴力窮舉所有可能 | ⭐⭐ 常考 |
| 14 | 堆積 / 優先佇列 | Heap / Priority Queue | 快速取最大/最小值，用在 Top K 問題 | ⭐⭐ 常考 |
| 15 | 區間問題 | Intervals | 合併/插入/判斷重疊區間 | ⭐⭐ 常考 |
| 16 | 拓撲排序 | Topological Sort | 有先後順序的依賴關係排序（如課程安排） | ⭐ 加分 |
| 17 | 聯合查找 | Union Find (Disjoint Set) | 快速判斷兩個節點是否連通 | ⭐ 加分 |
| 18 | 分治法 | Divide and Conquer | 拆成子問題 → 各自解 → 合併結果 | ⭐ 加分 |
| 19 | Trie 前綴樹 | Trie | 高效字串前綴搜尋（自動補全） | ⭐ 加分 |
| 20 | 位元操作 | Bit Manipulation | 用位元運算解 XOR、子集、狀態壓縮 | ⭐ 加分 |
| 21 | 圖論演算法 | Graph Algorithms | Dijkstra、Floyd、最短路徑等 | ⭐ 加分 |

---

## 🎯 面試優先順序（後端工程師）

### 第一梯隊 — 必考（先搞定這些）

1. Sliding Window（滑動視窗）
2. Two Pointers（雙指針）
3. Dynamic Programming（動態規劃）
4. Binary Search（二分搜尋）
5. BFS / DFS（廣度 / 深度優先搜尋）
6. Hash Map Patterns（雜湊表技巧）
7. Prefix Sum（前綴和）
8. Monotonic Stack（單調堆疊）
9. Linked List Patterns（鏈結串列技巧）
10. Tree Traversal（樹的遍歷）

### 第二梯隊 — 常考（穩固基礎後學這些）

11. Greedy（貪心演算法）
12. Backtracking（回溯法）
13. Heap / Priority Queue（堆積 / 優先佇列）
14. Intervals（區間問題）

### 第三梯隊 — 加分（有餘力再碰）

15. Topological Sort（拓撲排序）
16. Union Find（聯合查找）
17. Divide and Conquer（分治法）
18. Trie（前綴樹）
19. Bit Manipulation（位元操作）
20. Graph Algorithms（圖論演算法）
