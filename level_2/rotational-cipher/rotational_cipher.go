package rotationalcipher

import (
    "unicode"
)


func RotationalCipher(s string, key int) (string) {
    var res = make([]rune, len(s))
    for i, v := range s {
        if unicode.IsLower(v) {
            res[i] = rune(((int(v) + key - 'a') % 26) + 'a')
        } else if unicode.IsUpper(v) {
            res[i] = rune(((int(v) + key - 'A') % 26) + 'A')
        } else {
            res[i] = v
        }
    }
    return string(res)
}
