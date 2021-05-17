package main

//LRU队列
//请使用常用数据结构实现一个基于LRU队列的本地缓存。



type Value struct {
node *Node
}
type Node struct {
key int
prev *Node
next *Node
}

type list struct {
maxSize int
head  *Node
tail *Node
}

type localCache struct {
cacheTable map[int]Value
list  *list
}

func (c *localCache) put(k int, v Value) {
//添加缓存
c.cacheTable[k] = v
//添加列表至尾部
newNode := &Node{key:k}
newNode.prev = c.list.tail
c.list.tail.next = newNode
c.list.tail = newNode

//缓存超出  清除最久未被使用 即队首元素
if len(c.cacheTable) > c.maxSize {
deleteOne := c.list.head
delete(c.cacheTable, deleteOne.key)
c.list.head.prev = nil
c.list.head = c.list.head.next
}
}

func (c *localCache) get(k int) Value {
if  val,ok := c.cacheTable[k]; ok {
//通过key 找链表节点
curNode := val.node
//删除队列当前节点
curNode.next.prev = curNode.prev
curNode.prev.next = curNode.next
//最新查找为热key,追加到最后 curNode.next = nil curNode.prev = tail
c.list.tail = curNode
return val
}

return nil
}

func (c *localCache) del(k int) {
if val.ok := c.cacheTable[k]; ok {
delete(c.cacheTable, k)
//删除当前节点
curNode.next.prev = curNode.prev
curNode.prev.next = curNode.next
}
}




