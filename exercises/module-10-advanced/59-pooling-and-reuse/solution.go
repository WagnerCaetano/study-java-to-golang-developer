package pooling

import "bytes"

type BufferPool struct {
	// FIX ME: wrap sync.Pool
}

func NewBufferPool() *BufferPool {
	return &BufferPool{}
}

func (p *BufferPool) Get() *bytes.Buffer {
	// FIX ME
	return &bytes.Buffer{}
}

func (p *BufferPool) Put(b *bytes.Buffer) {
	// FIX ME: Reset before put
}

func WriteJSON(p *BufferPool, v any) ([]byte, error) {
	// FIX ME: get buffer, json encode, copy bytes, put buffer
	return nil, nil
}
