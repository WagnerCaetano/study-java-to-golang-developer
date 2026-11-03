# Exercise 59: Pooling and Reuse

## Problem

High-throughput JSON APIs allocate heavily if every response builds a new buffer. Reuse `bytes.Buffer` values with `sync.Pool`.

Package **`pooling`**:

```go
type BufferPool struct { ... }

func NewBufferPool() *BufferPool
func (p *BufferPool) Get() *bytes.Buffer
func (p *BufferPool) Put(b *bytes.Buffer) // Reset before returning to pool

func WriteJSON(p *BufferPool, v any) ([]byte, error)
```

Rules:

- `BufferPool` uses **`sync.Pool`** internally.
- `Put` must **`Reset`** the buffer before storing it.
- `WriteJSON` gets a buffer, encodes `v` with `encoding/json`, **copies** bytes out (caller owns the slice), then puts the buffer back.

## Learning objectives

- Reducing allocation pressure in hot paths.
- Safe reuse with `Reset` and returning copied bytes.
- When `sync.Pool` is appropriate (and when it is not).

## How to test

```bash
go test -v
```
