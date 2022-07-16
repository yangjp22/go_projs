package link

import (
	"errors"
	"fmt"
)

// LinkNode 单向链表
type LinkNode struct {
	Data int64
	NextNode *LinkNode
}

// PrintLinkNode 打印出单链表所有元素
func (lnode *LinkNode) PrintLinkNode() {
	nowNode := lnode
	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
		} else {
			break
		}
	}
}

// Length 返回单向链表的长度
func (lnode *LinkNode) Length() int {
	nowNode := lnode
	length := 0
	for {
		if nowNode != nil {
			length += 1
			nowNode = nowNode.NextNode
		} else {
			break
		}
	}
	return length
}

// TestLinkNode 单链表的测试
func TestLinkNode() {
	// 新的节点
	node := new(LinkNode)
	node.Data = 2

	// 下一个节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1   // 将node1 连接到node上

	// 下一个节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2 // 将node2 链接到node1上

	node.PrintLinkNode()
	fmt.Println("length of node is", node.Length())
}

// Ring 循环链表
type Ring struct {
	next, prev *Ring   // 前后驱节点
	Value interface{}    // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// NewNRing 创建N个节点的循环链表
func NewNRing(n int) (*Ring, error) {
	if n <= 0 {
		return nil, errors.New("Invalid n")
	}

	r := new(Ring)
	p := r
	for i := 0; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r, nil
}

// Next 返回下一个节点
func (r *Ring) Next() *Ring {
	if r.next != nil {
		return r.next
	}
	return r.init()
}

// Prev 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// Move 获取循环链表的第N个节点
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {   // 当n为负数时，表示从前面往前遍历，否则往后遍历
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// AddRing 在r节点后插入一个新的节点，返回插入之前r的第一个后驱节点
func (r *Ring) AddRing(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// RemoveRing 删除节点后面的n个节点
func (r *Ring) RemoveRing (n int) *Ring{
	if n < 0 {
		return nil
	}
	return r.AddRing(r.Move(n + 1))
}

// Length 返回循环链表的长度
func (r *Ring) Length() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}