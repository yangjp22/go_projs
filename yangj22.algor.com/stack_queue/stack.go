package stack_queue

import "sync"

type ArrayStack struct {
	array []string   // 底层切片
	size int         // 栈的元素数量
	lock sync.Mutex  // 为了并发安全
}



