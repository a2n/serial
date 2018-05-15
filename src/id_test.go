package serial

import "testing"

func BenchmarkIdIncrease(b *testing.B) {
	is := NewIDService()
	is.Set(0)
	for i := 0; i < b.N; i++ {
		is.Increase()
	}
	b.Logf("Value %d", is.Get())
}

func BenchmarkIdIncreaseParallel(b *testing.B) {
	is := NewIDService()
	is.Set(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			is.Increase()
		}
		b.Logf("Value %d", is.Get())
	})
}
