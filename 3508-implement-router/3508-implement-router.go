type Router struct {
	memLeft int
	packets [][]int
	seen map[[3]int]bool
	byDest map[int][]int
}

func Constructor(memoryLimit int) Router {
	return Router {
		memLeft: memoryLimit,
		packets: [][]int{},
		seen: make(map[[3]int]bool),
		byDest: make(map[int][]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
    key := [3]int{source, destination, timestamp}
	if this.seen[key] {
		return false
	}
	if this.memLeft == 0 {
		old := this.packets[0]
		this.packets = this.packets[1:]
		delete(this.seen, [3]int{old[0], old[1], old[2]})
		tlist := this.byDest[old[1]]
		idx := sort.SearchInts(tlist, old[2])
		if idx < len(tlist) && tlist[idx] == old[2] {
			this.byDest[old[1]] = append(tlist[:idx], tlist[idx+1:]...)
		}
	} else {
		this.memLeft--
	}
	this.packets = append(this.packets, []int{source, destination, timestamp})
	this.seen[key] = true
	tlist := this.byDest[destination]
	idx := sort.SearchInts(tlist, timestamp)
	tlist = append(tlist, 0)
	copy(tlist[idx + 1:], tlist[idx:])
	tlist[idx] = timestamp
	this.byDest[destination] = tlist

	return true
}

func (this *Router) ForwardPacket() []int {
    if len(this.packets) == 0 {
		return []int{}
	}
	packet := this.packets[0]
    this.packets = this.packets[1:]
    destination, timestamp := packet[1], packet[2]
    tlist := this.byDest[destination]
	idx := sort.SearchInts(tlist, timestamp)
    if idx < len(tlist) - 1 {
	    copy(tlist[idx:], tlist[idx + 1:])
    }
	tlist = tlist[:len(tlist) - 1]
	this.byDest[destination] = tlist
    delete(this.seen, [3]int{packet[0], destination, timestamp})
    this.memLeft++
	return packet
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    tlist := this.byDest[destination]
	if len(tlist) == 0 {
		return 0
	}
	left := sort.SearchInts(tlist, startTime)
	right := sort.Search(len(tlist), func(i int) bool {
		return tlist[i] > endTime
	})
    return right - left
}