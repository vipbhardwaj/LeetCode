func compareVersion(version1 string, version2 string) int {
    revisions1 := strings.Split(version1, ".")
    revisions2 := strings.Split(version2, ".")

    batch1 := process(revisions1)
    batch2 := process(revisions2)

    batch1 = trailRemover(batch1)
    batch2 = trailRemover(batch2)

    return compareBatch(batch1, batch2)
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

    for i:=s; i<l; i++ {
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
        if batch > 0 {
            processed = append(processed, batch)
        } else {
            processed = append(processed, 0)
        }
    }
    return processed
}