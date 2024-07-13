type Specs struct {
    pos int
    hp int
    dir bool
    ded bool
    i int
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    var robots []Specs
    for i := 0; i < len(healths); i++ {
        var dir bool
        if directions[i] == 'R' {
            dir = true
        } else {
            dir = false
        }
        r := Specs {
            pos: positions[i],
            hp: healths[i],
            dir: dir,
            ded: false,
            i: i,
        }
        robots = append(robots, r)
    }
    sort.Slice(robots, func(i, j int) bool {
        return robots[i].pos < robots[j].pos
    })
    // for _, r := range robots {
    //     fmt.Println("hp:", r.hp, "dir:", r.dir)
    // }
    // fmt.Println("=================")
    // return []int{}
    
    var stack []*Specs
    for i := range robots {
        r := &robots[i]
        // fmt.Println("hp:", r.hp, "dir:", r.dir)
        // fmt.Println(len(stack))
        if r.dir {
            stack = append(stack, r)
        } else {
            for len(stack) > 0 {
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if r.hp > top.hp {
                    r.hp--
                    top.ded = true
                } else if r.hp == top.hp {
                    r.ded = true
                    top.ded = true
                    break
                } else {
                    top.hp--
                    r.ded = true
                    if top.hp > 0 {
                        stack = append(stack, top)
                    } else {
                        top.ded = true
                    }
                    break
                }
            }
        }
    }
    var res []int
    sort.Slice(robots, func(i, j int) bool {
        return robots[i].i < robots[j].i
    })
    // fmt.Println("=================")
    for _, r := range robots {
        // fmt.Println("hp:", r.hp, "ded:", r.ded)
        if r.ded {
            continue
        }
        res = append(res, r.hp)
    }
    return res
}