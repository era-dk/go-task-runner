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

type MyParams struct {
	Key1 string
	Key2 string
}

var counterFnResolver = func (counter int) Resolver {
	return func(ctx context.Context, task *Task, params *ParamsInterface) error {
		for i := range counter {
			time.Sleep(200 * time.Millisecond)
			task.Msg(fmt.Sprintf("fn-counter[1-%d] - %d", counter, i+1))
		}
		return nil
	}
}

var counterBashResolver = func (counter int) Resolver {
	return func(ctx context.Context, task *Task, params *ParamsInterface) error {
		execCmd := exec.Command(
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
	Hidden: false,
	Subtasks: []*Task{
		{
			Title: "Configure task",
			OutputLines: 1,
			Resolver: func (ctx context.Context, task *Task, params *ParamsInterface) error {
				task.Log().Info().
					Msg("log configure task")
				myParams := (*params).(MyParams)
				myParams.Key2 = "value2"
				*params = myParams

				task.Msg(fmt.Sprintf("param key 1: %s", myParams.Key1))
				return nil
			},
		},
		{
			Title: "Validate",
			Collapse: true,
			SubtasksConcurrent: true,
			SkipOnFail: true,
			Resolver: func(ctx context.Context, task *Task, params *ParamsInterface) error {
				return errors.New("exception")
			},
			Subtasks: []*Task{
				{
					Title: "Validate task 1",
					OutputLines: 1,
					Resolver: func(ctx context.Context, task *Task, params *ParamsInterface) error {
						myParams := (*params).(MyParams)
						task.Msg(fmt.Sprintf("param key 2: %s", myParams.Key2))
	
						return nil
					},
				},
				{
					Title: "Validate task 2",
					SubtasksConcurrent: true,
					Subtasks: []*Task{
						{
							Title: "Validate task 2 - counter A",
							OutputLines: 2,
							Collapse: true,
							Resolver: counterFnResolver(4),
						},
						{
							Title: "Validate task 2 - counter B",
							OutputLines: 2,
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
					Resolver: func(ctx context.Context, task *Task, params *ParamsInterface) error {
						return nil
						//return errors.New("it's an error exception")
					},
				},
			},
		},
	},
}

func TestExample(t *testing.T) {
	if err := NewRunner(mainTask).Run(context.Background(), MyParams{
		Key1: "value 1",
	}); err != nil {
		log.Fatal(err)
	}
}