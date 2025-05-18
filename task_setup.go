package runner

import (
	"bytes"
	"context"
)

func (t *Task) Setup(ctx context.Context) error {
    t.Output = new(bytes.Buffer)

    for _, subtask := range t.Subtasks {
        if err := subtask.Setup(ctx); err != nil {
            return err
        }
    }

    return nil
}