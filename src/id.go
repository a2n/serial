package serial

import (
	"sync/atomic"
)

type IdService struct {
	id uint64
}

func NewIdService() *IdService {
	return &IdService{}
}

// Getting the value.
func (this *IdService) Get() uint64 {
	return atomic.LoadUint64(&this.id)
}

// Setting the value.
func (this *IdService) Set(n uint64) {
	atomic.StoreUint64(&this.id, n)
}

// Increasing the value by one.
func (this *IdService) Increase() uint64 {
	// Math.MaxUint64++ -> 0
	return atomic.AddUint64(&this.id, 1)
}
