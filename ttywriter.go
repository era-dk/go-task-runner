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

func (w *ttyWriter) HideCursor() {
    fmt.Fprint(w.out, aec.Hide)
}

func (w *ttyWriter) ShowCursor() {
    fmt.Fprint(w.out, aec.Show)
}

func (w *ttyWriter) PrintLines(lines Lines, clearLastLines bool) {
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

    if clearLastLines {
        for i := numLines; i < w.numLines; i++ {
            fmt.Fprintln(w.out, aec.EraseLine(aec.EraseModes.All))
            numLines++
        }
    }

    w.numLines = numLines
}

