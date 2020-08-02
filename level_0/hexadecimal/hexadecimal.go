package hexadecimal

import (
    "fmt"
    "math"
)

type SyntaxError string
type RangeError string

func (e *SyntaxError) Error() string {
    return fmt.Sprintf("syntax")
}

func (e *RangeError) Error() string {
    return fmt.Sprintf("range")
}

func ParseHex(hex string) (int64, error) {
    if len(hex) < 1 {
        err := SyntaxError("")
        return 0, &err
    }
    var val int64 = 0
    var val1 int64 = val
    var pow int64 = 1
    for i := 0; i < len(hex); i++ {
        c := hex[len(hex)-1 - i]
        var d byte = 0
        switch {
        case '0' <= c && c <= '9':
            d = c - '0'
        case 'a' <= c && c <= 'f':
            d = 10 + c - 'a'
        case 'A' <= c && c <= 'F':
            d = 10 + c - 'A'
        default:
            err := SyntaxError("")
            return 0, &err
        }
        if val >= math.MaxInt64/16 + 1 {
            err := RangeError("")
            return math.MaxInt64, &err
        }
        val1 = val
        val1 += pow * int64(d)
        if val1 < val {
            err := RangeError("")
            return math.MaxInt64, &err
        }
        val = val1
        pow *= 16
    }
    return val, nil
}

func HandleErrors(hexStrings []string) (errors []string) {
    for _, hex := range hexStrings {
        _, err := ParseHex(hex)
        if err != nil {
            errors = append(errors, err.Error())
        } else {
            errors = append(errors, "none")
        }
    }
    return
}
