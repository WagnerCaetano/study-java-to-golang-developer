package ratelimit

import (
	"testing"
	"time"
)

func TestAllowConsumesCapacity(t *testing.T) {
	lim := NewLimiter(2, time.Hour)

	if !lim.Allow() {
		t.Fatal("first Allow() = false, want true")
	}
	if !lim.Allow() {
		t.Fatal("second Allow() = false, want true")
	}
	if lim.Allow() {
		t.Fatal("third Allow() = true, want false")
	}
}

func TestAllowRefillsAfterWait(t *testing.T) {
	lim := NewLimiter(1, 20*time.Millisecond)

	if !lim.Allow() {
		t.Fatal("first Allow() = false, want true")
	}
	if lim.Allow() {
		t.Fatal("second Allow() = true, want false (bucket empty)")
	}

	time.Sleep(25 * time.Millisecond)

	if !lim.Allow() {
		t.Fatal("Allow() after refill = false, want true")
	}
}
