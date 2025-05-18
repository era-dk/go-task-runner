package runner

import "bytes"

type TaskState int

const (
    TaskStateIdle TaskState = iota
    TaskStateProgress
    TaskStateNoticed
    TaskStateCompleted
    TaskStateError
)

type Task struct {    
    Title string
    Resolver Resolver
    Hidden bool
    Collapse bool
    OutputLines int
    SubtasksConcurrent bool
    Subtasks []*Task

    Output *bytes.Buffer
    state TaskState
    spinnerIndex int
    notice error
    err error
}