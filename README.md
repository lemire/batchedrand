# batchedrand

This package provides a batched random shuffle implementation for Go. It implements the batched algorithm
from Brackett-Rozinsky and Lemire (2025)---a mathematically correct shuffling. It should be between
2x and 3x faster than the standard Go library (`"math/rand/v2"`) while producing the same high-quality 
results.


- [Batched Ranged Random Integer Generation](https://arxiv.org/pdf/2408.06213), Software: Practice and Experience 55 (1), 2025

## Usage

To create a new `BatchedRand` instance, you may choose between ChaCha8 and PCG.

To use ChaCha8 RNG:

```go
rng := batchedrand.Rand{rand.New(rand.NewChaCha8([32]byte{1, 2, 3, /* ... */}))}
```

To use PCG RNG:

```go
rng := batchedrand.Rand{rand.New(rand.NewPCG(1, 2))}
```

To shuffle a slice using the batched shuffle:

```go
data := []int{1, 2, 3, 4, 5}
rng.Shuffle(len(data), func(i, j int) {
    data[i], data[j] = data[j], data[i]
})
```


## Running Tests

To run the tests:

```bash
go test
```

## Running Benchmarks

To run the benchmarks:

```bash
go test -bench=.
```

This will execute all benchmark functions in the package.

If you have python and uv, you can run...

```bash
uv run --with=pandas --with=matplotlib --with=tabulate bench.p
```

Result on an M4 processor:


| Benchmark     |   Size |   Batched (ns/item) |   Standard (ns/item) |   speedup |
|:--------------|-------:|--------------------:|---------------------:|----------:|
| ChaChaShuffle |     30 |                 1.8 |                  4.6 |       2.6 |
| ChaChaShuffle |    100 |                 1.8 |                  4.7 |       2.5 |
| ChaChaShuffle | 500000 |                 2.6 |                  5.1 |       1.9 |
| PCGShuffle    |     30 |                 1.5 |                  3.9 |       2.6 |
| PCGShuffle    |    100 |                 1.5 |                  4.2 |       2.8 |
| PCGShuffle    | 500000 |                 1.9 |                  3.8 |       2   |
