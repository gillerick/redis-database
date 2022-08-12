package db

type Memory struct {
	MainStore     Store
	ReversedStore ReversedStore
}

type Stack []Memory

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(store Store, reversedStore ReversedStore) {
	*s = append(*s, Memory{store, reversedStore})
}

func (s *Stack) Pop() (Memory, bool) {
	if s.IsEmpty() {
		return Memory{}, false
	} else {
		lastIndex := len(*s) - 1
		lastStruct := (*s)[lastIndex]

		*s = (*s)[:lastIndex]

		return lastStruct, true
	}
}
