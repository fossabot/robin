package robin

import (
	"sync"
)

type Disposable interface {
	Dispose()
	Identify() string
}

type Disposer struct {
	disposables ConcurrentMap
	lock        *sync.Mutex
}

func (d *Disposer) init() *Disposer {
	d.disposables = NewConcurrentMap()
	d.lock = new(sync.Mutex)
	return d
}

func NewDisposer() *Disposer {
	return new(Disposer).init()
}

func (d *Disposer) Add(disposable Disposable) {
	d.disposables.Set(disposable.Identify(), disposable)
}

func (d *Disposer) Remove(disposable Disposable) {
	d.disposables.Remove(disposable.Identify())
}

func (d *Disposer) Count() int {
	return d.disposables.Count()
}

func (d *Disposer) Dispose() {
	d.lock.Lock()
	defer d.lock.Unlock()
	for _, key := range d.disposables.Keys() {
		if tmp, ok := d.disposables.Pop(key); ok {
			tmp.(Disposable).Dispose()
		}
	}
}
