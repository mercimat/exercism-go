package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
    var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

    for _, r := range d {
        if _, ok := h[r]; ok {
            h[r] += 1
        } else {
            return h, errors.New("invalid nucleotide found in DNA")
        }
    }

    return h, nil
}
