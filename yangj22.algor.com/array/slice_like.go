package array

import "sync"

// Array 可变长数组
type Array struct {
	array []int  // 固定大小的数组，用满容量和满大小的切片来代替
	len int  // 真正长度
	cap int  // 容量
	lock sync.Mutex  // 为了并发安全使用的锁
}

// Make 初始化数组
func Make(len, cap int) *Array {
	s := new(Array)

	// 把切片当数组用
	array := make([]int, len, cap)

	// 元数据
	s.array = array
	s.cap = cap
	s.len = len
	return s
}

// Append 将一个元素添加到数组的末尾
func (a *Array) Append(element int) {
	// 并发锁
	a.lock.Lock()
	defer a.lock.Unlock()

	// 大小等于容量，表示没有多余位置
	if a.len == a.cap {
		// 没容量 数组需要扩容 到两倍
		newCap := 2 * a.cap

		// 如果之前的容量为0，那么新容量为1
		if a.cap == 0 {
			newCap = 1
		}

		// 创建新数组
		newArray := make([]int, a.len, newCap)

		// 把老数组的数据移动到新数组
		for k, v := range a.array {
			newArray[k] = v
		}

		// 替换掉数组
		a.array = newArray
		a.cap = newCap
	}

	// 把新元素放到数组里
	a.array[a.len] = element
	a.len += 1
}

// AppendMany 一次性添加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

