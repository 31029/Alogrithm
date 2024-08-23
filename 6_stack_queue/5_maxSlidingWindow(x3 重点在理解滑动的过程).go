package stack_queue

func findLarger(nums []int) int {
	larger := nums[0]
	for _, v := range nums {
		if v > larger {
			larger = v
		}
	}
	return larger
}

func MaxSlidingWindow_Overtime(nums []int, k int) []int {
	// x1 超时
	// x2 剪枝，仍然超时
	result := []int{}

	lastMax := nums[0]
	for i := 0; i < len(nums)-k+1; i++ {
		if i != 0 && i+k < len(nums)-1 {
			n := nums[i+k-1]
			if n  >= lastMax {
				result = append(result, n )
				lastMax = n 
				continue
			} 
		}
		l := findLarger(nums[i: i+k])
		result = append(result, l)
		lastMax = l
	}

	return result
}


func MaxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	growQueue := IntQueueConstructor()
	for i, v := range nums {
		for v > growQueue.Bottom() && !growQueue.Empty() {
			growQueue.PopBottom()
		}
		growQueue.Push(v)

		if i>=k && nums[i-k] == growQueue.Peek() {
			growQueue.Pop()
		}

		if i>=k-1 {
			result = append(result, growQueue.Peek())
		}
	}

	return result
}



// ************************* 封装单调队列的方式解题 ********************************
type MyQueue struct {
    queue []int
}

func NewMyQueue() *MyQueue {
    return &MyQueue{
        queue: make([]int, 0),
    }
}

func (m *MyQueue) Front() int {
    return m.queue[0]
}

func (m *MyQueue) Back() int {
    return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
    return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
    for !m.Empty() && val > m.Back() {
        m.queue = m.queue[:len(m.queue)-1]
    }
    m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
    if !m.Empty() && val == m.Front() {
        m.queue = m.queue[1:]
    }
}

func maxSlidingWindow(nums []int, k int) []int {
    queue := NewMyQueue()
    length := len(nums)
    res := make([]int, 0)
    // 先将前k个元素放入队列
    for i := 0; i < k; i++ {
        queue.Push(nums[i])
    }
    // 记录前k个元素的最大值
    res = append(res, queue.Front())

    for i := k; i < length; i++ {
        // 滑动窗口移除最前面的元素
        queue.Pop(nums[i-k])
        // 滑动窗口添加最后面的元素
        queue.Push(nums[i])
        // 记录最大值
        res = append(res, queue.Front())
    }
    return res
}