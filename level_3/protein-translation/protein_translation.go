package protein

import "errors"

var ErrStop = errors.New("stop error")
var ErrInvalidBase = errors.New("invalid base error")

var codonToProtein = map[string]string{
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": "STOP",
    "UAG": "STOP",
    "UGA": "STOP",
}

func FromCodon(codon string) (string, error) {
    if len(codon) != 3 {
        return "", ErrInvalidBase
    }
    v, ok := codonToProtein[codon]
    if !ok {
        return "", ErrInvalidBase
    }
    if v == "STOP" {
        return "", ErrStop
    }
    return v, nil
}

func FromRNA(seq string) ([]string, error) {
    var res = make([]string, 0)
    for i:=0; i<len(seq); i+=3 {
        v, err := FromCodon(seq[i:i+3])
        if err == ErrStop {
            break
        } else if err != nil {
            return res, err
        }
        res = append(res, string(v))
    }
    return res, nil
}
