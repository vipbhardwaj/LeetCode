func intersect(nums1 []int, nums2 []int) []int {
    var res []int
    // Sort the two arrays.
    sort.Ints(nums1)
    sort.Ints(nums2)
    
    // Initiate the two pointers and place them at the common elements.
    for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
        if nums1[i] == nums2[j] {
            res = append(res, nums1[i])
            i++
            j++
        } else if nums1[i] < nums2[j] {
            i++
        } else if nums2[j] < nums1[i] {
            j++
        }
    }
    return res
}