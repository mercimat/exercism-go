package romannumerals

import (
    "bytes"
    "fmt"
)

type mapping struct {
    number int
    roman string
}

var romanMappings = []mapping{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

func ToRomanNumeral(i int) (string, error) {
    buffer := bytes.NewBufferString("")
    if i < 1 || i > 3000 {
        return "", fmt.Errorf("not a roman number")
    }

    for _, mapping := range romanMappings {
        for i >= mapping.number {
            buffer.WriteString(mapping.roman)
            i -= mapping.number
        }
    }
    return buffer.String(), nil
}
