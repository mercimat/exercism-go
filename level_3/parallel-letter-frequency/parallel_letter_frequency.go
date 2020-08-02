package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
    m := FreqMap{}
    for _, r := range s {
        m[r]++
    }
    return m
}

func ConcurrentFrequency(texts []string) FreqMap {
    res := make(FreqMap)
    ch := make(chan FreqMap, len(texts))
    for i := 0; i < len(texts); i++ {
        s := texts[i]
        go func() {
            m := make(FreqMap)
            m = Frequency(s)
            ch <- m
        }()
    }
    for i := 0; i < len(texts); i++ {
        m := <-ch
        for r, v := range m {
            if _, ok := res[r]; ok {
                res[r] += v
            } else {
                res[r] = v
            }
        }
    }
    return res
}
