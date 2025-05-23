package runner

import (
	"context"
	"time"
)

func (r *Runner) listenNoProgress(ctx context.Context) {
    loop:
    for {
        select {
        case <-ctx.Done():
            break loop
        }
    }
}

func (r *Runner) listenProgress(ctx context.Context) {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()

    r.progressWriter.Start()
    defer r.progressWriter.End()

    loop:
    for {
        select {
        case <-ticker.C:
            lines := r.task.Progress(ctx, 0)
            r.progressWriter.PrintLines(lines)
        case <-ctx.Done():
            lines := r.task.Progress(ctx, 0)
            r.progressWriter.PrintLines(lines)
            break loop
        }
    }
}