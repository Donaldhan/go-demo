package structx

import (
	"fmt"
	"log"
)

type Node struct {
	data int
	next *Node
}

func Shownode(p *Node) { //遍历
	for p != nil {
		fmt.Println("Shownode", *p)
		p = p.next //移动指针
	}
}

// 单向链表中每个结点包含两部分，分别是数据域和指针域，上一个结点的指针指向下一结点，依次相连，形成链表。

// 这里介绍三个概念：首元结点、头结点和头指针。
// 首元结点：就是链表中存储第一个元素的结点，如下图中 a1 的位置。
// 头结点：它是在首元结点之前附设的一个结点，其指针域指向首元结点。头结点的数据域可以存储链表的长度或者其它的信息，也可以为空不存储任何信息。
// 头指针：它是指向链表中第一个结点的指针。若链表中有头结点，则头指针指向头结点；若链表中没有头结点，则头指针指向首元结点。

// 头结点在链表中不是必须的，但增加头结点有以下几点好处：
// 增加了头结点后，首元结点的地址保存在头结点的指针域中，对链表的第一个数据元素的操作与其他数据元素相同，无需进行特殊处理。
// 增加头结点后，无论链表是否为空，头指针都是指向头结点的非空指针，若链表为空的话，那么头结点的指针域为空。

// 头插法
func InsertFrontLinkListTest() {
	log.Println("====InsertFrontLinkListTest===")
	var head = new(Node)
	head.data = 0
	var tail *Node
	for i := 1; i < 10; i++ {
		var node = Node{data: i}
		node.next = head //将新插入的node的next指向头结点
		head = &node     //重新赋值头结点
		if 1 == 1 {
			tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
		}
	}
	Shownode(head) //遍历结果
	log.Println("====InsertFrontLinkListTest from tail===")
	Shownode(tail) //遍历结果
}

// 尾插法,错误，有问题
func InsertTailLinkListTest() {
	log.Println("====InsertTailLinkListTest===")
	var head = new(Node)
	head.data = 0
	var tail *Node
	tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = Node{data: i}
		node.next = tail //将新插入的node的next指向头结点
		tail = &node     //重新赋值头结点
	}
	Shownode(tail) //遍历结果
}
