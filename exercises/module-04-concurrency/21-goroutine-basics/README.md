# Exercise 21: Goroutine Basics

## Problem
Fix `ParallelSum` so it waits for every goroutine and combines all of their
partial sums. The current implementation returns after receiving only one
goroutine's result.

The function should return `0` for an empty slice. If `numGoroutines` is zero
or negative, treat it as one. It should also work when there are more requested
goroutines than numbers.

## How to Test
```bash
go test -v
```
