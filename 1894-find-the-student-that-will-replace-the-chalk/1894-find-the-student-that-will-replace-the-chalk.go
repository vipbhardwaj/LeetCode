func chalkReplacer(chalk []int, k int) int {
    chalkSum := 0
    for _, i := range chalk {
        chalkSum += i
    }
    
    k %= chalkSum
    i := 0
    for k >= chalk[i] {
        k -= chalk[i]
        i++
    }
    return i
}