func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids)
    for _, ass := range asteroids {
        if mass < ass { return false }
        mass += ass
    }
    return true
}