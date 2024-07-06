func passThePillow(n int, time int) int {
    side := true
    for time >= n {
        time -= (n-1)
        side = !side
    }
    // fmt.Println(side)
    if side {
        return time+1
    } else {
        return n-time
    }
    
}