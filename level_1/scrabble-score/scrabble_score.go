package scrabble

var letterValues = map[rune]int{
    'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1,
    'f': 4, 'g': 2, 'h': 4, 'i': 1, 'j': 8,
    'k': 5, 'l': 1, 'm': 3, 'n': 1, 'o': 1,
    'p': 3, 'q': 10, 'r': 1, 's': 1, 't': 1,
    'u': 1, 'v': 4, 'w': 4, 'x': 8, 'y': 4,
    'z': 10,
}

func Score(w string) int {
    score := 0
    for i := 0; i < len(w); i++ {
        r := rune(w[i])
        if 'A' <= r && r <= 'Z' {
            r = 'a' + r - 'A'
        }
        score += letterValues[r]
    }
    return score
}
