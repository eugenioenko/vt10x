package vt10x

import (
	"io"
	"testing"
)

// TestAltScreenScrollDoesNotLeakToScrollback proves, independent of any
// specific CLI's behavior, that content scrolled off within the alternate
// screen buffer (DECSET 1049) must never be recorded to scrollback -- real
// terminals discard it entirely when the app exits the alt screen, since
// it's transient UI, not shell history. The same scroll performed in the
// primary screen must still be recorded normally.
func TestAltScreenScrollDoesNotLeakToScrollback(t *testing.T) {
	term := New(WithSize(80, 10), WithScrollback(100))

	write := func(s string) {
		if _, err := term.Write([]byte(s)); err != nil && err != io.EOF {
			t.Fatal(err)
		}
	}

	// Enter alt screen, pin a 5-row scroll region, and scroll well past it.
	write("\x1b[?1049h")
	write("\x1b[1;5r")
	for i := 0; i < 20; i++ {
		write("line\r\n")
	}
	if got := term.ScrollbackLen(); got != 0 {
		t.Errorf("alt screen: ScrollbackLen() = %d, want 0 (alt screen must never leak to scrollback)", got)
	}

	// Exit alt screen; the same scroll in the primary screen must behave normally.
	write("\x1b[?1049l")
	write("\x1b[1;5r")
	for i := 0; i < 20; i++ {
		write("line\r\n")
	}
	if got := term.ScrollbackLen(); got == 0 {
		t.Errorf("primary screen: ScrollbackLen() = 0, want > 0 (primary screen scroll must still be recorded)")
	}
}
