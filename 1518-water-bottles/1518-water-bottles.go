func numWaterBottles(numBottles int, numExchange int) int {
    res := numBottles
    rem := 0
    for numBottles >= numExchange {
        rem = numBottles % numExchange
        numBottles /= numExchange
        res += numBottles
        numBottles += rem
    }
    return res
}