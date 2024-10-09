package others

// LeetCode-146: LRU 缓存. [https://leetcode.cn/problems/lru-cache/description/]
/* 描述：
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。
如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/
type LRUCache struct {
	Capacity    int
	LRUNodeMap  map[int]*BiLinkNode
	LRUListHead *BiLinkNode
	LRUListTail *BiLinkNode
}

type BiLinkNode struct {
	Key  int
	Val  int
	Pre  *BiLinkNode
	Next *BiLinkNode
}

func Constructor(capacity int) LRUCache {
	// 思路：掌握关键的思路，LRUNodeMap: make(map[int]*BiLinkNode, capacity)，以及两个注意点即可.
	// 在进行LRUList节点操作的时候，要保证思路清晰，可以适当辅助画图。
	dummyH := &BiLinkNode{}
	dummyT := &BiLinkNode{}

	// 注意：dummy节点也要初始化
	dummyH.Next = dummyT
	dummyT.Pre = dummyH

	return LRUCache{
		Capacity:    capacity,
		LRUNodeMap:  make(map[int]*BiLinkNode, capacity),
		LRUListHead: dummyH,
		LRUListTail: dummyT,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.LRUNodeMap[key] == nil {
		return -1
	}

	old := this.LRUNodeMap[key]
	PreN := old.Pre
	NextN := old.Next

	if NextN == this.LRUListTail {
		return old.Val
	}

	// 删除原节点
	PreN.Next = NextN
	NextN.Pre = PreN

	// 加入到末尾
	preTail := this.LRUListTail.Pre
	preTail.Next = old
	this.LRUListTail.Pre = old
	old.Next = this.LRUListTail
	old.Pre = preTail

	return old.Val
}

func (this *LRUCache) Put(key int, value int) {
	if this.LRUNodeMap[key] != nil {
		this.LRUNodeMap[key].Val = value
		this.Get(key)
		return
	}

	if len(this.LRUNodeMap)+1 > this.Capacity {
		preHead := this.LRUListHead.Next
		this.LRUListHead.Next = preHead.Next
		// 注意：是preHead.Next.Pre而不是preHead.Pre
		preHead.Next.Pre = this.LRUListHead

		delete(this.LRUNodeMap, preHead.Key)
	}

	preTail := this.LRUListTail.Pre
	this.LRUNodeMap[key] = &BiLinkNode{
		Key:  key,
		Val:  value,
		Pre:  preTail,
		Next: this.LRUListTail,
	}
	this.LRUListTail.Pre = this.LRUNodeMap[key]
	preTail.Next = this.LRUNodeMap[key]
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
