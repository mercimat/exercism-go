package scale

import (
    "strings"
)

var chromaticScaleWithSharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var chromaticScaleWithFlats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var isFlat = map[string]bool{
    "F": true,
    "Bb": true,
    "Eb": true,
    "Ab": true,
    "Db": true,
    "Gb": true,
    "d": true,
    "g": true,
    "c": true,
    "f": true,
    "bb": true,
    "eb": true,
}

func Scale(tonic, interval string) []string {
    if _, ok := isFlat[tonic]; !ok {
        return scale(strings.Title(tonic), interval, chromaticScaleWithSharps)
    } else {
        return scale(strings.Title(tonic), interval, chromaticScaleWithFlats)
    }
}

func scale(tonic string, interval string, scale []string) []string {
    i := 0
    for ; i < len(scale); i++ {
        if scale[i] == tonic {
            break
        }
    }
    if interval == "" {
        return append(scale[i:], scale[:i]...)
    }
    res := []string{}
    res = append(res, scale[i])
    for _, v := range interval {
        switch v {
        case 'm':
            i += 1
        case 'M':
            i += 2
        case 'A':
            i += 3
        }
        i %= len(scale)
        if scale[i] == res[0] {
            break
        }
        res = append(res, scale[i])
    }
    return res
}
