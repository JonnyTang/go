package main

import "fmt"

/*
  简单的实现一个单链表
*/

// 所有数据类型
type Object interface{}

// 定义节点
type Node struct {
	data Object
	next *Node
}

// 定义列表
type List struct {
	size uint64 // 链表节点数
	head *Node  // 头节点
	tail *Node  // 尾节点
}

// 初始化函数
func (list *List) Init() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

// 判断链表是否为空
func (list *List) IsEmpty() bool {
	return list.size <= 0
}

// 添加最后
func (list *List) Append(node *Node) bool {
	// 入参校验
	if node == nil {
		return false
	}

	node.next = nil
	if list.size == 0 {
		list.head = node
	} else { // 有元素时，在后面添加
		oldList := list.tail
		oldList.next = node
	}

	list.tail = node
	list.size++ //节点数增加

	return true
}

// 插入到任意位置
func (list *List) Insert(i uint64, node *Node) bool {
	if node == nil || i > list.size || i == 0 || list.size == 0 {
		return false
	}

	if i == 1 { // 插入表头
		node.next = list.head
		list.head = node
	} else {
		// 先找到第i-1个元素
		preNode := list.head
		for j := 1; uint64(j) < i; j++ {
			preNode = preNode.next
		}
		node.next = preNode.next
		preNode.next = node
	}

	list.size++

	return true
}

// 删除每i个元素
func (list *List) Remove(i uint64) bool {
	if list.size == 0 || i > list.size || i == 0 {
		return false
	}

	// 找到前一个元素
	node := list.head
	if list.size == 1 { //只有一个元素的情况
		list.head = nil
		list.tail = nil
	} else {
		for j := 1; uint64(j) < i; j++ {
			node = node.next
		}
		node.next = node.next.next
	}

	list.size--

	return true
}

// 改
func (list *List) Update() bool {

	return true
}

// 查找第i个节点
func (list *List) Find(i uint64) *Node {
	if i > list.size || i == 0 || list.size == 0 {
		return nil
	}

	// 遍历
	target := list.head
	for j := 1; uint64(j) <= i; j++ {
		target = target.next
	}

	return target
}

func main() {
	fmt.Println("this is a list-learning!")

	a := Node{110, nil}
	fmt.Println(a.data)
}
