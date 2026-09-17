package vt10x

import (
	"bytes"
	"testing"
)

func TestCSIParse(t *testing.T) {
	var csi csiEscape
	csi.reset()
	csi.buf = []byte("s")
	csi.parse()
	if csi.mode != 's' || csi.arg(0, 17) != 17 || len(csi.args) != 0 {
		t.Fatal("CSI parse mismatch")
	}

	csi.reset()
	csi.buf = []byte("31T")
	csi.parse()
	if csi.mode != 'T' || csi.arg(0, 0) != 31 || len(csi.args) != 1 {
		t.Fatal("CSI parse mismatch")
	}

	csi.reset()
	csi.buf = []byte("48;2f")
	csi.parse()
	if csi.mode != 'f' || csi.arg(0, 0) != 48 || csi.arg(1, 0) != 2 || len(csi.args) != 2 {
		t.Fatal("CSI parse mismatch")
	}

	csi.reset()
	csi.buf = []byte("?25l")
	csi.parse()
	if csi.mode != 'l' || csi.arg(0, 0) != 25 || csi.priv != true || len(csi.args) != 1 {
		t.Fatal("CSI parse mismatch")
	}
}

func TestEraseScrollback(t *testing.T) {
	term := New(WithSize(10, 3), WithScrollback(100))
	term.Write([]byte("\033[20h")) // set CRLF mode

	// Push enough lines to overflow the 3-row screen into scrollback.
	for i := 0; i < 10; i++ {
		term.Write([]byte("line\n"))
	}
	if term.ScrollbackLen() == 0 {
		t.Fatal("expected scrollback to be populated")
	}

	// ESC[3J (xterm E3) erases saved lines but not the visible screen.
	term.Write([]byte("visible"))
	term.Write([]byte("\033[3J"))

	if got := term.ScrollbackLen(); got != 0 {
		t.Fatalf("expected empty scrollback after ESC[3J, got %d lines", got)
	}
	if got := extractStr(term, 0, 6, term.Cursor().Y); got != "visible" {
		t.Fatalf("ESC[3J must not clear the screen, got %q", got)
	}

	// Scrollback must keep working after the reset.
	for i := 0; i < 10; i++ {
		term.Write([]byte("more\n"))
	}
	if term.ScrollbackLen() == 0 {
		t.Fatal("expected scrollback to repopulate after ESC[3J")
	}
	if line := term.ScrollbackLine(0); line == nil {
		t.Fatal("expected ScrollbackLine(0) to be readable after repopulate")
	}
}

func TestDeviceAttributes(t *testing.T) {
	var buf bytes.Buffer
	term := New(WithWriter(&buf))
	term.Write([]byte("\033[c"))
	if got := buf.String(); got != "\033[?6c" {
		t.Fatalf("expected \\033[?6c, got %q", got)
	}

	buf.Reset()
	term.Write([]byte("\033[0c"))
	if got := buf.String(); got != "\033[?6c" {
		t.Fatalf("expected \\033[?6c, got %q", got)
	}
}
