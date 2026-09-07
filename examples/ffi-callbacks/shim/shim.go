package shim

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var active atomic.Int64

type Counter struct {
	value atomic.Int64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Add(value int64) int64 {
	return c.value.Add(value)
}

func (c *Counter) Value() int64 {
	return c.value.Load()
}

type Subscription struct {
	mu       sync.Mutex
	callback func(int64) int64
}

func Register(callback func(int64) int64) *Subscription {
	if callback != nil {
		active.Add(1)
	}
	return &Subscription{callback: callback}
}

func (s *Subscription) Unregister() {
	s.mu.Lock()
	if s.callback != nil {
		s.callback = nil
		active.Add(-1)
	}
	s.mu.Unlock()
}

type Batch struct {
	gate chan struct{}
	done chan int64
	once sync.Once
}

func (s *Subscription) Start(count int64, value int64) *Batch {
	batch := &Batch{gate: make(chan struct{}), done: make(chan int64, 1)}
	go func() {
		<-batch.gate
		var workers sync.WaitGroup
		var calls atomic.Int64
		for i := int64(0); i < count; i++ {
			workers.Add(1)
			go func() {
				defer workers.Done()
				s.mu.Lock()
				callback := s.callback
				s.mu.Unlock()
				if callback != nil {
					callback(value)
					calls.Add(1)
				}
			}()
		}
		workers.Wait()
		batch.done <- calls.Load()
	}()
	return batch
}

func (b *Batch) Release() {
	b.once.Do(func() { close(b.gate) })
}

func (b *Batch) Wait() int64 {
	return <-b.done
}

func Active() int64 {
	return active.Load()
}

func Collect() {
	runtime.GC()
}

func Reenter(callback func(int64) int64, value int64) int64 {
	return callback(value)
}
