func countPrimeSetBits(left int, right int) int {
    var res int
    for i:= left; i <= right; i++ {
        if isPrime(setBits(i)) {
            res++
        }
    }
    return res
}

func setBits(i int) int {
    var set int
    for i > 0 {
        if i & 1 == 1 {
            set++
        }
        i >>= 1
    }
    return set
}

func isPrime(n int) bool {
    // Check if n is 1 or 0
    if n <= 1 {
        return false
    }

    // Check if n is 2 or 3
    if n == 2 || n == 3 {
        return true
    }

    // Check whether n is divisible by 2 or 3
    if n % 2 == 0 || n % 3 == 0 {
        return false
    }
    
    // Check from 5 to square root of n
    // Iterate i by (i+6)
    for i := 5; i*i <= n ; i += 6 {
        if n % i == 0 || n % (i + 2) == 0 {
            return false
        }
    }

    return true
}