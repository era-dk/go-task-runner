package runner

import (
	"context"

	"github.com/rs/zerolog/log"
)

func NewRunner(task *Task) *Runner {
    return &Runner{
        task: task,
        progressWriter: NewTtyWriter(),
    }
}

type Runner struct {
    logfile string
    task *Task
    noProgress bool
    progressWriter *ttyWriter
}

func (r *Runner) UseLogger(logfile string) *Runner {
    r.logfile = logfile
    return r
}

func (r *Runner) NoProgress() *Runner {
    r.noProgress = true
    return r
}

func (r *Runner) Run(ctx context.Context, params ParamsInterface) error {
    if err := r.setupLogger(); err != nil {
        return err
    }

    log.Info().Msg("start runner")
    if err := r.task.Setup(ctx); err != nil {
        return err
    }

    ctx, cancel := context.WithCancel(ctx)
    go func() {
        if err := r.task.Run(ctx, &params); err != nil {
            r.task.state = TaskStateError
        }
        cancel()
    }()

    if r.noProgress {
        r.listenNoProgress(ctx)
    } else {
        r.listenProgress(ctx)
    }
    
    log.Info().Msg("runner completed")
    return nil
}