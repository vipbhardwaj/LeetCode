type FoodRatings struct {
    foodToCuisine map[string]string
    foodToRating map[string]int
    cuisineToHeap map[string]*FoodHeap
}


func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    fr := FoodRatings{
        foodToCuisine: make(map[string]string),
        foodToRating:  make(map[string]int),
        cuisineToHeap: make(map[string]*FoodHeap),
    }

    for i := range foods {
        food := foods[i]
        cuisine := cuisines[i]
        rating := ratings[i]

        fr.foodToCuisine[food] = cuisine
        fr.foodToRating[food] = rating
        if _, ok := fr.cuisineToHeap[cuisine]; !ok {
            fr.cuisineToHeap[cuisine] = &FoodHeap{}
            heap.Init(fr.cuisineToHeap[cuisine])
        }
        heap.Push(fr.cuisineToHeap[cuisine], &Food{name: food, rating: rating})
    }
    return fr
}


func (fr *FoodRatings) ChangeRating(food string, newRating int) {
    fr.foodToRating[food] = newRating
    cuisine := fr.foodToCuisine[food]
    heap.Push(fr.cuisineToHeap[cuisine], &Food{name: food, rating: newRating})
}


func (fr *FoodRatings) HighestRated(cuisine string) string {
    pq := fr.cuisineToHeap[cuisine]
    for {
        top := (*pq)[0]
        if fr.foodToRating[top.name] == top.rating {
            return top.name
        }
        heap.Pop(pq)
    }
}

type Food struct {
    name   string
    rating int
    index  int // required for heap updates
}

type FoodHeap []*Food

func (fh FoodHeap) Len() int            { return len(fh) }
func (fh FoodHeap) Less(i, j int) bool  {
    if fh[i].rating == fh[j].rating {
        return fh[i].name < fh[j].name // lexicographical tie-breaker
    }
    return fh[i].rating > fh[j].rating
}
func (fh FoodHeap) Swap(i, j int) {
    fh[i], fh[j] = fh[j], fh[i]
    fh[i].index = i
    fh[j].index = j
}

func (fh *FoodHeap) Push(x interface{}) {
    item := x.(*Food)
    item.index = len(*fh)
    *fh = append(*fh, item)
}

func (fh *FoodHeap) Pop() interface{} {
    old := *fh
    n := len(old)
    item := old[n-1]
    *fh = old[0 : n-1]
    return item
}