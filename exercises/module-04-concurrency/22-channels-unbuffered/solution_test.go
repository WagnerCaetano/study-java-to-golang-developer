package channel

import "testing"

func TestProducerConsumer(t *testing.T) {
	ch := Producer(10)
	result := Consumer(ch)
	expected := 55 // 1+2+...+10
	if result != expected {
		t.Errorf("Consumer got %d, want %d", result, expected)
	}
}

func TestProducerConsumerEmpty(t *testing.T) {
	ch := Producer(0)
	result := Consumer(ch)
	if result != 0 {
		t.Errorf("Empty producer: got %d, want 0", result)
	}
}
