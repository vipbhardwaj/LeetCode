func averageWaitingTime(customers [][]int) float64 {
    prevTime := 0
    totalWait := 0
    for _, p := range customers {
        wait := 0
        if prevTime < p[0] {
            wait = p[1]
            prevTime = p[1] + p[0]
        } else {
            prevTime += p[1]
            wait = prevTime - p[0]
        }

        // fmt.Println(currTime)
        // fmt.Println(wait)
        totalWait += wait
        // fmt.Println(totalWait)
        // fmt.Println()
    }
    return float64(totalWait) / float64(len(customers))
}