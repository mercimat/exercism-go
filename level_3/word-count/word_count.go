package wordcount

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    var word = make([]byte, 0)
    var freqs = make(Frequency)
    for i:=0; i<len(phrase); i++ {
        c := normalizeChar(phrase[i])
        if c == ' ' {
            if len(word) > 0 {
                if word[0] == '\'' && word[len(word)-1] == '\'' {
                    freqs[string(word[1:len(word)-1])] += 1
                } else {
                    freqs[string(word)] += 1
                }
                word = make([]byte, 0)
            }
            continue
        }
        word = append(word, c)
    }
    if len(word) > 0 {
        if word[0] == '\'' && word[len(word)-1] == '\'' {
            freqs[string(word[1:len(word)-1])] += 1
        } else {
            freqs[string(word)] += 1
        }
    }
    return freqs
}

func normalizeChar(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return byte(c-'A'+'a')
    }
    if c >= 'a' && c <= 'z' {
        return c
    }
    if c >= '0' && c <= '9' {
        return c
    }
    if c == '\'' {
        return c
    }
    return ' '
}
