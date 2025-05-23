package runner

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"testing"
	"time"
)


var counterFnResolver = func (counter int) Resolver {
	return func(ctx context.Context, task *Task) error {
		for i := range counter {
			time.Sleep(200 * time.Millisecond)
			task.Msg(fmt.Sprintf("fn-counter[1-%d] - %d", counter, i+1))
		}
		return nil
	}
}

var counterBashResolver = func (counter int) Resolver {
	return func(ctx context.Context, task *Task) error {
		execCmd := exec.CommandContext(
			ctx,
			"/bin/sh",
			"-c",
			fmt.Sprintf("for i in `seq %d`; do echo \"bash-counter[1-%d] - $i\"; sleep 1; done;", counter, counter),
		)
		execCmd.Stdout = task.Output
		if err := execCmd.Run(); err != nil {
			return err
		}
		return nil
	}
}

var mainTask = &Task{
	Hidden: true,
	Subtasks: []*Task{
		{
			Title: "Configure task",
			OutputLines: 1,
			Resolver: func (ctx context.Context, task *Task) error {
				task.Log().Info().Msg("log configure task")
				return nil
			},
		},
		{
			Title: "Validate",
			Collapse: true,
			SubtasksConcurrent: true,
			SkipOnFail: true,
			Resolver: func(ctx context.Context, task *Task) error {
				return errors.New("my exception")
			},
			Subtasks: []*Task{
				{
					Title: "Validate task 1",
					OutputLines: 1,
					Resolver: func(ctx context.Context, task *Task) error {
						return nil
					},
				},
				{
					Title: "Validate task 2",
					SubtasksConcurrent: true,
					Subtasks: []*Task{
						{
							Title: "Validate task 2 - counter A",
							OutputLines: 1,
							Collapse: true,
							Resolver: counterBashResolver(14),
						},
						{
							Title: "Validate task 2 - counter B",
							OutputLines: 1,
							Resolver: counterFnResolver(9),
						},
					},
				},
			},
		},
		{
			Title: "Counter tasks",
			Collapse: true,
			SubtasksConcurrent: true,
			Subtasks: []*Task{
				{
					Title: "Counter task 1",
					OutputLines: 2,
					Resolver: counterFnResolver(15),
				},
				{
					Title: "Counter task 2",
					OutputLines: 5,
					Resolver: counterFnResolver(10),
				},
			},
		},
		{
			Title: "Complete task",
			Subtasks: []*Task{
				{
					Title: "Quick counter task",
					OutputLines: 1,
					Resolver: counterFnResolver(1),
				},
				{
					Title: "An error task",
					Resolver: func(ctx context.Context, task *Task) error {
						return nil
						//return errors.New("it's an error exception")
					},
				},
			},
		},
	},
}

func TestExample(t *testing.T) {
	if err := NewRunner(mainTask).Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}