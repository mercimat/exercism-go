package tree

import "fmt"

type Record struct {
    ID      int
    Parent  int
}

type Node struct {
    ID          int
    Children    []*Node
}


func Build(records []Record) (*Node, error) {
    if len(records) == 0 {
        return nil, nil
    }

    nodes := make([]Node, len(records))
    parents := make([]*Node, len(records))
    seen := make([]bool, len(records))

    for _, record := range records {
        if record.ID >= len(records) {
            return nil, fmt.Errorf("Too high id %d", record.ID)
        }
        if record.ID <= record.Parent && record.ID > 0 {
            return nil, fmt.Errorf("ID lower than its parent's: id %d, parent %d", record.ID, record.Parent)
        }
        if seen[record.ID] {
            return nil, fmt.Errorf("Duplicate ID: id %d", record.ID)
        }
        seen[record.ID] = true
        if record.ID == 0 && record.Parent > 0 {
            return nil, fmt.Errorf("Unexpected values for root record: id %d, parent %d", record.ID, record.Parent)
        }

        parents[record.ID] = &nodes[record.Parent]
    }

    for i:=1; i<len(records); i++ {
        parents[i].Children = append(parents[i].Children, &nodes[i])
    }
    for i:=1; i<len(records); i++ {
        nodes[i].ID = i
    }

    return &nodes[0], nil
}
