package serial

import (
	"sync/atomic"
)

// IDService 識別服務
type IDService struct {
	id uint64
}

// NewIDService 創建服務
func NewIDService() *IDService {
	return &IDService{}
}

// Get 取值
func (is *IDService) Get() uint64 {
	return atomic.LoadUint64(&is.id)
}

// Set 設值
func (is *IDService) Set(n uint64) {
	atomic.StoreUint64(&is.id, n)
}

// Increase 遞增
func (is *IDService) Increase() uint64 {
	// Math.MaxUint64++ -> 0
	return atomic.AddUint64(&is.id, 1)
}
