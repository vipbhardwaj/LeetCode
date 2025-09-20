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
		this.ForwardPacket()
	}
	this.packets = append(this.packets, []int{source, destination, timestamp})
	this.seen[key] = true
	this.byDest[destination] = append(this.byDest[destination], timestamp)
    this.memLeft--
	return true
}

func (this *Router) ForwardPacket() []int {
    if len(this.packets) == 0 {
		return []int{}
	}
	packet := this.packets[0]
    this.packets = this.packets[1:]
    delete(this.seen, [3]int{packet[0], packet[1], packet[2]})
    this.memLeft++

    times := this.byDest[packet[1]][1:]
    if len(times) == 0 {
        delete(this.byDest, packet[1])
    } else {
        this.byDest[packet[1]] = times
    }
	return packet
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    times, ok := this.byDest[destination]
	if !ok {
		return 0
	}
	left := sort.SearchInts(times, startTime)
    right := sort.Search(len(times), func(i int) bool { return times[i] > endTime })
    return right - left
}