package acronym

import (
    "unicode"
    "regexp"
)

var re = regexp.MustCompile("[a-zA-Z']+\\W*")

// More optimal solution
func Abbreviate(s string) string {
    abbreviation := make([]rune, 0)

    var newWord = true
    for _, r := range s {
        if unicode.IsLetter(r) {
            if newWord {
                abbreviation = append(abbreviation, unicode.ToUpper(r))
            }
            newWord = false
        } else if r == '\'' {
            continue
        } else {
            newWord = true
        }
    }

    return string(abbreviation)
}

// Exercise around using regexp - less optimal
func Abbreviate2(s string) string {
    abbreviation := make([]rune, 0)
    for _, match := range re.FindAllString(s, -1) {
        abbreviation = append(abbreviation, unicode.ToUpper(rune(match[0])))
    }
    return string(abbreviation)
}
