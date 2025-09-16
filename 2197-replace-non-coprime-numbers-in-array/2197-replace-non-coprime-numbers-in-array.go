func replaceNonCoprimes(nums []int) []int {
    result := []int{}
    for _, num := range nums {
        result = append(result, num)
        for len(result) > 1 {
            a := result[len(result)-1]
            b := result[len(result)-2]
            g := gcd(a, b)
            if g > 1 {
                result = result[:len(result)-2]
                lcm := a / g * b
                result = append(result, lcm)
            } else {
                break
            }
        }
    }
    return result
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}