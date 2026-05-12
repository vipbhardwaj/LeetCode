func minimumEffort(tasks [][]int) int {
    for i := range tasks {
        tasks[i] = append(tasks[i], tasks[i][1] - tasks[i][0])
    }
    sort.Slice(tasks, func (i, j int) bool {
        return tasks[i][2] > tasks[j][2]
    })

    var l, bal, loan int = len(tasks), tasks[0][2], 0
    for i:=1; i < l; i++ {
        if bal < tasks[i][1] {
            loan += tasks[i][1] - bal
            bal = tasks[i][1]
        }
        bal -= tasks[i][0]
    }
    return tasks[0][1] + loan
}