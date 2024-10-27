package tree

// import (
// 	"errors"
// 	"sort"
// )

// // Record ...
// type Record struct {
// 	ID     int
// 	Parent int
// }

// // Node ...
// type Node struct {
// 	ID       int
// 	Children []*Node
// }

// // Build ...
// func Build(records []Record) (*Node, error) {
// 	var err error

// 	if len(records) == 0 {
// 		return nil, nil
// 	}

// 	// Sort the records so the root can be picked out.
// 	sort.Slice(records, func(i, j int) bool {
// 		return records[i].ID < records[j].ID
// 	})

// 	nodeMap := make(map[int][]int)

// 	var rootID, counter int

// 	// Biuld a map of the records to use for building the tree.
// 	for i, rec := range records {

// 		// Set the ID of the root.
// 		if i == 0 {
// 			if rootID = records[0].ID; records[0].ID != records[0].Parent {
// 				return nil, errors.New("root has invalid parent value")
// 			}
// 			counter = rootID + 1
// 			continue
// 		}

// 		if rec.ID != counter {
// 			return nil, errors.New("node IDs non-continuous")
// 		}
// 		counter++

// 		if rec.ID <= rec.Parent {
// 			return nil, errors.New("parent node has larger or equal value to child")
// 		}

// 		nodeMap[rec.Parent] = append(nodeMap[rec.Parent], rec.ID)
// 	}

// 	// Create the children trees here with the recursive
// 	// createChildren function.
// 	var root *Node
// 	if root, err = createChildren(rootID, nodeMap); nil != err {
// 		return nil, err
// 	}

// 	return root, nil
// }

// // createChildren is a recursive function to
// // create the tree of nodes based on a nodeMap.
// func createChildren(rootID int, nodeMap map[int][]int) (*Node, error) {
// 	rootChildIDs := nodeMap[rootID]

// 	sort.Ints(rootChildIDs)

// 	var childNodes []*Node
// 	for _, childID := range rootChildIDs {
// 		var err error

// 		var node *Node
// 		if node, err = createChildren(childID, nodeMap); nil != err {
// 			return nil, err
// 		}
// 		childNodes = append(childNodes, node)
// 	}

// 	return &Node{rootID, childNodes}, nil
// }
