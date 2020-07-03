package leetcode

//面试题 16.25. LRU缓存
//设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，
//并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。
//
//它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
//示例:
//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4

type lruListNode struct {
	key  int
	val  int
	prev *lruListNode
	next *lruListNode
}

type LRUCache struct {
	dataMap  map[int]*lruListNode
	capacity int
	length   int
	listHead *lruListNode
	listTail *lruListNode
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("capacity must > 0")
	}
	lc := LRUCache{
		dataMap:  make(map[int]*lruListNode),
		capacity: capacity,
		length:   0,
		listHead: nil,
		listTail: nil,
	}

	return lc
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.dataMap[key]
	if !exist {
		return -1
	}

	this.MoverToListHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	_, exist := this.dataMap[key]
	if exist {
		this.dataMap[key].val = value
		this.MoverToListHead(this.dataMap[key])
		return
	}

	if this.length >= this.capacity {
		removedKey := this.RemoveFromTail()
		delete(this.dataMap, removedKey)
		this.length--
	}

	node := &lruListNode{
		key:  key,
		val:  value,
		prev: nil,
		next: nil,
	}

	this.AddToListHead(node)
	this.dataMap[key] = node
	this.length++
}

func (this *LRUCache) AddToListHead(node *lruListNode) {
	if this.length == 0 {
		this.listHead = node
		this.listTail = node
		return
	}

	node.next = this.listHead
	this.listHead.prev = node
	this.listHead = node
}

func (this *LRUCache) MoverToListHead(node *lruListNode) {
	if node == this.listHead {
		return
	}
	if node == this.listTail {
		this.listTail = this.listTail.prev
		this.listTail.next = nil
		this.AddToListHead(node)
		return
	}

	nodeNext := node.next
	node.prev.next = node.next
	nodeNext.prev = node.prev

	this.AddToListHead(node)
}

func (this *LRUCache) RemoveFromTail() (key int) {
	key = this.listTail.key
	if this.listTail.prev == nil {
		this.listTail = nil
		this.listHead = nil
		return key
	}

	tail := this.listTail
	this.listTail = this.listTail.prev
	this.listTail.next = nil
	tail.prev = nil

	return key
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
