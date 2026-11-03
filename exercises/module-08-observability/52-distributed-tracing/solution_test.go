package ddtraceex

import "testing"

func TestRecordingTracerRecordsSpan(t *testing.T) {
	var rt RecordingTracer
	finish := rt.StartSpan("http.request")
	finish()
	if len(rt.Spans) != 1 || rt.Spans[0] != "http.request" {
		t.Fatalf("spans = %v, want [http.request]", rt.Spans)
	}
}

func TestDefaultStartSpanUsesNoop(t *testing.T) {
	SetDefaultTracer(NoopTracer{})
	finish := StartSpan("noop")
	finish() // should not panic
}

func TestDefaultStartSpanWithRecording(t *testing.T) {
	rt := &RecordingTracer{}
	SetDefaultTracer(rt)
	finish := StartSpan("bff.aggregate")
	finish()
	if len(rt.Spans) != 1 || rt.Spans[0] != "bff.aggregate" {
		t.Fatalf("spans = %v", rt.Spans)
	}
	SetDefaultTracer(NoopTracer{})
}
