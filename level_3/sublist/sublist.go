package sublist

type Relation string

func isSubListOf(a, b []int) bool {
    if len(a) > len(b) {
        a, b = b, a
    }
    j := 0
    if len(a) == 0 {
        return true
    }

    for {
        if j >= len(b) - len(a) + 1 {
            break
        }
        if a[0] == b[j] {
            i := 0
            for  {
                if i >= len(a) {
                    return true
                }
                if a[i] != b[j+i] {
                    break
                }
                i++
            }
        }
        j++
    }
    return false
}

func Sublist(a, b []int) Relation {
    if !isSubListOf(a, b) {
        return "unequal"
    }
    if len(a) < len(b) {
        return "sublist"
    } else if len(a) > len(b) {
        return "superlist"
    } else {
        return "equal"
    }
}
