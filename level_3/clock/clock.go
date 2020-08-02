package clock

import "fmt"

type Clock int

func New(h, m int) Clock {
    c := Clock((h*60 + m) % 1440)
    if c < 0 {
        c += 1440
    }
    return c
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d", c/60, c%60)
}


func (c Clock) Add(minutes int) Clock {
    return New(0, int(c)+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
    return New(0, int(c)-minutes)
}
