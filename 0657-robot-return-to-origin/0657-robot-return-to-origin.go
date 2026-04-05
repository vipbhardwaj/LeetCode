func judgeCircle(moves string) bool {
    var x, y int
    for _, move := range moves {
        switch move {
            case 'U':
                x++
            case 'D':
                x--
            case 'L':
                y++
            case 'R':
                y--
            default:
                continue
            }
    }
    return x == 0 && y == 0
}