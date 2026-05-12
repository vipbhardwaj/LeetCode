func minimumEffort(tasks [][]int) int {
    sort.Slice(tasks, func (i, j int) bool {
        return tasks[i][1] - tasks[i][0] > tasks[j][1] - tasks[j][0]
    })

    var l, bal, loan int = len(tasks), tasks[0][1]-tasks[0][0], 0
    for i:=1; i < l; i++ {
        var spend int = tasks[i][1]
        if bal < spend {
            loan += spend - bal
            bal = spend
        }
        bal -= tasks[i][0]
    }
    return tasks[0][1] + loan
}