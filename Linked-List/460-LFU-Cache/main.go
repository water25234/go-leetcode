package main

// Design and implement a data structure for a Least Frequently Used (LFU) cache.

// Implement the LFUCache class:

// LFUCache(int capacity) Initializes the object with the capacity of the data structure.
// int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
// void put(int key, int value) Update the value of the key if present, or inserts the key if not already present. When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item. For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.
// To determine the least frequently used key, a use counter is maintained for each key in the cache. The key with the smallest use counter is the least frequently used key.

// When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter for a key in the cache is incremented either a get or put operation is called on it.

// The functions get and put must each run in O(1) average time complexity.

// Example 1:

// Input
// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
// Output
// [null, null, null, 1, null, -1, 3, null, -1, 3, 4]

// Explanation
// // cnt(x) = the use counter for key x
// // cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
// LFUCache lfu = new LFUCache(2);
// lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
// lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
// lfu.get(1);      // return 1
//                  // cache=[1,2], cnt(2)=1, cnt(1)=2
// lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
//                  // cache=[3,1], cnt(3)=1, cnt(1)=2
// lfu.get(2);      // return -1 (not found)
// lfu.get(3);      // return 3
//                  // cache=[3,1], cnt(3)=2, cnt(1)=2
// lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
//                  // cache=[4,3], cnt(4)=1, cnt(3)=2
// lfu.get(1);      // return -1 (not found)
// lfu.get(3);      // return 3
//                  // cache=[3,4], cnt(4)=1, cnt(3)=3
// lfu.get(4);      // return 4
//                  // cache=[4,3], cnt(4)=2, cnt(3)=3

// Constraints:

// 1 <= capacity <= 104
// 0 <= key <= 105
// 0 <= value <= 109
// At most 2 * 105 calls will be made to get and put.

// type LFUCache struct {

// }

// func Constructor(capacity int) LFUCache {

// }

// func (this *LFUCache) Get(key int) int {

// }

// func (this *LFUCache) Put(key int, value int)  {

// }

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type Node struct {
	key  int
	val  int
	freq int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewList() *DoublyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{
		head: head,
		tail: tail,
	}
}

func (l *DoublyLinkedList) AddLast(node *Node) {
	node.prev = l.tail.prev
	node.next = l.tail
	l.tail.prev.next = node
	l.tail.prev = node
	l.size++
}

func (l *DoublyLinkedList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
	l.size--
}

func (l *DoublyLinkedList) RemoveFirst() *Node {
	if l.size == 0 {
		return nil
	}
	first := l.head.next
	l.Remove(first)
	return first
}

type LFUCache struct {
	capacity   int
	size       int
	minFreq    int
	keyToNode  map[int]*Node
	freqToList map[int]*DoublyLinkedList
}

// main
func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		size:       0,
		minFreq:    0,
		keyToNode:  make(map[int]*Node),
		freqToList: make(map[int]*DoublyLinkedList),
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.keyToNode[key]
	if !ok {
		return -1
	}

	this.increaseFreq(node)
	return node.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if node, ok := this.keyToNode[key]; ok {
		node.val = value
		this.increaseFreq(node)
		return
	}

	if this.size == this.capacity {
		list := this.freqToList[this.minFreq]
		toRemove := list.RemoveFirst()
		delete(this.keyToNode, toRemove.key)
		this.size--
	}

	newNode := &Node{
		key:  key,
		val:  value,
		freq: 1,
	}

	if this.freqToList[1] == nil {
		this.freqToList[1] = NewList()
	}
	this.freqToList[1].AddLast(newNode)
	this.keyToNode[key] = newNode
	this.minFreq = 1
	this.size++
}

func (this *LFUCache) increaseFreq(node *Node) {
	oldFreq := node.freq
	oldList := this.freqToList[oldFreq]
	oldList.Remove(node)

	if oldFreq == this.minFreq && oldList.size == 0 {
		this.minFreq++
	}

	node.freq++

	if this.freqToList[node.freq] == nil {
		this.freqToList[node.freq] = NewList()
	}
	this.freqToList[node.freq].AddLast(node)
}
