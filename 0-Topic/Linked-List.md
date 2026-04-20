# Linked List（鏈結串列）完整教學

---

## 📖 第一部分：什麼是 Linked List？

### 生活化比喻

想像一列**尋寶遊戲**：

- 每個站點有一張紙條，上面寫著：**這一站的寶物** + **下一站的地址**
- 你從第一站開始，照著紙條一站一站走
- 走到最後一站，紙條寫著「沒有下一站了」（nil）

這就是 Linked List —— 每個節點（Node）存「自己的值」和「下一個節點的指標」。

### Linked List vs Array 的差別

| | Array（陣列） | Linked List（鏈結串列） |
|---|--------------|----------------------|
| **記憶體** | 連續一整塊 | 散落各處，用指標串起來 |
| **查詢第 i 個** | O(1) 直接算位置 | O(n) 要從頭走過去 |
| **插入/刪除** | O(n) 要搬移後面元素 | O(1) 改指標就好 |
| **適合場景** | 頻繁讀取 | 頻繁插入/刪除 |

### Go 中 Linked List 的基本結構

```go
// 節點定義（LeetCode 標準寫法）
type ListNode struct {
    Val  int
    Next *ListNode
}
```

視覺化：

```
[1] → [2] → [3] → [4] → nil
 ↑
head
```

每個框框就是一個 `ListNode`，箭頭就是 `Next` 指標。

### Linked List 題目的四大技巧

| 技巧 | 用途 | 代表題 |
|------|------|--------|
| **快慢指針** | 找中點、判斷有環 | 環形鏈結串列、找中間節點 |
| **反轉鏈結串列** | 翻轉部分或全部節點 | 反轉鏈結串列 |
| **Dummy Head（虛擬頭節點）** | 簡化邊界處理 | 合併排序鏈結串列、刪除節點 |
| **遞迴** | 從尾巴往回處理 | 反轉鏈結串列、兩數相加 |

---

## 🎯 第二部分：例題實戰

---

## 例題一：反轉鏈結串列（LeetCode 206）— 最經典必會

**題目**：把一個鏈結串列反轉過來。

```
輸入：1 → 2 → 3 → 4 → 5 → nil
輸出：5 → 4 → 3 → 2 → 1 → nil
```

### 🤔 思路

想像你在拆一條**鏈子**，每次把當前這一節拆下來，接到一條新鏈子的最前面：

```
原本：  1 → 2 → 3 → 4 → nil
新鏈：  nil

第一步：把 1 拆下來，接到新鏈前面
原本：  2 → 3 → 4 → nil
新鏈：  1 → nil

第二步：把 2 拆下來，接到新鏈前面
原本：  3 → 4 → nil
新鏈：  2 → 1 → nil

...以此類推
```

### 🚀 Go 解法（迭代版）

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil // 新鏈的頭（一開始是 nil）
    curr := head             // 目前要處理的節點

    for curr != nil {
        next := curr.Next // 1. 先記住下一個（不然等等斷掉就找不到了）
        curr.Next = prev  // 2. 把當前節點指向前面（反轉）
        prev = curr       // 3. prev 前進
        curr = next       // 4. curr 前進
    }

    return prev // prev 就是新的 head
}
```

### 📊 圖解步驟

`1 → 2 → 3 → nil`

#### 初始狀態

```
prev    curr
 ↓       ↓
nil     [1] → [2] → [3] → nil
```

#### Step 1：處理節點 1

```
1. next = curr.Next        → next 指向 [2]
2. curr.Next = prev        → [1].Next = nil（斷開，指向 prev）
3. prev = curr             → prev 指向 [1]
4. curr = next             → curr 指向 [2]

       prev    curr
        ↓       ↓
nil ← [1]     [2] → [3] → nil
```

#### Step 2：處理節點 2

```
1. next = curr.Next        → next 指向 [3]
2. curr.Next = prev        → [2].Next = [1]
3. prev = curr             → prev 指向 [2]
4. curr = next             → curr 指向 [3]

              prev    curr
               ↓       ↓
nil ← [1] ← [2]     [3] → nil
```

#### Step 3：處理節點 3

```
1. next = curr.Next        → next = nil
2. curr.Next = prev        → [3].Next = [2]
3. prev = curr             → prev 指向 [3]
4. curr = next             → curr = nil（結束）

                     prev    curr
                      ↓       ↓
nil ← [1] ← [2] ← [3]     nil
```

#### 迴圈結束，回傳 prev

```
[3] → [2] → [1] → nil  ✅
```

### 🔑 關鍵心法：四步口訣

```
存下一個 → 反轉指標 → prev 前進 → curr 前進
next      → reverse   → prev++   → curr++
```

面試寫反轉鏈結串列，就默念這四步，一行一行寫不會錯。

---

## 例題二：判斷鏈結串列有沒有環（LeetCode 141）— 快慢指針經典

**題目**：判斷一個鏈結串列是否有環（某個節點的 Next 指向前面的節點，形成循環）。

```
有環：  1 → 2 → 3 → 4
                ↑       ↓
                └───────┘

無環：  1 → 2 → 3 → nil
```

### 🤔 策略：快慢指針（龜兔賽跑）

想像一個環形跑道：

- **烏龜（slow）**：每次走 1 步
- **兔子（fast）**：每次走 2 步
- 如果有環 → 兔子一定會繞回來追上烏龜
- 如果無環 → 兔子先跑到 nil（終點）

### 🚀 Go 解法

```go
func hasCycle(head *ListNode) bool {
    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next      // 慢指針走 1 步
        fast = fast.Next.Next // 快指針走 2 步

        if slow == fast {
            return true // 相遇了，有環！
        }
    }

    return false // fast 走到 nil，無環
}
```

### 📊 圖解步驟

#### ✅ 有環的情況

```
節點：  1 → 2 → 3 → 4
                ↑       ↓
                └───────┘
（節點 4 的 Next 指向節點 2，形成環）
```

```
初始：    slow=1, fast=1

Step 1:   slow=2, fast=3    （slow 走 1 步，fast 走 2 步）
Step 2:   slow=3, fast=3    （fast 繞了一圈追上來！）
                   ↑
              slow == fast → return true ✅
```

#### ❌ 無環的情況

```
節點：  1 → 2 → 3 → nil
```

```
初始：    slow=1, fast=1

Step 1:   slow=2, fast=3
Step 2:   fast.Next = nil → 迴圈結束 → return false ✅
```

### 🤔 為什麼快慢指針一定會相遇？

如果有環，當兩個指針都進入環之後：

- 每走一步，fast 比 slow **多走 1 步**
- 所以兩者之間的**距離差每次減 1**
- 距離差遲早會變成 0 → **一定會相遇**

就像在環形跑道上，跑得快的人一定會追上（套圈）跑得慢的人。

---

## 例題三：合併兩個排序鏈結串列（LeetCode 21）— Dummy Head 技巧

**題目**：把兩個已排序的鏈結串列合併成一個排序的鏈結串列。

```
輸入：1 → 2 → 4
     1 → 3 → 4

輸出：1 → 1 → 2 → 3 → 4 → 4
```

### 🤔 思路

跟打撲克牌整理手牌一樣：

- 兩疊牌都從最上面翻
- 每次比較兩疊牌頂，**小的那張**先放到新的那疊
- 某一疊翻完了，另一疊剩的直接接上去

### 🔑 Dummy Head 技巧

合併鏈結串列時，「第一個節點」的處理很麻煩（因為還沒有 head）。

**Dummy Head** 就是在最前面放一個**假的節點**，讓所有操作都一樣，最後回傳 `dummy.Next` 就好。

```
dummy → [?] → [1] → [1] → [2] → [3] → [4] → [4] → nil
         ↑
      這個是假的，最後不要它
```

### 🚀 Go 解法

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{} // 虛擬頭節點
    curr := dummy

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            curr.Next = list1    // 接上較小的節點
            list1 = list1.Next   // list1 前進
        } else {
            curr.Next = list2
            list2 = list2.Next
        }
        curr = curr.Next // curr 前進
    }

    // 其中一條還有剩，直接接上
    if list1 != nil {
        curr.Next = list1
    } else {
        curr.Next = list2
    }

    return dummy.Next // 跳過 dummy，回傳真正的 head
}
```

### 📊 圖解步驟

```
list1:  [1] → [2] → [4] → nil
list2:  [1] → [3] → [4] → nil
dummy:  [D]
         ↑
        curr
```

#### Step 1：比較 1 vs 1，list1 <= list2，接 list1

```
list1:        [2] → [4] → nil
list2:  [1] → [3] → [4] → nil
dummy:  [D] → [1]
               ↑
              curr
```

#### Step 2：比較 2 vs 1，list2 較小，接 list2

```
list1:  [2] → [4] → nil
list2:        [3] → [4] → nil
dummy:  [D] → [1] → [1]
                      ↑
                     curr
```

#### Step 3：比較 2 vs 3，list1 較小，接 list1

```
list1:        [4] → nil
list2:        [3] → [4] → nil
dummy:  [D] → [1] → [1] → [2]
                            ↑
                           curr
```

#### Step 4：比較 4 vs 3，list2 較小，接 list2

```
list1:  [4] → nil
list2:        [4] → nil
dummy:  [D] → [1] → [1] → [2] → [3]
                                   ↑
                                  curr
```

#### Step 5：比較 4 vs 4，list1 <= list2，接 list1

```
list1:        nil
list2:  [4] → nil
dummy:  [D] → [1] → [1] → [2] → [3] → [4]
                                         ↑
                                        curr
```

#### Step 6：list1 == nil，把 list2 剩的直接接上

```
dummy:  [D] → [1] → [1] → [2] → [3] → [4] → [4] → nil
```

#### 回傳 dummy.Next

```
[1] → [1] → [2] → [3] → [4] → [4] → nil  ✅
```

---

## 🧠 第三部分：心法整理

### 1. 反轉鏈結串列：四步口訣

```go
next := curr.Next  // 存下一個
curr.Next = prev   // 反轉指標
prev = curr        // prev 前進
curr = next        // curr 前進
```

這四行背起來，反轉鏈結串列的題目都能解。變形題（反轉部分、K 個一組反轉）只是加上邊界控制。

### 2. 快慢指針的三大用途

| 用途 | slow 走幾步 | fast 走幾步 |
|------|-----------|-----------|
| **判斷有環** | 1 | 2 |
| **找中間節點** | 1 | 2（fast 到底時，slow 在中間）|
| **找倒數第 k 個** | 1 | 先走 k 步，然後一起走 1 |

### 3. Dummy Head 是鏈結串列的瑞士刀

只要你的操作會**改動 head**（合併、刪除第一個節點、插入排序等），就加一個 Dummy Head，最後回傳 `dummy.Next`。

好處：不用特別處理「head 是 nil」或「操作的是第一個節點」這種邊界情況。

```go
dummy := &ListNode{}
curr := dummy

// ... 操作完 ...

return dummy.Next
```

### 4. 鏈結串列操作的通用陷阱

**斷鏈問題**：改 Next 之前一定要先保存下一個節點！

```go
// ❌ 錯誤：直接改，後面就找不到了
curr.Next = prev
curr = curr.Next  // curr.Next 已經被改掉了！

// ✅ 正確：先存起來
next := curr.Next
curr.Next = prev
curr = next
```

### 5. 面試畫圖很重要

鏈結串列題在面試白板上，**一定要畫圖**。

畫出每一步的指標變化，面試官看得到你的思路，就算 code 有小 bug 也不會扣太多分。反而什麼都不畫，直接寫 code 容易出錯還解釋不清楚。

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 核心技巧 |
|------|------|---------|
| Easy | LeetCode 206. Reverse Linked List | 反轉四步口訣 |
| Easy | LeetCode 141. Linked List Cycle | 快慢指針 |
| Easy | LeetCode 21. Merge Two Sorted Lists | Dummy Head |
| Easy | LeetCode 876. Middle of the Linked List | 快慢指針找中點 |
| Medium | LeetCode 19. Remove Nth Node From End | 快慢指針找倒數第 k |
| Medium | LeetCode 142. Linked List Cycle II | 快慢指針 + 找環入口 |
| Medium | LeetCode 24. Swap Nodes in Pairs | 兩兩反轉 |
| Hard | LeetCode 25. Reverse Nodes in k-Group | K 個一組反轉 |