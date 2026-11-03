package ddtraceex

type Tracer interface {
	StartSpan(name string) (finish func())
}

type RecordingTracer struct {
	Spans []string
}

func (t *RecordingTracer) StartSpan(name string) (finish func()) {
	t.Spans = append(t.Spans, name)
	return func() {}
}

type NoopTracer struct{}

func (NoopTracer) StartSpan(name string) (finish func()) {
	return func() {}
}

var defaultTracer Tracer = NoopTracer{}

func SetDefaultTracer(t Tracer) {
	defaultTracer = t
}

func StartSpan(name string) (finish func()) {
	return defaultTracer.StartSpan(name)
}
