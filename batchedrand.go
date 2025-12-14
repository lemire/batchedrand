package batchedrand

import (
	"math/bits"
	"math/rand/v2"
)

type Rand struct {
	*rand.Rand
}

// Shuffle pseudo-randomizes the order of elements.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func (r *Rand) Shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}
	i := uint64(n)

	// Single swaps for sizes > 2^30
	for ; i > (1 << 30); i-- {
		var index0 uint64
		randVal := r.Uint64()
		size := i
		bound := i
		// Unroll loop for k=1
		a := size
		hi, lo := bits.Mul64(a, randVal)
		randVal = lo
		index0 = hi
		if randVal < bound {
			newBound := size
			t := (-newBound) % newBound
			for randVal < t {
				randVal = r.Uint64()
				a = size
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=1
		pos1 := size - 1
		pos2 := index0
		swap(int(pos1), int(pos2))
	}

	// Batches of 2 down to 2^19
	for ; i > (1 << 19); i -= 2 {
		// Inline partialShuffle64b(storage, i, 2, bound, rng)
		var index0, index1 uint64
		randVal := r.Uint64()
		size := i
		bound := uint64(1) << 60
		// Unroll loop for k=2
		a := size - 1
		hi, lo := bits.Mul64(a, randVal)
		randVal = lo
		index1 = hi
		a = size
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index0 = hi
		if randVal < bound {
			newBound := size * (size - 1)
			t := (-newBound) % newBound
			for randVal < t {
				randVal = r.Uint64()
				a = size - 1
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index1 = hi
				a = size
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=2
		pos1 := size - 2
		pos2 := index1
		swap(int(pos1), int(pos2))
		pos1 = size - 1
		pos2 = index0
		swap(int(pos1), int(pos2))
	}

	// Batches of 3 down to 2^14
	for ; i > (1 << 14); i -= 3 {
		// Inline partialShuffle64b(storage, i, 3, bound, rng)
		var index0, index1, index2 uint64
		randVal := r.Uint64()
		size := i
		bound := uint64(1) << 57
		// Unroll loop for k=3
		a := size - 2
		hi, lo := bits.Mul64(a, randVal)
		randVal = lo
		index2 = hi
		a = size - 1
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index1 = hi
		a = size
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index0 = hi
		if randVal < bound {
			newBound := size * (size - 1) * (size - 2)
			t := (-newBound) % newBound
			for randVal < t {
				randVal = r.Uint64()
				a = size - 2
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index2 = hi
				a = size - 1
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index1 = hi
				a = size
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=3
		pos1 := size - 3
		pos2 := index2
		swap(int(pos1), int(pos2))
		pos1 = size - 2
		pos2 = index1
		swap(int(pos1), int(pos2))
		pos1 = size - 1
		pos2 = index0
		swap(int(pos1), int(pos2))
	}

	// Batches of 4 down to 2^11
	for ; i > (1 << 11); i -= 4 {
		// Inline partialShuffle64b(storage, i, 4, bound, rng)
		var index0, index1, index2, index3 uint64
		randVal := r.Uint64()
		size := i
		bound := uint64(1) << 56
		// Unroll loop for k=4
		a := size - 3
		hi, lo := bits.Mul64(a, randVal)
		randVal = lo
		index3 = hi
		a = size - 2
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index2 = hi
		a = size - 1
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index1 = hi
		a = size
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index0 = hi
		if randVal < bound {
			newBound := size * (size - 1) * (size - 2) * (size - 3)
			t := (-newBound) % newBound
			for randVal < t {
				randVal = r.Uint64()
				a = size - 3
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index3 = hi
				a = size - 2
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index2 = hi
				a = size - 1
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index1 = hi
				a = size
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=4
		pos1 := size - 4
		pos2 := index3
		swap(int(pos1), int(pos2))
		pos1 = size - 3
		pos2 = index2
		swap(int(pos1), int(pos2))
		pos1 = size - 2
		pos2 = index1
		swap(int(pos1), int(pos2))
		pos1 = size - 1
		pos2 = index0
		swap(int(pos1), int(pos2))
	}

	// Batches of 5 down to 2^9
	for ; i > (1 << 9); i -= 5 {
		// Inline partialShuffle64b(storage, i, 5, bound, rng)
		var index0, index1, index2, index3, index4 uint64
		randVal := r.Uint64()
		size := i
		bound := uint64(1) << 55
		// Unroll loop for k=5
		a := size - 4
		hi, lo := bits.Mul64(a, randVal)
		randVal = lo
		index4 = hi
		a = size - 3
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index3 = hi
		a = size - 2
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index2 = hi
		a = size - 1
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index1 = hi
		a = size
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index0 = hi
		if randVal < bound {
			newBound := size * (size - 1) * (size - 2) * (size - 3) * (size - 4)
			t := (-newBound) % newBound
			for randVal < t {
				randVal = r.Uint64()
				a = size - 4
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index4 = hi
				a = size - 3
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index3 = hi
				a = size - 2
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index2 = hi
				a = size - 1
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index1 = hi
				a = size
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=5
		pos1 := size - 5
		pos2 := index4
		swap(int(pos1), int(pos2))
		pos1 = size - 4
		pos2 = index3
		swap(int(pos1), int(pos2))
		pos1 = size - 3
		pos2 = index2
		swap(int(pos1), int(pos2))
		pos1 = size - 2
		pos2 = index1
		swap(int(pos1), int(pos2))
		pos1 = size - 1
		pos2 = index0
		swap(int(pos1), int(pos2))
	}

	// Batches of 6 down to 6
	for ; i > 6; i -= 6 {
		// Inline partialShuffle64b(storage, i, 6, bound, rng)
		var index0, index1, index2, index3, index4, index5 uint64
		randVal := r.Uint64()
		size := i
		bound := uint64(1) << 54
		// Unroll loop for k=6
		a := size - 5
		hi, lo := bits.Mul64(a, randVal)
		randVal = lo
		index5 = hi
		a = size - 4
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index4 = hi
		a = size - 3
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index3 = hi
		a = size - 2
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index2 = hi
		a = size - 1
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index1 = hi
		a = size
		hi, lo = bits.Mul64(a, randVal)
		randVal = lo
		index0 = hi
		if randVal < bound {
			newBound := size * (size - 1) * (size - 2) * (size - 3) * (size - 4) * (size - 5)
			t := (-newBound) % newBound
			for randVal < t {
				randVal = r.Uint64()
				a = size - 5
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index5 = hi
				a = size - 4
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index4 = hi
				a = size - 3
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index3 = hi
				a = size - 2
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index2 = hi
				a = size - 1
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index1 = hi
				a = size
				hi, lo = bits.Mul64(a, randVal)
				randVal = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=6
		pos1 := size - 6
		pos2 := index5
		swap(int(pos1), int(pos2))
		pos1 = size - 5
		pos2 = index4
		swap(int(pos1), int(pos2))
		pos1 = size - 4
		pos2 = index3
		swap(int(pos1), int(pos2))
		pos1 = size - 3
		pos2 = index2
		swap(int(pos1), int(pos2))
		pos1 = size - 2
		pos2 = index1
		swap(int(pos1), int(pos2))
		pos1 = size - 1
		pos2 = index0
		swap(int(pos1), int(pos2))
	}

	// Final small shuffle if anything remains (i <= 6)
	if i > 1 {
		// For remaining i elements, we need to shuffle the last (i-1) positions.
		// 720 = 6! is a safe bound that works for i <= 7.
		// Inline partialShuffle64b(storage, i, i-1, 720, rng)
		var index0, index1, index2, index3, index4 uint64
		randVal := r.Uint64()
		size := i
		k := i - 1
		bound := uint64(720)
		// Unroll loop for k = i-1, but since i <=6, k<=5, we can handle cases
		// For simplicity, since k varies, we'll use a switch based on k
		var a, hi, lo uint64
		switch k {
		case 5:
			a = size - 4
			hi, lo = bits.Mul64(a, randVal)
			randVal = lo
			index4 = hi
			fallthrough
		case 4:
			a = size - 3
			hi, lo = bits.Mul64(a, randVal)
			randVal = lo
			index3 = hi
			fallthrough
		case 3:
			a = size - 2
			hi, lo = bits.Mul64(a, randVal)
			randVal = lo
			index2 = hi
			fallthrough
		case 2:
			a = size - 1
			hi, lo = bits.Mul64(a, randVal)
			randVal = lo
			index1 = hi
			fallthrough
		case 1:
			a = size
			hi, lo = bits.Mul64(a, randVal)
			randVal = lo
			index0 = hi
		}
		if randVal < bound {
			newBound := uint64(1)
			for j := uint64(1); j < k; j++ {
				newBound *= size - j
			}
			t := (-newBound) % newBound
			for randVal < t {
				randVal = r.Uint64()
				switch k {
				case 5:
					a = size - 4
					hi, lo = bits.Mul64(a, randVal)
					randVal = lo
					index4 = hi
					fallthrough
				case 4:
					a = size - 3
					hi, lo = bits.Mul64(a, randVal)
					randVal = lo
					index3 = hi
					fallthrough
				case 3:
					a = size - 2
					hi, lo = bits.Mul64(a, randVal)
					randVal = lo
					index2 = hi
					fallthrough
				case 2:
					a = size - 1
					hi, lo = bits.Mul64(a, randVal)
					randVal = lo
					index1 = hi
					fallthrough
				case 1:
					a = size
					hi, lo = bits.Mul64(a, randVal)
					randVal = lo
					index0 = hi
				}
			}
		}
		// Unroll swaps based on k
		var pos1, pos2 uint64
		switch k {
		case 5:
			pos1 = size - 5
			pos2 = index4
			swap(int(pos1), int(pos2))
			fallthrough
		case 4:
			pos1 = size - 4
			pos2 = index3
			swap(int(pos1), int(pos2))
			fallthrough
		case 3:
			pos1 = size - 3
			pos2 = index2
			swap(int(pos1), int(pos2))
			fallthrough
		case 2:
			pos1 = size - 2
			pos2 = index1
			swap(int(pos1), int(pos2))
			fallthrough
		case 1:
			pos1 = size - 1
			pos2 = index0
			swap(int(pos1), int(pos2))
		}
	}
}

func shuffleBatch23456(storage []int, rng func() uint64) {
	i := uint64(len(storage))

	// Single swaps for sizes > 2^30
	for ; i > (1 << 30); i-- {
		var index0 uint64
		r := rng()
		n := i
		bound := i
		// Unroll loop for k=1
		a := n
		hi, lo := bits.Mul64(a, r)
		r = lo
		index0 = hi
		if r < bound {
			newBound := n
			t := (-newBound) % newBound
			for r < t {
				r = rng()
				a = n
				hi, lo = bits.Mul64(a, r)
				r = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=1
		pos1 := n - 1
		pos2 := index0
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
	}

	// Batches of 2 down to 2^19
	for ; i > (1 << 19); i -= 2 {
		// Inline partialShuffle64b(storage, i, 2, bound, rng)
		var index0, index1 uint64
		r := rng()
		n := i
		bound := uint64(1) << 60
		// Unroll loop for k=2
		a := n - 1
		hi, lo := bits.Mul64(a, r)
		r = lo
		index1 = hi
		a = n
		hi, lo = bits.Mul64(a, r)
		r = lo
		index0 = hi
		if r < bound {
			newBound := n * (n - 1)
			t := (-newBound) % newBound
			for r < t {
				r = rng()
				a = n - 1
				hi, lo = bits.Mul64(a, r)
				r = lo
				index1 = hi
				a = n
				hi, lo = bits.Mul64(a, r)
				r = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=2
		pos1 := n - 2
		pos2 := index1
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 1
		pos2 = index0
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
	}

	// Batches of 3 down to 2^14
	for ; i > (1 << 14); i -= 3 {
		// Inline partialShuffle64b(storage, i, 3, bound, rng)
		var index0, index1, index2 uint64
		r := rng()
		n := i
		bound := uint64(1) << 57
		// Unroll loop for k=3
		a := n - 2
		hi, lo := bits.Mul64(a, r)
		r = lo
		index2 = hi
		a = n - 1
		hi, lo = bits.Mul64(a, r)
		r = lo
		index1 = hi
		a = n
		hi, lo = bits.Mul64(a, r)
		r = lo
		index0 = hi
		if r < bound {
			newBound := n * (n - 1) * (n - 2)
			t := (-newBound) % newBound
			for r < t {
				r = rng()
				a = n - 2
				hi, lo = bits.Mul64(a, r)
				r = lo
				index2 = hi
				a = n - 1
				hi, lo = bits.Mul64(a, r)
				r = lo
				index1 = hi
				a = n
				hi, lo = bits.Mul64(a, r)
				r = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=3
		pos1 := n - 3
		pos2 := index2
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 2
		pos2 = index1
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 1
		pos2 = index0
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
	}

	// Batches of 4 down to 2^11
	for ; i > (1 << 11); i -= 4 {
		// Inline partialShuffle64b(storage, i, 4, bound, rng)
		var index0, index1, index2, index3 uint64
		r := rng()
		n := i
		bound := uint64(1) << 56
		// Unroll loop for k=4
		a := n - 3
		hi, lo := bits.Mul64(a, r)
		r = lo
		index3 = hi
		a = n - 2
		hi, lo = bits.Mul64(a, r)
		r = lo
		index2 = hi
		a = n - 1
		hi, lo = bits.Mul64(a, r)
		r = lo
		index1 = hi
		a = n
		hi, lo = bits.Mul64(a, r)
		r = lo
		index0 = hi
		if r < bound {
			newBound := n * (n - 1) * (n - 2) * (n - 3)
			t := (-newBound) % newBound
			for r < t {
				r = rng()
				a = n - 3
				hi, lo = bits.Mul64(a, r)
				r = lo
				index3 = hi
				a = n - 2
				hi, lo = bits.Mul64(a, r)
				r = lo
				index2 = hi
				a = n - 1
				hi, lo = bits.Mul64(a, r)
				r = lo
				index1 = hi
				a = n
				hi, lo = bits.Mul64(a, r)
				r = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=4
		pos1 := n - 4
		pos2 := index3
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 3
		pos2 = index2
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 2
		pos2 = index1
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 1
		pos2 = index0
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
	}

	// Batches of 5 down to 2^9
	for ; i > (1 << 9); i -= 5 {
		// Inline partialShuffle64b(storage, i, 5, bound, rng)
		var index0, index1, index2, index3, index4 uint64
		r := rng()
		n := i
		bound := uint64(1) << 55
		// Unroll loop for k=5
		a := n - 4
		hi, lo := bits.Mul64(a, r)
		r = lo
		index4 = hi
		a = n - 3
		hi, lo = bits.Mul64(a, r)
		r = lo
		index3 = hi
		a = n - 2
		hi, lo = bits.Mul64(a, r)
		r = lo
		index2 = hi
		a = n - 1
		hi, lo = bits.Mul64(a, r)
		r = lo
		index1 = hi
		a = n
		hi, lo = bits.Mul64(a, r)
		r = lo
		index0 = hi
		if r < bound {
			newBound := n * (n - 1) * (n - 2) * (n - 3) * (n - 4)
			t := (-newBound) % newBound
			for r < t {
				r = rng()
				a = n - 4
				hi, lo = bits.Mul64(a, r)
				r = lo
				index4 = hi
				a = n - 3
				hi, lo = bits.Mul64(a, r)
				r = lo
				index3 = hi
				a = n - 2
				hi, lo = bits.Mul64(a, r)
				r = lo
				index2 = hi
				a = n - 1
				hi, lo = bits.Mul64(a, r)
				r = lo
				index1 = hi
				a = n
				hi, lo = bits.Mul64(a, r)
				r = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=5
		pos1 := n - 5
		pos2 := index4
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 4
		pos2 = index3
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 3
		pos2 = index2
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 2
		pos2 = index1
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 1
		pos2 = index0
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
	}

	// Batches of 6 down to 6
	for ; i > 6; i -= 6 {
		// Inline partialShuffle64b(storage, i, 6, bound, rng)
		var index0, index1, index2, index3, index4, index5 uint64
		r := rng()
		n := i
		bound := uint64(1) << 54
		// Unroll loop for k=6
		a := n - 5
		hi, lo := bits.Mul64(a, r)
		r = lo
		index5 = hi
		a = n - 4
		hi, lo = bits.Mul64(a, r)
		r = lo
		index4 = hi
		a = n - 3
		hi, lo = bits.Mul64(a, r)
		r = lo
		index3 = hi
		a = n - 2
		hi, lo = bits.Mul64(a, r)
		r = lo
		index2 = hi
		a = n - 1
		hi, lo = bits.Mul64(a, r)
		r = lo
		index1 = hi
		a = n
		hi, lo = bits.Mul64(a, r)
		r = lo
		index0 = hi
		if r < bound {
			newBound := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) * (n - 5)
			t := (-newBound) % newBound
			for r < t {
				r = rng()
				a = n - 5
				hi, lo = bits.Mul64(a, r)
				r = lo
				index5 = hi
				a = n - 4
				hi, lo = bits.Mul64(a, r)
				r = lo
				index4 = hi
				a = n - 3
				hi, lo = bits.Mul64(a, r)
				r = lo
				index3 = hi
				a = n - 2
				hi, lo = bits.Mul64(a, r)
				r = lo
				index2 = hi
				a = n - 1
				hi, lo = bits.Mul64(a, r)
				r = lo
				index1 = hi
				a = n
				hi, lo = bits.Mul64(a, r)
				r = lo
				index0 = hi
			}
		}
		// Unroll swaps for k=6
		pos1 := n - 6
		pos2 := index5
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 5
		pos2 = index4
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 4
		pos2 = index3
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 3
		pos2 = index2
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 2
		pos2 = index1
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		pos1 = n - 1
		pos2 = index0
		storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
	}

	// Final small shuffle if anything remains (i <= 6)
	if i > 1 {
		// For remaining i elements, we need to shuffle the last (i-1) positions.
		// 720 = 6! is a safe bound that works for i <= 7.
		// Inline partialShuffle64b(storage, i, i-1, 720, rng)
		var index0, index1, index2, index3, index4 uint64
		r := rng()
		n := i
		k := i - 1
		bound := uint64(720)
		// Unroll loop for k = i-1, but since i <=6, k<=5, we can handle cases
		// For simplicity, since k varies, we'll use a switch based on k
		var a, hi, lo uint64
		switch k {
		case 5:
			a = n - 4
			hi, lo = bits.Mul64(a, r)
			r = lo
			index4 = hi
			fallthrough
		case 4:
			a = n - 3
			hi, lo = bits.Mul64(a, r)
			r = lo
			index3 = hi
			fallthrough
		case 3:
			a = n - 2
			hi, lo = bits.Mul64(a, r)
			r = lo
			index2 = hi
			fallthrough
		case 2:
			a = n - 1
			hi, lo = bits.Mul64(a, r)
			r = lo
			index1 = hi
			fallthrough
		case 1:
			a = n
			hi, lo = bits.Mul64(a, r)
			r = lo
			index0 = hi
		}
		if r < bound {
			newBound := uint64(1)
			for j := uint64(1); j < k; j++ {
				newBound *= n - j
			}
			t := (-newBound) % newBound
			for r < t {
				r = rng()
				switch k {
				case 5:
					a = n - 4
					hi, lo = bits.Mul64(a, r)
					r = lo
					index4 = hi
					fallthrough
				case 4:
					a = n - 3
					hi, lo = bits.Mul64(a, r)
					r = lo
					index3 = hi
					fallthrough
				case 3:
					a = n - 2
					hi, lo = bits.Mul64(a, r)
					r = lo
					index2 = hi
					fallthrough
				case 2:
					a = n - 1
					hi, lo = bits.Mul64(a, r)
					r = lo
					index1 = hi
					fallthrough
				case 1:
					a = n
					hi, lo = bits.Mul64(a, r)
					r = lo
					index0 = hi
				}
			}
		}
		// Unroll swaps based on k
		var pos1, pos2 uint64
		switch k {
		case 5:
			pos1 = n - 5
			pos2 = index4
			storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
			fallthrough
		case 4:
			pos1 = n - 4
			pos2 = index3
			storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
			fallthrough
		case 3:
			pos1 = n - 3
			pos2 = index2
			storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
			fallthrough
		case 2:
			pos1 = n - 2
			pos2 = index1
			storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
			fallthrough
		case 1:
			pos1 = n - 1
			pos2 = index0
			storage[pos1], storage[pos2] = storage[pos2], storage[pos1]
		}
		// bound not used
	}
}
