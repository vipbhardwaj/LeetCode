type KthLargest struct {
	arr []int
	k   int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		arr: []int{},
		k:   k,
	}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if len(this.arr) < this.k {
		this.arr = append(this.arr, val)
		this.heapifyUp(len(this.arr) - 1)
	} else if val > this.arr[0] {
		this.arr[0] = val
		this.heapifyDown(0)
	}
	return this.arr[0]
}

// heapifyUp maintains the min-heap property when a new element is added
func (this *KthLargest) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if this.arr[parent] <= this.arr[index] {
			break
		}
		this.arr[parent], this.arr[index] = this.arr[index], this.arr[parent]
		index = parent
	}
}

// heapifyDown maintains the min-heap property after the root is replaced
func (this *KthLargest) heapifyDown(index int) {
	n := len(this.arr)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && this.arr[left] < this.arr[smallest] {
			smallest = left
		}
		if right < n && this.arr[right] < this.arr[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		this.arr[index], this.arr[smallest] = this.arr[smallest], this.arr[index]
		index = smallest
	}
}