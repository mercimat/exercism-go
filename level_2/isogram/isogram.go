package isogram

func IsIsogram(word string) bool {
    var byteArray [256]uint8

    for _, c := range word {
        if 'A' <= c && c <= 'Z' {
            c = 'a' + c - 'A'
        }
        if 'a' <= c && c <= 'z' {
            if byteArray[c] > 0 {
                return false
            }
            byteArray[c] = 1
        }
    }

    return true
}
