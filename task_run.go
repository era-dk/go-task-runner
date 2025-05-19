package runner

import (
	"context"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func (t *Task) Run(ctx context.Context, params *ParamsInterface) error {
    t.state = TaskStateProgress
    if t.Resolver != nil {
        log.Info().
            Str("title", t.Title).
            Msg("resolve task")
        if err := t.Resolver(ctx, t, params); err != nil {
            log.Error().
                Err(err).
                Str("title", t.Title).
                Msg("task failed")

            t.err = err
            t.state = TaskStateError
            if t.SkipOnFail {
                return nil
            }
            return err
        }
    }

    if t.SubtasksConcurrent {
        eg, egCtx := errgroup.WithContext(ctx)
        for _, subtask := range t.Subtasks {
            eg.Go(func() error {
                if err := subtask.Run(egCtx, params); err != nil {
                    subtask.state = TaskStateError
                    return err
                }
                return nil
            })
        }
        if err := eg.Wait(); err != nil {
            t.state = TaskStateError
            return err
        }
    } else {
        for _, subtask := range t.Subtasks {
            if err := subtask.Run(ctx, params); err != nil {
                subtask.state = TaskStateError
                return err
            }
        }
    }

    t.state = TaskStateCompleted
    return nil
}