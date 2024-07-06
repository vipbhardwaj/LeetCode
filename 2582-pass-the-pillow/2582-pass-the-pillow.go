func passThePillow(n int, time int) int {
    return O1(n, time)
    return On(n, time)
}

func O1(n int, time int) int {
    side := time / (n-1)
    time %= (n-1)
    
    if side % 2 == 0 {
        return time+1
    } else {
        return n-time
    }
}

func On(n int, time int) int {
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