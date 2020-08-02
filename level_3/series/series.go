package series

func All(n int, s string) (res []string) {
    for i:=0; n <= len(s); i++ {
        res = append(res, s[i:n])
        n++
    }
    return
}

func UnsafeFirst(n int, s string) string {
    if n > len(s) {
        return ""
    }
    return s[:n]
}

func First(n int, s string) (string, bool) {
    if n > len(s) {
        return "", false
    }
    return s[:n], true
}
