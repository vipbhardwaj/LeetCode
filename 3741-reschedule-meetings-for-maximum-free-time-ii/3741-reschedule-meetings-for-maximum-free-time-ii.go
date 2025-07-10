func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    n := len(startTime)

    gaps := make([]int, n+1)
    gaps[0] = startTime[0]
    for i := 1; i < n; i++ {
        gaps[i] = startTime[i] - endTime[i-1]
    }
    gaps[n] = eventTime - endTime[n-1]

    maxRight := make([]int, n+1)
    for i := n - 1; i >= 0; i-- {
        maxRight[i] = max(maxRight[i+1], gaps[i+1])
    }

    maxLeft, res := 0, 0
    for i := 1; i <= n; i++ {
        duration := endTime[i-1] - startTime[i-1]
        if maxLeft >= duration || maxRight[i] >= duration {
            res = max(res, gaps[i-1] + duration + gaps[i])
        }
        res = max(res, gaps[i-1] + gaps[i])
        maxLeft = max(maxLeft, gaps[i-1])
    }
    return res
}