package trinary

import (
    "errors"
)

func ParseTrinary(oct string) (int64, error) {
    var val int64 = 0
    var val1 int64 = 0

    for i := 0 ; i < len(oct); i++ {
        c := oct[i]
        var d int64 = 0
        switch {
        case '0' <= c && c <= '2':
            d = int64(c - '0')
        default:
            return 0, errors.New("syntax")
        }
        val1 = val
        val1 *= 3
        val1 += d
        if val1 < val {
            return 0, errors.New("range")
        }
        val = val1
    }
    return val, nil
}
