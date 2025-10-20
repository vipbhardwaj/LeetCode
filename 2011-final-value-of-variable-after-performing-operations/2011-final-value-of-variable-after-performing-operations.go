func finalValueAfterOperations(operations []string) int {
    var x int
    for _, o := range operations {
        if o[0] == '+' || o[2] == '+' {
            x++
        } else if o[0] == '-' || o[2] == '-' {
            x--
        }
    }
    return x
}