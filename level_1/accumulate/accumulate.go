package accumulate


func Accumulate(t []string, f func(string) string) []string {
    res := make([]string, len(t))
    for i, s := range t {
        res[i] = f(s)
    }
    return res
}
