type ValueInterface interface {
	*ListValue
}

type ListValue struct {
	freq  int
	key   int
	value int
}

const MaxFreq = 200001

type LFUCache struct {
	storage map[int]*Item[*ListValue]
	lists   [MaxFreq]LinkedList[*ListValue]
	size    int
	cap     int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		storage: make(map[int]*Item[*ListValue], capacity),
		lists:   [MaxFreq]LinkedList[*ListValue]{},
		cap:     capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if i := this.get(key); i != nil {
		return i.value.value
	}

	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if i := this.get(key); i != nil {
		i.value.value = value
		return
	}

	if this.size == this.cap {
		minF := 1
		for ; this.lists[minF].Len() == 0; minF++ {
		}

		i := this.lists[minF].PopFront()
		delete(this.storage, i.value.key)

		this.size--
	}

	i := this.lists[1].PushBack(&ListValue{
		freq:  1,
		key:   key,
		value: value,
	})

	this.storage[key] = i
	this.size++
}

func (this *LFUCache) get(key int) *Item[*ListValue] {
	if i, ok := this.storage[key]; ok {
		this.touch(i)
		return i
	}

	return nil
}

func (this *LFUCache) touch(i *Item[*ListValue]) {
	this.lists[i.value.freq].Remove(i)
	val := i.value
	val.freq++

	newI := this.lists[val.freq].PushBack(val)
	this.storage[val.key] = newI
}

// Generic implementation Linked List below

type Item[TValue ValueInterface] struct {
	prev *Item[TValue]
	next *Item[TValue]

	value TValue
}

func NewItem[TValue ValueInterface](value TValue) *Item[TValue] {
	return &Item[TValue]{
		value: value,
	}
}

func (i *Item[TValue]) String() string {
	return fmt.Sprintf("%+v", i.value)
}

func (i *Item[TValue]) Remove() {
	if i.prev != nil {
		i.prev.next = i.next
	}

	if i.next != nil {
		i.next.prev = i.prev
	}

	i.prev = nil
	i.next = nil
}

func (i *Item[TValue]) InsertAfter(nextVal TValue) *Item[TValue] {
	next := NewItem[TValue](nextVal)

	if i.next != nil {
		i.next.prev = next
	}

	next.next = i.next
	next.prev = i
	i.next = next

	return next
}

func (i *Item[TValue]) InsertBefore(prevVal TValue) *Item[TValue] {
	prev := NewItem[TValue](prevVal)

	if i.prev != nil {
		i.prev.next = prev
	}

	prev.next = i
	prev.prev = i.prev
	i.prev = prev

	return prev
}

type LinkedList[TValue ValueInterface] struct {
	head *Item[TValue]
	tail *Item[TValue]
	size int
}

func (l *LinkedList[TValue]) InsertAfter(i *Item[TValue], nextVal TValue) *Item[TValue] {
	next := i.InsertAfter(nextVal)
	l.size++

	if i == l.tail {
		l.tail = next
	}

	return next
}

func (l *LinkedList[TValue]) InsertBefore(i *Item[TValue], prevVal TValue) *Item[TValue] {
	prev := i.InsertBefore(prevVal)
	l.size++

	if i == l.head {
		l.head = prev
	}

	return prev
}

func (l *LinkedList[TValue]) PushBack(val TValue) *Item[TValue] {
	if l.size == 0 {
		i := NewItem[TValue](val)
		l.head = i
		l.tail = i
		l.size++
		return i
	} else {
		return l.InsertAfter(l.tail, val)
	}
}

func (l *LinkedList[TValue]) PushFront(val TValue) *Item[TValue] {
	if l.size == 0 {
		i := NewItem[TValue](val)
		l.head = i
		l.tail = i
		l.size++
		return i
	} else {
		return l.InsertBefore(l.head, val)
	}
}

func (l *LinkedList[TValue]) PopBack() *Item[TValue] {
	i := l.tail
	l.Remove(i)
	return i
}

func (l *LinkedList[TValue]) PopFront() *Item[TValue] {
	i := l.head
	l.Remove(i)
	return i
}

func (l *LinkedList[TValue]) Remove(i *Item[TValue]) {
	newHead := l.head
	newTail := l.tail

	if l.head == i {
		newHead = l.head.next
	}

	if l.tail == i {
		newTail = l.tail.prev
	}

	i.Remove()
	l.size--

	l.head = newHead
	l.tail = newTail
}

func (l LinkedList[TValue]) Head() *Item[TValue] {
	return l.head
}

func (l LinkedList[TValue]) Tail() *Item[TValue] {
	return l.tail
}

func (l LinkedList[TValue]) Len() int {
	return l.size
}

func (l LinkedList[TValue]) String() string {
	items := ""
	for i := l.head; i != l.tail; i = i.next {
		items += fmt.Sprintf("%+v ", i)
	}

	items += fmt.Sprintf("%+v", l.tail)

	return fmt.Sprintf("{items: [%v], size: %d}", items, l.size)
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */