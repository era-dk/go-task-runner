package runner

import (
	"fmt"

	"github.com/morikuni/aec"
)

var StatusLabels = map[TaskState]string{
	TaskStateProgress: "in progress",
	TaskStateError: "error",
}

var StatusStyles = map[TaskState]aec.ANSI{
	TaskStateProgress: aec.YellowF,
	TaskStateCompleted: aec.GreenF,
	TaskStateError: aec.RedF,
}

func (t *Task) Status() string {
	label, ok := StatusLabels[t.state]
	if ok {
		label = fmt.Sprintf("[%s]", label)
		statusStyle, ok := StatusStyles[t.state]
		if ok {
			return statusStyle.Apply(label)
		}
		return label
	}
	return ""
}