type LFUCache struct {
    m map[int]int
    cap int
    h MinHeap
    time int
    n map[int]*CacheItem
}


func Constructor(capacity int) LFUCache {
    return LFUCache {
        m: make(map[int]int),
        cap: capacity,
        h: make(MinHeap, 0),
        time: 0,
        n: make(map[int]*CacheItem),
    }
}


func (this *LFUCache) Get(key int) int {
    this.time++
    if v, ok := this.m[key]; ok {
        node := this.n[key]
        node.count++
        node.timestamp = this.time
        heap.Init(&this.h)
        return v
    }
    return -1
}


func (this *LFUCache) Put(key int, value int)  {
    if this.cap == 0 {
		return
	}

	this.time++

	// Scenario A: Key already exists, update value and frequency
	if _, exists := this.m[key]; exists {
		this.m[key] = value
		item := this.n[key]
		item.count++
		item.timestamp = this.time
		heap.Init(&this.h)
		return
	}

	// Scenario B: Cache is full, evict the top item from the Min-Heap
	if len(this.m) >= this.cap {
		evictedItem := heap.Pop(&this.h).(*CacheItem)
		delete(this.m, evictedItem.key)
		delete(this.n, evictedItem.key)
	}

	// Scenario C: Insert the completely new item
	newItem := &CacheItem{
		key:       key,
		count:     1,
		timestamp: this.time,
	}
	this.m[key] = value
	this.n[key] = newItem
	heap.Push(&this.h, newItem)
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
 
type CacheItem struct {
    key int
    count int
    timestamp int
}

type MinHeap []*CacheItem
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].timestamp < h[j].timestamp
	}
	return h[i].count < h[j].count
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*CacheItem))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return item
}

