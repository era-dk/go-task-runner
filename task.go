package runner

import (
	"bytes"
)

type TaskState int
const (
    TaskStateIdle TaskState = iota
    TaskStateProgress
    TaskStateCompleted
    TaskStateError
)

type Task struct {    
    Title string
    Resolver Resolver
    Hidden bool
    Collapse bool
    SkipOnFail bool
    OutputLines int
    SubtasksConcurrent bool
    Subtasks []*Task

    Output *bytes.Buffer
    state TaskState
    spinnerIndex int
    err error
}