func totalMoney(n int) int {
    var money, week int
    for n > 7 {
        money += (7 * (7+1)) / 2 + 7 * week
        week++
        n -= 7
    }
    money += (n * (n+1)) / 2 + n * week
    return money
}