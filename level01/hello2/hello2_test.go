package hello2

import "testing"

func TestHello2(t *testing.T) {
    want := "Hello, world."
    if got := Hello2(); got != want {
        t.Errorf("Hello2() = %q, want %q", got, want)
    }
}