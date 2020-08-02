package bob

const (
    fine    = "Fine. Be that way!"
    sure    = "Sure."
    yell    = "Whoa, chill out!"
    yellQ   = "Calm down, I know what I'm doing!"
    casual  = "Whatever."
)

func Hey(s string) string {

    if len(s) == 0 {
        return fine
    }

    var lastChar rune
    onlySpaces := true
    hasCaps := false
    allCaps := true

    for _, c := range s {
        if c > ' ' {
            onlySpaces = false
            lastChar = c
        }
        if 'a' <= c && c <= 'z' {
            allCaps = false
        } else if 'A' <= c && c <= 'Z' {
            hasCaps = true
        }
    }

    isQuestion := (lastChar == '?')

    if onlySpaces {
        return fine
    }

    if hasCaps && allCaps {
        if isQuestion {
            return yellQ
        }
        return yell
    }

    if isQuestion {
        return sure
    }

    return casual
}
