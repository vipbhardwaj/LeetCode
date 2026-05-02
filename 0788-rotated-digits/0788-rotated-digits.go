func rotatedDigits(n int) int {
    count := 0
    for i := 1; i <= n; i++ {
        num := i
        isValid := true
        hasChange := false
        for num > 0 {
            digit := num % 10

            if digit == 3 || digit == 4 || digit == 7 {
                isValid = false
                break
            }
            if digit == 2 || digit == 5 || digit == 6 || digit == 9 {
                hasChange = true
            }
            num /= 10
        }
        if isValid && hasChange {
            count++
        }
    }
    return count
}