package binary

import "fmt"

func ParseBinary(s string) (int, error) {
    b := 0
    pow := 1
    for i := 0; i < len(s); i++ {
        c := s[len(s)-1-i]
        if c == '0' {
            pow *= 2
            continue
        } else if c == '1' {
            b += pow
            pow *= 2
        } else {
            return 0, fmt.Errorf("Unexpected char %v", c)
        }
    }
    return b, nil
}
