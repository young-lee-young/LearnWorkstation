/**
	集合
 */
package set

import (
	"sync"
)

var void = struct {}{}

type Set struct {
	sync.RWMutex
	Data map[interface{}]struct{}
}

func New() *Set {
	set := &Set{
		Data: make(map[interface{}]struct{}),
	}
	return set
}

func (s *Set) Add(data interface{}) {
	s.Lock()
	defer s.Unlock()
	s.Data[data] = void
}

func (s *Set) Remove(data interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.Data, data)
}

func (s *Set) Contains(data interface{}) bool {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Data[data]; ok {
		return true
	}
	return false
}

func (s *Set) GetLength() int {
	s.Lock()
	defer s.Unlock()
	return len(s.Data)
}

func (s *Set) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()
	if len(s.Data) == 0 {
		return true
	}
	return false
}
