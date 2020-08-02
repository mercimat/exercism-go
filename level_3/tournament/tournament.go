package tournament

import (
    "encoding/csv"
    "fmt"
    "io"
    "sort"
)

type outcome int

const (
    loss outcome = iota
    draw
    win
)


type inputEntry struct {
    team1   string
    team2   string
    outcome outcome
}

type Score struct {
    Team    string
    MP      int
    W       int
    D       int
    L       int
    P       int
}

type OrderedScores []Score

func (s OrderedScores) Len() int           { return len(s) }
func (s OrderedScores) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s OrderedScores) Less(i, j int) bool { return s[i].P < s[j].P || (s[i].P == s[j].P && s[i].Team > s[j].Team ) }

func readInput(r io.Reader) ([]inputEntry, error) {
    var entries []inputEntry
    csvReader := csv.NewReader(r)
    csvReader.Comma = ';'
    csvReader.Comment = '#'
    csvReader.FieldsPerRecord = -1

    for {
        record, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        if len(record) == 3 {
            entry := inputEntry{team1: record[0], team2: record[1]}
            switch record[2] {
            case "win":
                entry.outcome = win
            case "loss":
                entry.outcome = loss
            case "draw":
                entry.outcome = draw
            default:
                return nil, fmt.Errorf("Invalid record: %v", record)
            }
            entries = append(entries, entry)
        } else {
            return nil, fmt.Errorf("Invalid record: %v", record)
        }
    }
    return entries, nil
}

func tallyEntries(entries []inputEntry) map[string]Score {
    var scores = make(map[string]Score)
    for _, entry := range entries {
        team1 := entry.team1
        team2 := entry.team2
        result := entry.outcome
        if _, ok := scores[team1]; !ok {
            scores[team1] = Score{Team: team1}
        }
        if _, ok := scores[team2]; !ok {
            scores[team2] = Score{Team: team2}
        }
        s1 := scores[team1]
        s2 := scores[team2]
        s1.MP += 1
        s2.MP += 1
        if result == win {
            s1.W += 1
            s1.P += 3
            s2.L += 1
        } else if result == loss {
            s2.W += 1
            s2.P += 3
            s1.L += 1
        } else if result == draw {
            s1.D += 1
            s1.P += 1
            s2.D += 1
            s2.P += 1
        }
        scores[team1] = s1
        scores[team2] = s2
    }
    return scores
}

func report(w io.Writer, scores map[string]Score) {
    var orderedScores = make([]Score, 0)
    for _, v := range scores {
        orderedScores = append(orderedScores, Score{Team: v.Team, MP: v.MP, W: v.W, D: v.D, L: v.L, P: v.P})
    }
    sort.Sort(OrderedScores(orderedScores))

    fmt.Fprintf(w, "Team                           | MP |  W |  D |  L |  P\n")
    for i:=len(orderedScores)-1; i >=0; i-- {
        entry := orderedScores[i]
        fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
            entry.Team, entry.MP, entry.W, entry.D, entry.L, entry.P)
    }
}

func Tally(r io.Reader, w io.Writer) error {
    entries, err := readInput(r)
    if err != nil {
        return err
    }
    report(w, tallyEntries(entries))
    return nil
}
