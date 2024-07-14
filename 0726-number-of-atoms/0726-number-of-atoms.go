type Atom struct {
    e string
    c int
}

func countOfAtoms(formula string) string {
    var atoms []Atom
    var stack []int
    var element string
    
    for i:=0;  i<len(formula); i++ {
        r := rune(formula[i])
        if r == '(' {
            stack = append(stack, len(atoms))
        } else if unicode.IsDigit(r) {
            mul := false
            if formula[i-1] == ')' {
                mul = true
            }
            num := int(r-'0')
            for (i+1 < len(formula) && formula[i+1] >= '0' && formula[i+1] <= '9') {
                r = rune(formula[i+1])
                num = num*10 + int(r-'0')
                i++
            }
            if mul {
                past := stack[len(stack)-1]
                // fmt.Println(past)
                stack = stack[:len(stack)-1]
                for j:=past; j<len(atoms); j++ {
                    // fmt.Println(atoms[j].e)
                    atoms[j].c *= num
                }
                // fmt.Println("_________")
            } else {
                atoms[len(atoms)-1].c *= num
            }
        }
        if r >= 'A' && r <= 'Z' {
            element = string(r)
            for (i+1 < len(formula) && formula[i+1] >= 'a' && formula[i+1] <= 'z') {
                r = rune(formula[i+1])
                element += string(r)
                i++
            }
            // fmt.Println(element)
            atoms = append(atoms, Atom{element, 1})
        }
    }
    // for _, j := range atoms {
    //     fmt.Println("e:", j.e, ", c:", j.c)
    // }
    sort.Slice(atoms, func(i, j int) bool {
        return atoms[i].e < atoms[j].e
    })
    m := make(map[string]int)
    for _, a := range atoms {
        _, ok := m[a.e]
        if ok {
            m[a.e] += a.c
        } else {
            m[a.e] = a.c
        }
    }
    atoms = []Atom{}
    for k, v := range m {
        atoms = append(atoms, Atom{k, v})
    }
    sort.Slice(atoms, func(i, j int) bool {
        return atoms[i].e < atoms[j].e
    })
    res := ""
    for _, a := range atoms {
        res += string(a.e)
        if a.c > 1 {
            res += strconv.Itoa(a.c)
        }
    }
    return res
}