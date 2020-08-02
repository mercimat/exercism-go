package cryptosquare

import (
    "math"
    "strings"
)

func Encode(s string) string {
    s = strings.Map(norm, s)
    c := int(math.Ceil(math.Sqrt(float64(len(s)))))
    r := c-1
    if c*r < len(s) {
        r = c
    }
    res := make([]byte, 0, c*r)
    for i := 0; i < c; i++ {
        for j:=0; j < r; j++ {
            if c*j + i >= len(s) {
                res = append(res, ' ')
            } else {
                res = append(res, byte(s[c*j + i]))
            }
        }
        if i < c-1 {
            res = append(res, ' ')
        }
    }
    return string(res)
}

func norm(r rune) rune {
    switch {
    case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
        return r
    case r >= 'A' && r <= 'Z':
        return r + 'a' - 'A'
    }
    return -1
}

