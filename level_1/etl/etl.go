package etl

func toLower(s string) string {
    return string(rune(s[0]) - 'A' + 'a')
}

func Transform(oldData map[int][]string) map[string]int {
    newData := make(map[string]int)
    for k, t := range oldData {
        for _, s := range t {
            newData[toLower(s)] = k
        }
    }
    return newData
}
