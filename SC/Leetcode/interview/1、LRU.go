package interview

// 核心结构：哈希表 + 双向链表
type Node struct { // 节点口诀：前后键值四件套
	Key, Val   int
	Prev, Next *Node
}

type _LRUCache struct {
	cap        int           // 容量
	cache      map[int]*Node // 哈希表存节点
	head, tail *Node         // 头尾哨兵节点
}

// 初始化口诀：造头尾，连彼此
func lru_Constructor(cap int) _LRUCache {
	lru := _LRUCache{
		cap:   cap,
		cache: make(map[int]*Node),
		head:  &Node{}, // 头哨兵
		tail:  &Node{}, // 尾哨兵
	}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

// Get 口诀：查哈希，有则提(删旧+插头)
func (c *_LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.remove(node)    // 删旧位置
		c.addToHead(node) // 插到头部
		return node.Val
	}
	return -1
}

// Put 口诀：分存在/不存在，超容删尾
func (c *_LRUCache) Put(key, val int) {
	// 已存在：更新值并提到头部
	if node, ok := c.cache[key]; ok {
		node.Val = val
		c.remove(node)
		c.addToHead(node)
		return
	}

	// 不存在：创建新节点，插入头部
	newNode := &Node{Key: key, Val: val}
	c.cache[key] = newNode
	c.addToHead(newNode)

	// 超容时：删尾节点(最久未用)
	if len(c.cache) > c.cap {
		last := c.tail.Prev
		delete(c.cache, last.Key)
		c.remove(last)
	}
}

// 链表操作口诀：头插四步走，删节点两线连
func (c *_LRUCache) addToHead(node *Node) {
	node.Next = c.head.Next // 新节点指向头后
	node.Prev = c.head      // 新节点前驱指向头
	c.head.Next.Prev = node // 原头后节点前驱指向新节点
	c.head.Next = node      // 头节点指向新节点
}

func (c *_LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next // 前节点直连后节点
	node.Next.Prev = node.Prev // 后节点回连前节点
}
