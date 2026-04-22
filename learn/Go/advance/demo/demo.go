package demo

import (
	"context"
	"sync"
)

const (
	maxParallelMemfileSnapshotting = 64
)

type DynamicSemaphore struct {
	mu      sync.Mutex
	cond    *sync.Cond
	current int
	max     int
}

func NewDynamicSemaphore(max int) *DynamicSemaphore {
	s := &DynamicSemaphore{max: max}
	s.cond = sync.NewCond(&s.mu)

	go func() {
	}()

	return s
}

func (s *DynamicSemaphore) Acquire(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for s.current >= s.max {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		s.cond.Wait()
	}

	s.current++
	return nil
}

func (s *DynamicSemaphore) Release() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current--
	s.cond.Signal()
}

func (s *DynamicSemaphore) SetMax(max int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.max = max
	s.cond.Broadcast()
}

func (s *DynamicSemaphore) GetMax() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.max
}

func (s *DynamicSemaphore) GetCurrent() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.current
}

var snapshotCacheQueue = NewDynamicSemaphore(maxParallelMemfileSnapshotting)

func SetMaxDynamicSemaphore(max int) {
	snapshotCacheQueue.SetMax(max)
}

func GetMaxDynamicSemaphore() int {
	return snapshotCacheQueue.GetMax()
}

func GetCurDynamicSemaphore() int {
	return snapshotCacheQueue.GetCurrent()
}
