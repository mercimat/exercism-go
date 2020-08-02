package anagram

func Detect(subject string, candidates []string) []string {
    var res = make([]string, 0)
    for _, candidate := range candidates {
        if isAnagram(subject, candidate) {
            res = append(res, candidate)
        }
    }
    return res
}

func isAnagram(subject string, candidate string) bool {
    if len(subject) != len(candidate) {
        return false
    }
    if equalWord(subject, candidate) {
        return false
    }
    var letters = make(map[byte]int)
    for i:=0; i<len(subject); i++ {
        r := normalizeRune(subject[i])
        letters[r] += 1
    }
    for i:=0; i<len(candidate); i++ {
        r := normalizeRune(candidate[i])
        letters[r] -= 1
        if letters[r] < 0 {
            return false
        }
    }
    return true
}

func equalWord(w1, w2 string) bool {
    if len(w1) != len(w2) {
        return false
    }
    for i:=0; i<len(w1); i++ {
        if normalizeRune(w1[i]) != normalizeRune(w2[i]) {
            return false
        }
    }
    return true
}

func normalizeRune(r byte) byte {
    if r >= 'A' && r <= 'Z' {
        return byte(r-'A'+'a')
    }
    return r
}
