package pooling

import (
	"bytes"
	"encoding/json"
	"sync"
)

type BufferPool struct {
	pool sync.Pool
}

func NewBufferPool() *BufferPool {
	bp := &BufferPool{}
	bp.pool.New = func() any {
		return &bytes.Buffer{}
	}
	return bp
}

func (p *BufferPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

func (p *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	p.pool.Put(b)
}

func WriteJSON(p *BufferPool, v any) ([]byte, error) {
	buf := p.Get()
	defer p.Put(buf)

	enc := json.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	out := make([]byte, buf.Len())
	copy(out, buf.Bytes())
	return out, nil
}
