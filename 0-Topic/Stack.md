# Stack（堆疊）完整教學

---

## 📖 第一部分：什麼是 Stack？

### 生活化比喻

想像你在**疊盤子**：

- 洗好一個盤子 → 放到盤子塔**最上面**
- 要拿盤子 → 從**最上面**拿一個

**後放的先拿、先放的後拿**，這就是 Stack 的特性 —— **LIFO（Last In, First Out，後進先出）**。

### Stack 的核心操作

| 操作 | 時間 | 說明 |
|------|------|------|
| **Push** | O(1) | 推入棧頂 |
| **Pop** | O(1) | 從棧頂取出 |
| **Peek / Top** | O(1) | 看棧頂但不取 |
| **IsEmpty** | O(1) | 檢查空不空 |

### Go 用 slice 實作 Stack

Go 沒有內建 stack，通常直接用 slice：

```go
stack := []int{}

// Push
stack = append(stack, 1)

// Peek（看棧頂）
top := stack[len(stack)-1]

// Pop（取出棧頂）
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]

// IsEmpty
empty := len(stack) == 0
```

### Stack 的典型使用場景

| 場景 | 為什麼用 Stack | 代表題 |
|------|---------------|--------|
| **括號/符號匹配** | 最近的開括號 = 最深的未配對 | LeetCode 20 |
| **運算式計算** | 運算順序天然後進先出 | LeetCode 150 |
| **函式呼叫 / 遞迴迭代化** | call stack 就是 stack | 前序/中序遍歷迭代版 |
| **保存「上一個」狀態** | LIFO 找最近的歷史 | Min Stack |
| **單調性維護** | 見 `Monotonic-Stack.md` | 日溫度 |

### Stack vs Queue

| | Stack | Queue |
|---|------|------|
| 順序 | LIFO（後進先出） | FIFO（先進先出） |
| 用途 | DFS、歷史回溯、匹配 | BFS、排程、緩衝 |
| 類比 | 疊盤子、返回鍵 | 排隊買票 |

---

## 🎯 第二部分：例題實戰

---

## 例題一：有效的括號（LeetCode 20）— Stack 入門必會

**題目**：給一個字串，只含 `()`, `[]`, `{}` 三種括號。判斷字串是否**有效**（所有括號正確配對、正確嵌套）。

```
"()[]{}"   → true
"(]"       → false
"({[)}"    → false
"({[]})"   → true
```

### 🤔 思路

**最近開的括號要最先配對**—— 這就是 LIFO！

- 遇到**開括號**：push 進 stack
- 遇到**閉括號**：看 stack 頂是不是對應的開括號
  - 是 → pop
  - 否 → 直接回傳 false

最後檢查 stack 是否為空（剩下的就是沒配對的開括號）。

### 🚀 Go 解法

```go
func isValid(s string) bool {
    stack := []byte{}
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for i := 0; i < len(s); i++ {
        c := s[i]
        if opener, isCloser := pairs[c]; isCloser {
            // 閉括號
            if len(stack) == 0 || stack[len(stack)-1] != opener {
                return false
            }
            stack = stack[:len(stack)-1] // pop
        } else {
            // 開括號
            stack = append(stack, c)
        }
    }

    return len(stack) == 0
}
```

### 📊 圖解：`s = "({[]})"`

```
i=0 '(': push       stack = ['(']
i=1 '{': push       stack = ['(', '{']
i=2 '[': push       stack = ['(', '{', '[']
i=3 ']': 頂=']'配對  pop    stack = ['(', '{']
i=4 '}': 頂='{'配對  pop    stack = ['(']
i=5 ')': 頂='('配對  pop    stack = []

結束 stack 空 → true ✅
```

### 📊 反例：`s = "({[)}"`

```
i=0 '(': push       stack = ['(']
i=1 '{': push       stack = ['(', '{']
i=2 '[': push       stack = ['(', '{', '[']
i=3 ')': 頂='['不配對 → return false ✅
```

---

## 例題二：最小堆疊（LeetCode 155）— 設計類 Stack

**題目**：設計一個 stack，除了 push/pop/top，還要 `getMin()` 在 **O(1)** 時間拿到當前最小值。

### 🤔 思路：兩個 stack

- **主 stack**：正常存所有元素
- **輔助 min stack**：存「截至目前為止的最小值」

每次 push 時，把「當前最小」也同步 push 到 min stack 的頂端。pop 時兩邊都 pop。

### 🚀 Go 解法

```go
type MinStack struct {
    stack    []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{}
}

func (s *MinStack) Push(val int) {
    s.stack = append(s.stack, val)

    currMin := val
    if len(s.minStack) > 0 && s.minStack[len(s.minStack)-1] < val {
        currMin = s.minStack[len(s.minStack)-1]
    }
    s.minStack = append(s.minStack, currMin)
}

func (s *MinStack) Pop() {
    s.stack = s.stack[:len(s.stack)-1]
    s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
    return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
    return s.minStack[len(s.minStack)-1]
}
```

### 📊 圖解

```
操作                stack        minStack
Push(3)             [3]          [3]
Push(5)             [3,5]        [3,3]     ← min 還是 3
Push(2)             [3,5,2]      [3,3,2]   ← min 變 2
Push(1)             [3,5,2,1]    [3,3,2,1] ← min 變 1
GetMin() → 1
Pop()               [3,5,2]      [3,3,2]
GetMin() → 2  ✅
```

### 🔑 為什麼 min stack 能「自動還原」？

因為 min stack 跟主 stack 同步推拉，每次 pop 都會把當時的「min 快照」也拿掉。**當時是 2 的最小值，pop 到那時自然又變回 2**。

---

## 例題三：逆波蘭運算式（LeetCode 150）— 運算式計算

**題目**：計算**逆波蘭（後綴）**運算式的值。`["2","1","+","3","*"]` 代表 `(2+1)*3 = 9`。

```
tokens = ["2", "1", "+", "3", "*"] → 9
tokens = ["4", "13", "5", "/", "+"] → 6 （4 + (13/5) = 4+2 = 6）
```

### 🤔 思路

後綴運算式天生為 stack 設計：

- 遇到**數字** → push
- 遇到**運算符** → pop 兩個數字做運算，結果再 push 回去

### 🚀 Go 解法

```go
import "strconv"

func evalRPN(tokens []string) int {
    stack := []int{}

    for _, token := range tokens {
        switch token {
        case "+", "-", "*", "/":
            // 注意：先 pop 的是右運算元
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2]

            var result int
            switch token {
            case "+":
                result = a + b
            case "-":
                result = a - b
            case "*":
                result = a * b
            case "/":
                result = a / b
            }
            stack = append(stack, result)
        default:
            num, _ := strconv.Atoi(token)
            stack = append(stack, num)
        }
    }

    return stack[0]
}
```

### 📊 圖解：`["2", "1", "+", "3", "*"]`

```
"2": push           stack = [2]
"1": push           stack = [2, 1]
"+": pop 1, pop 2, 2+1=3, push   stack = [3]
"3": push           stack = [3, 3]
"*": pop 3, pop 3, 3*3=9, push   stack = [9]

回傳 9 ✅
```

### 🔑 注意運算子的運算元順序

```go
b := stack[len(stack)-1] // 後 pop 的是第二運算元
a := stack[len(stack)-2] // 先 pop 的是第一運算元

// a - b 才對，不是 b - a
// a / b 才對，不是 b / a
```

加法和乘法交換律沒差，**減法和除法會搞錯**，面試容易掉這個坑。

---

## 🧠 第三部分：心法整理

### 1. 辨識題型的關鍵字

看到這些，先想 Stack：

- 「**括號匹配**」「**符號配對**」
- 「**後綴/前綴運算式**」
- 「**最近的 X**」（最近開括號、最近打開的標籤）
- 「**撤銷**操作」「**回上一步**」
- 「迭代化**遞迴**」（DFS 不想用遞迴時）

### 2. Stack 的經典 pattern

**配對型**：

```go
stack := []byte{}
for _, c := range s {
    if /* 是開始符號 */ {
        stack = append(stack, c)
    } else if /* 是結束符號 */ {
        if len(stack) == 0 || /* 不匹配 */ {
            return false
        }
        stack = stack[:len(stack)-1]
    }
}
return len(stack) == 0
```

**運算式求值型**：

```go
stack := []int{}
for _, token := range tokens {
    if /* 是數字 */ {
        stack = append(stack, toInt(token))
    } else {
        b := stack[len(stack)-1]
        a := stack[len(stack)-2]
        stack = stack[:len(stack)-2]
        stack = append(stack, operate(a, b, token))
    }
}
return stack[0]
```

### 3. Stack 寫 DFS 的好處

遞迴版 DFS 有「**爆 stack 風險**」（樹太深 / 圖太複雜）。用 stack 手動模擬可以：

- 不怕函式呼叫深度爆掉
- 完全控制「何時繼續、何時回退」

缺點是程式碼變長、變醜。**面試通常先寫遞迴版，面試官問才改迭代版**。

### 4. Stack 配 HashMap 的組合

```go
pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
```

像括號匹配這種，用 map 存配對關係比多個 switch case 乾淨很多。

### 5. Stack 和 Monotonic Stack 的關係

**Monotonic Stack 是 Stack 的特化版**：多加「維持 stack 單調性」這條規則。進階題型大多用 Monotonic Stack 而不是普通 Stack（見 `Monotonic-Stack.md`）。

### 6. 面試小 tips

- Push/Pop 的空檢查：`len(stack) > 0` 寫在最前面，避免 panic
- Go 的 slice 當 stack 很方便，不用用 `container/list`
- Stack 的元素型別可以是 struct、interface{}、指標——靈活度很高

---

## 📚 推薦練習題（由易到難）

| 難度 | 題目 | 類型 |
|------|------|------|
| Easy | LeetCode 20. Valid Parentheses | 配對 |
| Easy | LeetCode 155. Min Stack | 設計類 |
| Easy | LeetCode 232. Implement Queue using Stacks | Stack 組 Queue |
| Easy | LeetCode 844. Backspace String Compare | Stack 模擬退格 |
| Medium | LeetCode 150. Evaluate Reverse Polish Notation | 運算式 |
| Medium | LeetCode 71. Simplify Path | 字串 stack 處理 |
| Medium | LeetCode 394. Decode String | Stack 解編碼 |
| Medium | LeetCode 227. Basic Calculator II | 計算機 |
| Hard | LeetCode 224. Basic Calculator | 含括號計算機 |
| Hard | LeetCode 42. Trapping Rain Water | Stack 解法之一 |
