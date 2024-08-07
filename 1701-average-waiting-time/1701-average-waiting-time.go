func averageWaitingTime(customers [][]int) float64 {
    prevTime := 0
    totalWait := 0
    for _, p := range customers {
        if prevTime < p[0] {
            totalWait += p[1]
            prevTime = p[1] + p[0]
        } else {
            prevTime += p[1]
            totalWait += prevTime - p[0]
        }
    }
    return float64(totalWait) / float64(len(customers))
}