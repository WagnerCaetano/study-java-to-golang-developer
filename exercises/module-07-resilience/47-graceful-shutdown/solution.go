package graceful

import "context"

// Server coordinates background work and shutdown.
type Server struct {
	// FIX ME: track context, cancel func, and WaitGroup
}

func NewServer() *Server {
	return &Server{}
}

// StartBackgroundWork runs fn in a goroutine until the server context is cancelled.
// FIX ME: Register work with WaitGroup and pass a cancellable context.
func (s *Server) StartBackgroundWork(fn func(ctx context.Context)) {
	go fn(context.Background())
}

// Shutdown stops background work and waits for it to finish, or returns when ctx is done.
// FIX ME: Cancel work and wait on WaitGroup (respect shutdown ctx).
func (s *Server) Shutdown(ctx context.Context) error {
	_ = ctx
	return nil
}
