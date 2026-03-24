package main

import "testing"

// TestAdd, Add fonksiyonunun doğruluğunu test eder.
func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("Add(4, 6) = %d; want %d", got, want)
	}
}

// BenchmarkAdd, Add fonksiyonunun performansını ölçer.
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}
