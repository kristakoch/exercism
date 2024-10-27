package tree

import (
	"fmt"
	"sort"
)

// Record holds data for Build input tree records.
type Record struct {
	ID     int
	Parent int
}

// Node holds data for Build output tree nodes.
type Node struct {
	ID       int
	Children []*Node
}

// Build outputs a tree of nodes based on input records.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// The root has the lowest-value ID and is its own parent.
	if records[0].ID != records[0].Parent {
		return nil, fmt.Errorf("invalid root %#v", records[0])
	}

	nodes := make([]*Node, len(records))

	for i, rec := range records {

		// Check for non-consecutive IDs and non-roots with invalid parent IDs.
		if rec.ID != i || ((rec.ID == rec.Parent) && rec.ID != 0) {
			return nil, fmt.Errorf("invalid record %#v", rec)
		}

		newNode := &Node{ID: rec.ID}

		nodes[i] = newNode

		// All non-roots should be assigned to parents.
		if rec.ID != 0 {

			// The parent should already have been added to the group of nodes.
			if nodes[rec.Parent] == nil {
				return nil, fmt.Errorf(
					"parent ID %d larger than child ID %d, cycling", 
					rec.Parent, 
					rec.ID,
				)
			}

			nodes[rec.Parent].Children = append(nodes[rec.Parent].Children, newNode)
		}
	}

	return nodes[0], nil
}
