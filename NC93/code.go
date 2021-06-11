package main

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

type LRUNode struct {
	key        int
	val        int
	next, prev *LRUNode
}

type LRUCache struct {
	cache          map[int]*LRUNode
	head, tail     *LRUNode
	size, capacity int
}

func initCache(k int) *LRUCache {
	l := LRUCache{
		cache: map[int]*LRUNode{},
		head: &LRUNode{
			key: 0,
			val: 0,
		},
		tail: &LRUNode{
			key: 0,
			val: 0,
		},
		size:     0,
		capacity: k,
	}

	l.head.next = l.tail
	l.tail.prev = l.head

	return &l
}

func (cache *LRUCache) set(key, val int) {
	if node, ok := cache.cache[key]; ok {
		cache.moveToHead(node)
		node.val = val
	} else {
		n := LRUNode{
			key: key,
			val: val,
		}
		cache.insertIntoHead(&n)
	}
}

func (cache *LRUCache) get(key int) int {
	if node, ok := cache.cache[key]; ok {
		cache.moveToHead(node)
		return node.val
	} else {
		return -1
	}
}

func (cache *LRUCache) moveToHead(node *LRUNode) {
	cache.removeNode(node)
	cache.insertIntoHead(node)
}

func (cache *LRUCache) insertIntoHead(node *LRUNode) {
	cache.head.next.prev = node
	node.next = cache.head.next
	cache.head.next = node
	node.prev = cache.head
	cache.cache[node.key] = node
	cache.size++
	if cache.size > cache.capacity {
		cache.removeNode(cache.tail.prev)
	}
}

func (cache *LRUCache) removeNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(cache.cache, node.key)
	cache.size--
}

func LRU(operators [][]int, k int) []int {
	// write code here
	ret := []int{}
	cache := initCache(k)

	for _, arr := range operators {
		if arr[0] == 2 {
			ret = append(ret, cache.get(arr[1]))
		} else if arr[0] == 1 {
			cache.set(arr[1], arr[2])
		}
	}
	return ret
}
