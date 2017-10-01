package main

import (
	"fmt"
	"sync"
)

// execute jobs
type worker struct {
	dispatcher *Dispatcher
	data       chan interface{}
	quit       chan struct{}
}

// Start worker
func (w *worker) start() {
	go func() {
		for {
			w.dispatcher.pool <- w
			select {
			case v := <-w.data:
				if str, ok := v.(string); ok {
					fmt.Println("Download ", str)
				}
				w.dispatcher.wg.Done()
			case <-w.quit:
				return
			}
		}
	}()
}

// Dispatch workers
type Dispatcher struct {
	pool    chan *worker
	queue   chan interface{}
	workers []*worker
	wg      sync.WaitGroup
	quit    chan struct{}
}

// Add values
func (d *Dispatcher) Add(v interface{}) {
	d.wg.Add(1)
	d.queue <- v
}

// Wait to finish job
func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

// Wait for the dispatcher
func (d *Dispatcher) Stop(immediately bool) {
	if !immediately {
		d.Wait()
	}

	d.quit <- struct{}{}
	for _, w := range d.workers {
		w.quit <- struct{}{}
	}
}

// Start dispatcher but does not wait for it completed
func (d *Dispatcher) Start() {
	for _, w := range d.workers {
		w.start()
	}
	go func() {
		for {
			select {
			case v := <-d.queue:
				(<-d.pool).data <- v
			case <-d.quit:
				return
			}
		}
	}()
}

const (
	maxWorkers = 3
	maxQueues  = 10000
)

// NewDispatcher returns a Dispatcher pointer.
func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		pool:  make(chan *worker, maxWorkers),
		queue: make(chan interface{}, maxQueues),
		quit:  make(chan struct{}),
	}
	d.workers = make([]*worker, cap(d.pool))
	for _, w := range d.workers {
		w = &worker{
			dispatcher: d,
			data:       make(chan interface{}),
			quit:       make(chan struct{}),
		}
		_ = w
	}
	return d
}

func main() {
	d := NewDispatcher()
	d.Start()

	for i := 0; i < 100; i++ {
		d.Add(fmt.Sprintf("http://example-%d", i))
	}
	d.Wait()
}
