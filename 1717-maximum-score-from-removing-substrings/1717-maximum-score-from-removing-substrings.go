type stack[T any] struct {
    Push func(T)
    Pop func() T
    Length func() int
    Top func() T
    Iter func() []T
}

func Stack[T any]() stack[T] {
    slice := make([]T, 0)
    return stack[T] {
        Push: func(i T) {
            slice = append(slice, i)
        },
        Pop: func() T {
            var t T
            if len(slice) == 0 {
                return t
            }
            t = slice[len(slice)-1]
            slice = slice[:len(slice)-1]
            return t
        },
        Length: func() int {
            return len(slice)
        },
        Top: func() T {
            if len(slice) == 0 {
                var t T
                return t
            }
            return slice[len(slice)-1]
        },
        Iter: func() []T {
            var elements []T
            for _, e := range slice {
                elements = append(elements, e)
            }
            return elements
        },
    }
}

func stackApproach(s string, x int, y int) int {
    st := stack[rune]{}
    var res int
    
    for _, c := range s {
        fmt.Println(string(c))
        if st.Length() == 0 {
            st.Push(c)
        } else if c == 'b' && st.Top() == 'a' {
            res += x
            st.Pop()
        } else {
            st.Push(c)
        }
    }
    
    for _, r := range st.Iter() {
        if r == 'a' && st.Top() == 'b' {
            res += y
            st.Pop()
        }
    }
    return res
}

func maximumGain(s string, x int, y int) int {
    return counterMethod(s, x, y)
    return stackApproach(s, x, y)
}

func reverse(str string) string {
    s := []rune(str)
    for i := 0; i < len(s) / 2; i++ {
        s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
    }
    return string(s)
}

func counterMethod(s string, x int, y int) int {
    if (x < y) {
        temp := y
        y = x
        x = temp
        s = reverse(s)
    }
    
    aCount := 0
    bCount := 0
    res := 0
    for _, c := range s {
        if c == 'a' {
            aCount++
        } else if c == 'b' {
            if aCount > 0 {
                res += x
                aCount--
            } else {
                bCount++
            }
        } else {
            res += min(aCount, bCount) * y
            aCount = 0
            bCount = 0
        }
    }
    return res + min(aCount, bCount) * y
    
}