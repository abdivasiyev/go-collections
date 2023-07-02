# Collections in Go

[![codecov](https://codecov.io/gh/abdivasiyev/go-collections/branch/master/graph/badge.svg?token=tNKcOjlxLo)](https://codecov.io/gh/abdivasiyev/go-collections)
[![Tests](https://github.com/abdivasiyev/go-collections/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/abdivasiyev/go-collections/actions/workflows/ci.yml)

*Hey, now Java collections are here!*

@author Asliddin Abdivasiyev

Collections implemented by Go generics, and
you can use any comparable type with them

Requirements:
- go version >= 1.18

Installation:
```shell
go get -u github.com/abdivasiyev/go-collections
```

Usage:

```go
package main

// import collections package first
import (
	"fmt"
	"github.com/abdivasiyev/go-collections"
)

func main() {
	// initialize ArrayList collection with type of int
	list := collections.NewArrayList[int]()

	// add elements with Add method
	list.Add(1)
	list.Add(2)

	// loop through elements
	// and print them out
	for i := 0; i < list.Size(); i++ {
		fmt.Println(list.Get(i))
	}
}
```

Available collections:

- [ ] List
  - [x] ArrayList
  - [x] LinkedList
- [ ] Queue
  - [ ] Queue
  - [ ] PriorityQueue
- [ ] Deque
  - [ ] ArrayDeque
- [ ] Set
  - [ ] HashSet
  - [ ] LinkedHashSet
  - [ ] SortedSet
- [ ] Stack
- [ ] Map
  - [ ] HashMap
  - [ ] LinkedHashMap
  - [ ] SortedMap
- [ ] Tree
  - [ ] BinaryTree
  - [ ] TreeMap
  - [ ] TreeSet
  - [ ] Heap
    - [ ] MaxHeap
    - [ ] MinHeap
