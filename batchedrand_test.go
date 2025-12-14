package batchedrand

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

// Setup: A slice of integers to shuffle
func getSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}
	return s
}

func BenchmarkChaChaShuffle(b *testing.B) {
	sizes := []int{30, 100, 500000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Batched_size_%d", size), func(b *testing.B) {
			rng := Rand{rand.New(rand.NewChaCha8([32]byte{1, 2, 3}))}
			data := getSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rng.Shuffle(len(data), func(i, j int) {
					data[i], data[j] = data[j], data[i]
				})
			}
		})
		b.Run(fmt.Sprintf("Standard_size_%d", size), func(b *testing.B) {
			rng := rand.New(rand.NewChaCha8([32]byte{1, 2, 3}))
			data := getSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rng.Shuffle(len(data), func(i, j int) {
					data[i], data[j] = data[j], data[i]
				})
			}
		})
	}
}

func BenchmarkPCGShuffle(b *testing.B) {
	sizes := []int{30, 100, 500000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Batched_size_%d", size), func(b *testing.B) {
			rng := Rand{rand.New(rand.NewPCG(1, 2))}
			data := getSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rng.Shuffle(len(data), func(i, j int) {
					data[i], data[j] = data[j], data[i]
				})
			}
		})
		b.Run(fmt.Sprintf("Standard_size_%d", size), func(b *testing.B) {
			rng := rand.New(rand.NewPCG(1, 2))
			data := getSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rng.Shuffle(len(data), func(i, j int) {
					data[i], data[j] = data[j], data[i]
				})
			}
		})
	}
}

func TestShuffle_ChaCha(t *testing.T) {
	const size = 1000
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	// Use ChaCha8 with a fixed key for reproducibility
	var key [32]byte
	for i := range key {
		key[i] = byte(i)
	}
	rng := Rand{rand.New(rand.NewChaCha8(key))}

	// Values to check: 0, 100, 200, ..., 900, 999
	var valuesToCheck []int
	for i := 0; i < size; i += 100 {
		valuesToCheck = append(valuesToCheck, i)
	}
	valuesToCheck = append(valuesToCheck, 999)

	positions := make(map[int]map[int]bool) // value -> position -> seen
	for _, val := range valuesToCheck {
		positions[val] = make(map[int]bool)
	}

	const numShuffles = 50000
	for i := 0; i < numShuffles; i++ {
		copyData := make([]int, size)
		copy(copyData, data)
		rng.Shuffle(len(copyData), func(i, j int) {
			copyData[i], copyData[j] = copyData[j], copyData[i]
		})
		// Find positions of the values to check
		for _, val := range valuesToCheck {
			for j, v := range copyData {
				if v == val {
					positions[val][j] = true
					break
				}
			}
		}
	}
	// Check that each value appeared in all positions
	for _, val := range valuesToCheck {
		for i := 0; i < size; i++ {
			if !positions[val][i] {
				t.Errorf("Position %d not seen for value %d after %d shuffles", i, val, numShuffles)
			}
		}
	}
}

func TestShuffle_PCG(t *testing.T) {
	const size = 1000
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	// Use PCG with fixed seeds for reproducibility
	rng := Rand{rand.New(rand.NewPCG(1, 2))}

	// Values to check: 0, 100, 200, ..., 900, 999
	var valuesToCheck []int
	for i := 0; i < size; i += 100 {
		valuesToCheck = append(valuesToCheck, i)
	}
	valuesToCheck = append(valuesToCheck, 999)

	positions := make(map[int]map[int]bool) // value -> position -> seen
	for _, val := range valuesToCheck {
		positions[val] = make(map[int]bool)
	}

	const numShuffles = 50000
	for i := 0; i < numShuffles; i++ {
		copyData := make([]int, size)
		copy(copyData, data)
		rng.Shuffle(len(copyData), func(i, j int) {
			copyData[i], copyData[j] = copyData[j], copyData[i]
		})
		// Find positions of the values to check
		for _, val := range valuesToCheck {
			for j, v := range copyData {
				if v == val {
					positions[val][j] = true
					break
				}
			}
		}
	}
	// Check that each value appeared in all positions
	for _, val := range valuesToCheck {
		for i := 0; i < size; i++ {
			if !positions[val][i] {
				t.Errorf("Position %d not seen for value %d after %d shuffles", i, val, numShuffles)
			}
		}
	}
}
