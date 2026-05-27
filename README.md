# vt10x

Fork of [hinshun/vt10x](https://github.com/hinshun/vt10x) maintained for use in [TTT](https://github.com/eugenioenko/ttt), a terminal text editor.

## Why a fork?

The upstream repository has been inactive for over 10 years. This fork adds features needed by TTT's integrated terminal emulator:

- **`SetOnScrollOut` callback** — fires when lines scroll off the top of the terminal, enabling host applications to capture scrollback history for their own scroll buffer.

## Usage

```go
import "github.com/eugenioenko/vt10x"
```

## Original

Package vt10x is a vt10x terminal emulation backend, influenced largely by st, rxvt, xterm, and iTerm as reference. Use it for terminal muxing, a terminal emulation frontend, or wherever else you need terminal emulation.

Original author: [hinshun](https://github.com/hinshun)
