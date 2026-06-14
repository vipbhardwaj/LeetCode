func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
    var n int = len(wage)
    indv := make([][]float64, n)
    for i:=0; i<n; i++ {
        indv[i] = []float64{float64(wage[i]) / float64(quality[i]), float64(quality[i])}
    }
    sort.Slice(indv, func(i, j int) bool {
        return indv[i][0] < indv[j][0]
    })
    fmt.Println(indv)
    
    var res float64 = 1e18
    var curr int
    h := &MaxHeap{}
    for i := range indv {
        curr += int(indv[i][1])
        heap.Push(h, int(indv[i][1]))
        if h.Len() > k {
            curr -= heap.Pop(h).(int)
        }
        if h.Len() == k && res > indv[i][0] * float64(curr) {
            res = indv[i][0] * float64(curr)
        }
    }
    return res
}

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(int))
}
func(h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}