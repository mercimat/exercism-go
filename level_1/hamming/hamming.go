package hamming

import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("strings of difference length")
    }

    d := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            d++
        }
    }
    return d, nil
}

