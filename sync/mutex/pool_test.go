package main

import (
	"sync"
	"testing"
)

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < 100*read; k++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				rw.Read()
			}()
		}

		for j := 0; j < 100*write; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				rw.Write()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkReadMore(b *testing.B) { benchmark(b, &Lock{}, 9, 1) }

func BenchmarkReadMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 9, 1) }

func BenchmarkWriteMore(b *testing.B) { benchmark(b, &Lock{}, 1, 9) }

func BenchmarkWriteMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 1, 9) }

func BenchmarkReadWriteEqual(b *testing.B) { benchmark(b, &Lock{}, 5, 5) }

func BenchmarkReadWriteEqualRW(b *testing.B) { benchmark(b, &RWLock{}, 5, 5) }
