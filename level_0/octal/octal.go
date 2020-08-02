package octal

import (
    "errors"
)

func ParseOctal(oct string) (int64, error) {
    var val int64 = 0

    for i := 0 ; i < len(oct); i++ {
        c := oct[i]
        var d int64 = 0
        switch {
        case '0' <= c && c <= '7':
            d = int64(c - '0')
        default:
            return 0, errors.New("syntax")
        }
        val *= 8
        val += d
    }
    return val, nil
}
