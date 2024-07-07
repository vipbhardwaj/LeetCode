func chalkReplacer(chalk []int, k int) int {
    chalkSum := 0
    for _, i := range chalk {
        chalkSum += i
    }
    
    rem := k % chalkSum
    // times := chalkSum / k
    // fmt.Println(rem)
    for i, n := range chalk {
        if rem < n {
            return i
        }
        rem -= n
    }
    return len(chalk)-1
}