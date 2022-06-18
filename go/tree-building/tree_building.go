package tree

import (
	"errors"
	"sort"
)

/*

Brute Force:

One option we have is to find the root first,
then look for the children, then its children, and so on.
But this would have a time complexity of O(n^2).

Let's look at the notes below to find a better approach.

Notes:
1. The records only contain an ID number and a parent ID number.
2. The ID number is always between 0 (inclusive) and the length of the record list (exclusive).
3. All records have a parent ID lower than their own ID, except for the root record, which has a parent ID that's equal to its own ID.

Giving that we have unique IDs, they are always between 0 and the length of the records,
and also parents will have a lower ID, we can:

- Sort the records by ID
- Create an array to store the nodes
- Iterate over the records
	- Create a new node and add to the array of nodes
	- Get its parent in the array and add itself as a child
- Return the first node in the array as the root

Time Complexity: O(nlgn)
Space Complexity: O(n)
*/

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodes := make([]*Node, 0, len(records))

	// Sort records by ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for _, r := range records {
		// Check parent comes before the node
		if r.Parent > r.ID {
			return nil, errors.New("Parent cannot come after the node")
		}

		// Check for cycles
		if r.ID == r.Parent && r.ID != 0 {
			return nil, errors.New("Node cannot point to itself")
		}

		// Check for invalid entries
		if r.ID > len(nodes) {
			return nil, errors.New("Record has invalid ID")
		}

		// Check for duplicates
		if r.ID < len(nodes) {
			return nil, errors.New("Cannot have duplicate records")
		}

		// Create node from record
		newNode := &Node{r.ID, nil}
		nodes = append(nodes, newNode)

		// Add node as child of the parent
		if r.Parent != r.ID {
			parent := nodes[r.Parent]
			parent.Children = append(parent.Children, newNode)
		}
	}

	return nodes[0], nil
}

/*
## Benchmarks ##

> Without allocation capacity for the nodes slice
var nodes []*Node

BenchmarkTwoTree
BenchmarkTwoTree-8                    74          14949914 ns/op         5801320 B/op     131102 allocs/op
BenchmarkTenTree
BenchmarkTenTree-8                   933           1965076 ns/op          954386 B/op      15023 allocs/op
BenchmarkShallowTree
BenchmarkShallowTree-8               676           2326532 ns/op         1092682 B/op      10043 allocs/op

> Allocating capacity for the nodes slice
nodes := make([]*Node, 0, len(records))

BenchmarkTwoTree
BenchmarkTwoTree-8                    82          14207013 ns/op         3407951 B/op     131075 allocs/op
BenchmarkTenTree
BenchmarkTenTree-8                  1072           1514248 ns/op          650009 B/op      15004 allocs/op
BenchmarkShallowTree
BenchmarkShallowTree-8              1207           1401872 ns/op          788305 B/op      10024 allocs/op
*/
