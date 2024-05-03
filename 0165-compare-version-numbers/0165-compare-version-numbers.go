func compareVersion(version1 string, version2 string) int {
    revisions1 := strings.Split(version1, ".")
    revisions2 := strings.Split(version2, ".")

    batch1 := process(revisions1)
    batch2 := process(revisions2)
    // fmt.Println("batch1 len,", len(batch1))
    // fmt.Println("batch2 len,", len(batch2))

//     if len(revisions1) > len(revisions2) {
//         return 1
//     } else if len(revisions2) > len(revisions1) {
//         return -1
//     }

    batch1 = trailRemover(batch1)
    batch2 = trailRemover(batch2)

    return compareBatch(batch1, batch2)

    minLen := int(math.Min(float64(len(revisions1)), float64(len(revisions2))))
    for i:=0; i < minLen; i++ {
        rev1 := revisions1[i]
        rev2 := revisions2[i]
        comparison := compare(rev1, rev2)
        if comparison != 0 {
            return comparison
        } else {
            continue
        }
    }

    maxLen := int(math.Max(float64(len(revisions1)), float64(len(revisions2))))
    for i := minLen; i < maxLen; i++ {
        var comparison int
        if minLen == len(revisions1) {
            comparison = compare("", revisions2[i])
        } else {
            comparison = compare(revisions1[i], "")
        }
        if comparison == 0 {
            continue
        } else {
            return comparison
        }
    }
    return 0
}

func trailRemover(batch []int) []int {
    if len(batch) == 0 {
        return []int{0}
    }
    if batch[len(batch)-1] == 0 {
        return trailRemover(batch[:len(batch)-1])
    } else {
        return batch
    }
}

func compareBatch(rev1 []int, rev2 []int) int {
    for i, j := 0, 0; i < len(rev1) && j < len(rev2); i,j = i+1, j+1 {
        if rev1[i] > rev2[j] {
            return 1
        } else if rev1[i] < rev2[j] {
            return -1
        } else {
            continue
        }
    }

    var rev []int
    var l int
    var s int
    if len(rev1) > len(rev2) {
        rev = rev1
        l = len(rev1)
        s = len(rev1) - len(rev2)
    } else if len(rev2) > len(rev1) {
        rev = rev2
        l = len(rev2)
        s = len(rev2) - len(rev1)
    } else {
        return 0
    }

    // fmt.Println(len(rev1))
    // fmt.Println(len(rev2))
    for i:=s; i<l; i++ {
        // fmt.Println(rev[i])
        if rev[i] == 0 {
            continue
        } else {
            break
        }
    }
    if len(rev1) > len(rev2) {
        return 1
    } else if len(rev2) > len(rev1) {
        return -1
    } else {
        return 0
    }
}

func process(revs []string) []int {
    var processed []int
    for _, rev := range revs {
        flag := true
        batch := 0
        for _, c := range rev {
            if flag && c == '0' {
                continue
            }
            flag = false
            batch += 10*batch + int(c-'0')
        }
        // fmt.Print(batch, ",")
        if batch > 0 {
            processed = append(processed, batch)
        } else {
            processed = append(processed, 0)
        }
    }
    // fmt.Println()
    return processed
}

func compare(rev1 string, rev2 string) int {
    minLen := int(math.Min(float64(len(rev1)), float64(len(rev2))))
    maxLen := int(math.Max(float64(len(rev1)), float64(len(rev2))))
    var rev string

    if maxLen == len(rev1) {
        rev = rev1
    } else {
        rev = rev2
    }

    for i, j := 0, 0; i<len(rev1) && j < len(rev2); i, j = i+1, j+1 {
        // for ; rev1[i] == '0' && i < len(rev1); i++ {}
        // for ; rev2[j] == '0' && j < len(rev2); j++ {}

        if i < len(rev1) || j < len(rev2) {
            if j == len(rev2) {
                return -1
            } else if i == len(rev1) {
                return 1
            } else if rev1[i] > rev2[j] {
                return 1
            } else if rev2[j] > rev1[i] {
                return -1
            } else if rev1[i] == rev2[j] {
                continue
            }   
        } else {
            return 0
        }
    }

    // if i < len(rev1) && j < len(rev2) {
    //     if i == len(rev1) {
    //         return -1
    //     } else if j == len(rev2) {
    //         return 1
    //     } else {
    //         for 
    //     }
    // } else {
    //     return 0
    // }

    for i := minLen; i < maxLen; i++ {
        // if rev[i] == '0' {
        //     continue
        // }
        if rev == rev1 {
            return 1
        } else {
            return -1
        }
    }
    return 0
}