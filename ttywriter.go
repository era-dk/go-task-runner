package runner

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/morikuni/aec"
)

func NewTtyWriter() *ttyWriter {
    return &ttyWriter{
        mtx: &sync.Mutex{},
        out: io.Writer(os.Stdout),
    }
}

type ttyWriter struct {
    mtx *sync.Mutex
    out io.Writer
    numLines uint
}

func (w *ttyWriter) Start() {
    fmt.Fprint(w.out, aec.Hide)
}

func (w *ttyWriter) End() {
    fmt.Fprint(w.out, aec.Show)
}

func (w *ttyWriter) PrintLines(lines Lines) {
    w.mtx.Lock()
    defer w.mtx.Unlock()

    b := aec.EmptyBuilder
    b = b.Up(w.numLines)
    fmt.Fprint(w.out, b.Column(0).ANSI)

    var numLines uint = 0
    for _, line := range lines {
        fmt.Fprint(w.out, aec.EraseLine(aec.EraseModes.All))
        fmt.Fprintln(w.out, line)
        numLines++
    }

    if w.numLines > numLines {
        for i := numLines; i < w.numLines; i++ {
            fmt.Fprintln(w.out, aec.EraseLine(aec.EraseModes.All))
        }
        b = aec.EmptyBuilder
        b = b.Up(w.numLines - numLines)
        fmt.Fprint(w.out, b.Column(0).ANSI)
    }

    w.numLines = numLines
}

