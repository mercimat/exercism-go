package isbn

import (
    "unicode"
    "regexp"
)

var isbnPattern = "^((\\d{9})|(\\d\\-\\d{3}-\\d{5}-))(\\d|X)$"
var pattern, _ = regexp.Compile(isbnPattern)

func IsValidISBNRegex(isbn string) bool {
    if ! pattern.MatchString(isbn) {
        return false
    }
    var multiplier = 10
    var acc = 0
    for _, v := range isbn {
        if v == '-' {
            continue
        }
        if v == 'X' {
            acc += 10
            multiplier -= 1
        } else {
            acc += int(v - '0') * multiplier
            multiplier -= 1
        }
    }
    return ((acc % 11) == 0)
}

func IsValidISBN(isbn string) bool {
    if !(len(isbn) == 10 || len(isbn) == 13) {
        return false
    }

    var multiplier = 10
    var acc = 0
    for i, v := range isbn {
        if ! isValidChar(v, i, len(isbn)) {
            return false
        }
        if v == '-' {
            continue
        }
        if v == 'X' {
            acc += 10
            multiplier -= 1
        } else {
            acc += int(v - '0') * multiplier
            multiplier -= 1
        }
    }
    return ((acc % 11) == 0)
}

func isValidChar(c rune, i int, l int) bool {
    if l == 10 {
        if unicode.IsDigit(c) {
            return true
        } else if c == 'X' && i == l-1 {
            return true
        }
    } else if l == 13 {
        if i == 1 || i == 5 || i == 11 {
            if c == '-' {
                return true
            }
        } else if unicode.IsDigit(c) {
            return true
        } else if c == 'X' && i == l-1 {
            return true
        }
    }
    return false
}
