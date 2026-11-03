package ddtraceex

// Tracer starts spans for distributed tracing.
type Tracer interface {
	StartSpan(name string) (finish func())
}

// RecordingTracer records span names for tests.
type RecordingTracer struct {
	Spans []string
}

func (t *RecordingTracer) StartSpan(name string) (finish func()) {
	// FIX ME: append name to t.Spans
	return func() {}
}

// NoopTracer discards spans.
type NoopTracer struct{}

func (NoopTracer) StartSpan(name string) (finish func()) {
	return func() {}
}

var defaultTracer Tracer = NoopTracer{}

// SetDefaultTracer sets the tracer used by StartSpan.
func SetDefaultTracer(t Tracer) {
	defaultTracer = t
}

// StartSpan starts a span on the default tracer.
func StartSpan(name string) (finish func()) {
	return defaultTracer.StartSpan(name)
}
