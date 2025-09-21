type MovieRentingSystem struct {
    entries  [][]int
    rented   map[[2]int]bool
    movieMap map[int][][]int
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
    sort.Slice(entries, func(i, j int) bool {
        if entries[i][2] == entries[j][2] {
            if entries[i][0] == entries[j][0] {
                return entries[i][1] < entries[j][1]
            }
            return entries[i][0] < entries[j][0]
        }
        return entries[i][2] < entries[j][2]
    })
    m := make(map[[2]int]bool)
    movieMap := make(map[int][][]int)
    for _, e := range entries {
        m[[2]int{e[0], e[1]}] = false
        movieMap[e[1]] = append(movieMap[e[1]], []int{e[0], e[2]})
    }
    return MovieRentingSystem{
        entries:  entries,
        rented:   m,
        movieMap: movieMap,
    }
}

func (this *MovieRentingSystem) Search(movie int) []int {
    var res []int
    seen := make(map[int]bool)
    for _, rec := range this.movieMap[movie] {
        shop := rec[0]
        if seen[shop] {
            continue
        }
        if val, ok := this.rented[[2]int{shop, movie}]; ok && !val {
            res = append(res, shop)
            seen[shop] = true
            if len(res) == 5 {
                break
            }
        }
    }
    return res
}

func (this *MovieRentingSystem) Rent(shop int, movie int)  {
    this.rented[[2]int{shop, movie}] = true
}

func (this *MovieRentingSystem) Drop(shop int, movie int)  {
    this.rented[[2]int{shop, movie}] = false
}

func (this *MovieRentingSystem) Report() [][]int {
    var res [][]int
    m := make(map[[2]int]bool)
    for _, e := range this.entries {
        if val, ok := this.rented[[2]int{e[0], e[1]}]; ok {
            if val {
                if _, ok := m[[2]int{e[0], e[1]}]; !ok {
                    res = append(res, []int{e[0], e[1]})
                }
                m[[2]int{e[0], e[1]}] = true
            }
        }
        if len(res) == 5 {
            break
        }
    }
    return res
}

// type Entry struct {
//     shop int
//     movie int
//     price int
//     index int
// }
// type PriorityQueue []*Entry
// func (pq PriorityQueue) Len() int { return len(pq) }
// func (pq PriorityQueue) Less(i, j int) bool {
//     if pq[i].price == pq[j].price {
//         if pq[i].shop == pq[j].shop {
//             return pq[i].movie < pq[j].movie
//         }
//         return pq[i].shop < pq[j].shop
//     }
//     return pq[i].price < pq[j].price
// }
// func (pq PriorityQueue) Swap(i, j int) {
//     pq[i], pq[j] = pq[j], pq[i]
//     pq[i].index = i
//     pq[j].index = j
// }
// func (pq *PriorityQueue) Push(x any) {
//     n := len(*pq)
//     item := x.(*Item)
//     item.index = n
//     *pq = append(*pq, item)
// }
// func (pq *PriorityQueue) Pop() any {
//     old := *pq
//     n := len(old)
//     item := old[n-1]
//     old[n-1] = nil // Avoid memory leak
//     item.index = -1 // For safety
//     *pq = old[0 : n-1]
//     return item
// }