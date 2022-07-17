package stack_queue

import (
	"errors"
	"sync"
)

// ArrayStack 用数组实现栈
type ArrayStack struct {
	array []string   // 底层切片
	size int         // 栈的元素数量
	lock sync.Mutex  // 为了并发安全
}

// Push 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.array = append(stack.array, v)
	stack.size = stack.size + 1
}

// Pop 出栈
func (stack *ArrayStack) Pop() (string, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		return "", errors.New("empty")
	}

	// 栈顶元素
	v := stack.array[stack.size - 1]

	//
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size - 1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray
	stack.size -= 1
	return v, nil
}

// Peek 返回栈顶元素 但不出栈
func (stack *ArrayStack) Peek() string {
	if stack.size == 0 {
		return "Empty stack"
	}
	return stack.array[stack.size - 1]
}

// Size 返回栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}


type LinkNode struct {
	Next *LinkNode
	Value string
}

// LinkStack 实现的栈
type LinkStack struct {
	root *LinkNode  // 链表起点
	size int             // 栈大小
	lock sync.Mutex      // 为了并发安全的锁
}

// Push 入栈
func (stack *LinkStack) Push(element string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈顶为空
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = element
	} else {
		node := new(LinkNode)
		node.Value = element
		node.Next = stack.root
		// 新节点放在头部
		stack.root = node
	}
	stack.size += 1
}

// Pop 出栈
func (stack *LinkStack) Pop() (string, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		return "", errors.New("empty stack")
	}

	topNode := stack.root.Next
	value := stack.root.Value
	stack.root = topNode
	stack.size -= 1

	return value, nil
}

// Peek 返回栈顶元素
func (stack *LinkStack) Peek() (string, error) {
	if stack.size == 0 {
		return "", errors.New("empty stack")
	}
	return stack.root.Value, nil
}

// Size 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}
//
//func TestLinkStack() {
//	linkStack := new(LinkStack)
//	linkStack.Push("cat")
//	linkStack.Push("dog")
//	linkStack.Push("hen")
//	fmt.Println("size:", linkStack.Size())
//	fmt.Println("pop:", linkStack.Pop())
//	fmt.Println("pop:", linkStack.Pop())
//	fmt.Println("size:", linkStack.Size())
//	linkStack.Push("drag")
//	fmt.Println("pop:", linkStack.Pop())
//}


