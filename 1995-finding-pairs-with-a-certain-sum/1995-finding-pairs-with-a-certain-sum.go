type FindSumPairs struct {
    nums1 []int
    nums2 []int
    cache map[int]int
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
    f := FindSumPairs{
        nums1: nums1,
        nums2: nums2,
        cache: make(map[int]int),
    }
    for _, v := range f.nums2 {
        f.cache[v]++
    }
    return f
}


func (this *FindSumPairs) Add(i int, val int)  {
    this.cache[this.nums2[i]]--
    this.nums2[i] += val
    this.cache[this.nums2[i]]++
}


func (this *FindSumPairs) Count(tot int) int {
    var res int
    for _, i := range this.nums1 {
        var comp int = tot - i
        res += this.cache[comp]
    }
    return res
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */