type Task struct {
	count     int
	coolDownUntil int
}

// 1. Correctly implemented TaskHeap
type TaskHeap []int // We only need to store frequencies/counts in the max-heap

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-Heap: highest frequency first
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	// Count frequencies
	frequencies := make(map[byte]int)
	for _, t := range tasks {
		frequencies[t]++
	}

	// 2. Properly instantiate and initialize the heap
	h := &TaskHeap{}
	heap.Init(h)
	for _, count := range frequencies {
		heap.Push(h, count)
	}

	type cooldownTask struct {
		count     int
		available int
	}
	
	// Queue to hold tasks during their cooling period
	var queue []cooldownTask
	time := 0

	for h.Len() > 0 || len(queue) > 0 {
		time++

		if h.Len() > 0 {
			// Process the most frequent available task
			count := heap.Pop(h).(int) - 1
			if count > 0 {
				// 3. Instead of pushing instantly, put it in a cooling queue
				queue = append(queue, cooldownTask{count: count, available: time + n})
			}
		}

		// Check if the task at the front of the queue has finished cooling down
		if len(queue) > 0 && queue[0].available == time {
			heap.Push(h, queue[0].count)
			queue = queue[1:]
		}
	}

	return time
}