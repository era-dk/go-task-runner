package runner

import (
	"context"
	"time"
)

func (r *Runner) listen(ctx context.Context) {
    ticker := time.NewTicker(100 * time.Millisecond)
    if r.noProgress {
        ticker.Stop()
    } else {
        defer ticker.Stop()
    }

    r.progressWriter.Start()
    defer r.progressWriter.End()

    loop:
    for {
        select {
        case <-ticker.C:
            lines := r.task.Progress(ctx, 0)
            r.progressWriter.PrintLines(lines)
        case <-ctx.Done():
            if !r.noProgress {
                lines := r.task.Progress(ctx, 0)
                r.progressWriter.PrintLines(lines)
            }
            break loop
        }
    }
}