package cache

import "fmt"

type Node struct {
	key  string
	pre  *Node
	next *Node
}

type LRUCache struct {
	Cap    int              //最大容量
	bucket map[string]*Node //HashMap
	head   *Node
	tail   *Node
	Size   int
}

func NewLruList(cap int) *LRUCache {
	cache := &LRUCache{
		Cap:    cap,
		bucket: make(map[string]*Node, cap),
		head:   &Node{"", nil, nil},
		tail:   &Node{"", nil, nil},
		Size:   0,
	}
	cache.head.next = cache.tail
	cache.tail.pre = cache.head
	return cache
}

//添加一个节点到首位(也就是最近一个访问的)
func (list *LRUCache) addNodeFirst(node *Node) *Node {
	//判断是否到容量上限
	if list.Size == list.Cap {
		//到达上限之后,删除最后一个节点
		deleteNode := list.tail.pre
		list.deleteNode(list.tail.pre)
		return deleteNode
	}
	//添加节点到首位
	node.pre = list.head
	node.next = list.head.next
	list.head.next.pre = node
	list.head.next = node
	//添加该映射
	list.bucket[node.key] = node
	list.Size++
	return nil
}

//将某个key变成最近使用的(假定该key一定存在)
func (list *LRUCache) makeNodeFirst(key string) {
	//根据key获取该节点
	node := list.bucket[key]
	//先删除该节点
	list.deleteNode(node)
	//再加入该节点到首位
	list.addNodeFirst(node)
}

//删除某个节点
func (list *LRUCache) deleteNode(node *Node) {
	//先删除映射
	delete(list.bucket, node.key)
	//从双向链表中删除该节点
	node.pre.next = node.next
	node.next.pre = node.pre
	list.Size--
}

func (list *LRUCache) Put(key string) string {
	//先判断是否已有该节点,有则更新
	node := list.bucket[key]
	if node == nil {
		//当没有该节点时,直接加入到首位
		node := &Node{key, nil, nil}
		deleteNode := list.addNodeFirst(node)
		if deleteNode != nil {
			return deleteNode.key
		}
		return ""
	} else {
		//如果已经有该节点,那么先直接更新该节点的值并且提前至首位
		list.makeNodeFirst(key)
	}
	return ""
}

func (list *LRUCache) Delete(key string) {
	node := list.bucket[key]
	list.deleteNode(node)
}

func (list *LRUCache) Print() {
	cur := list.head.next
	for ; cur != list.tail; cur = cur.next {
		fmt.Printf("%s ", cur.key)
	}
	fmt.Println()
}
