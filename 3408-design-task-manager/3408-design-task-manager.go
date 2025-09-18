type TaskManager struct {
    taskToPriority map[int]int       // taskId → priority
    taskToUser     map[int]int       // taskId → userId
    deleted        map[int]bool      // taskId → true if removed
    globalHeap     *TaskHeap         // max-heap of all tasks
}

func Constructor(tasks [][]int) TaskManager {
    tm := TaskManager{
        taskToPriority: make(map[int]int),
        taskToUser:     make(map[int]int),
        deleted:        make(map[int]bool),
        globalHeap:     &TaskHeap{},
    }
    heap.Init(tm.globalHeap)
    for _, t := range tasks {
        userId, taskId, priority := t[0], t[1], t[2]
        tm.taskToPriority[taskId] = priority
        tm.taskToUser[taskId] = userId
        heap.Push(tm.globalHeap, &Task{taskId: taskId, priority: priority})
    }
    return tm
}

func (tm *TaskManager) Add(userId int, taskId int, priority int)  {
    tm.taskToPriority[taskId] = priority
    tm.taskToUser[taskId] = userId
    tm.deleted[taskId] = false
    heap.Push(tm.globalHeap, &Task{taskId: taskId, priority: priority})
}

func (tm *TaskManager) Edit(taskId int, newPriority int)  {
    tm.taskToPriority[taskId] = newPriority
    heap.Push(tm.globalHeap, &Task{taskId: taskId, priority: newPriority})
}

func (tm *TaskManager) Rmv(taskId int)  {
    tm.deleted[taskId] = true
    delete(tm.taskToPriority, taskId)
    delete(tm.taskToUser, taskId)
}

func (tm *TaskManager) ExecTop() int {
    for tm.globalHeap.Len() > 0 {
        top := (*tm.globalHeap)[0]
        actualPriority, exists := tm.taskToPriority[top.taskId]
        if !exists || actualPriority != top.priority || tm.deleted[top.taskId] {
            heap.Pop(tm.globalHeap)
            continue
        }
        userId := tm.taskToUser[top.taskId]
        tm.Rmv(top.taskId)
        return userId
    }
    return -1
}

type Task struct {
    taskId   int
    priority int
}
type TaskHeap []*Task

func (th TaskHeap) Len() int { return len(th) }
func (th TaskHeap) Less(i, j int) bool {
    if th[i].priority == th[j].priority {
        return th[i].taskId > th[j].taskId
    }
    return th[i].priority > th[j].priority
}
func (th TaskHeap) Swap(i, j int)       { th[i], th[j] = th[j], th[i] }
func (th *TaskHeap) Push(x interface{}) { *th = append(*th, x.(*Task)) }
func (th *TaskHeap) Pop() interface{} {
    old := *th
    n := len(old)
    item := old[n-1]
    *th = old[:n-1]
    return item
}