package counter

import (
    "testing"
)

var testCases = []struct {
    s       string
    lines   int
    letters int
    chars   int
}{
    {"", 0, 0, 0},
    {"\n", 1, 0, 1},
    {"àb!", 1, 2, 3},
    {"abc\ndef", 2, 6, 7},
    {"\xbd\xb2=\xbc \u2318", 1, 0, 6},
}

func TestNoAdd(t *testing.T) {
    c := makeCounter()
    if l := c.Lines(); l != 0 {
        t.Fatalf("Wrong number of lines: expected %d but got %d", 0, l)
    }
    if l := c.Letters(); l != 0 {
        t.Fatalf("Wrong number of letters: expected %d but got %d", 0, l)
    }
    if l := c.Characters(); l != 0 {
        t.Fatalf("Wrong number of lines: expected %d but got %d", 0, l)
    }
}

func TestStringCounter(t *testing.T) {
    for _, tt := range testCases {
        c := makeCounter()
        c.AddString(tt.s)
        if l := c.Lines(); tt.lines != l {
            t.Fatalf("Wrong number of lines: expected %d but got %d", tt.lines, l)
        }
        if l := c.Letters(); tt.letters != l {
            t.Fatalf("Wrong number of letters: expected %d but got %d", tt.letters, l)
        }
        if l := c.Characters(); tt.chars != l {
            t.Fatalf("Wrong number of lines: expected %d but got %d", tt.chars, l)
        }
    }
}
