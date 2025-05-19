package runner

import (
	"github.com/morikuni/aec"
)

var SpinnerIcons = map[TaskState][]string{
	TaskStateIdle: {" "},
	TaskStateProgress: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	TaskStateCompleted: {"✔"},
	TaskStateError: {"✘"},
}

var SpinnerStyles = map[TaskState]aec.ANSI{
	TaskStateProgress: aec.YellowF,
	TaskStateCompleted: aec.GreenF,
	TaskStateError: aec.RedF,
}

func (t *Task) Spinner() string {
	icons, ok := SpinnerIcons[t.state]
	if ok {
		t.spinnerIndex++
		if t.spinnerIndex >= len(icons) {
			t.spinnerIndex = 0
		}
		spinStyle, ok := SpinnerStyles[t.state]
		if ok {
			return spinStyle.Apply(icons[t.spinnerIndex])
		}
		return icons[t.spinnerIndex]
	}
	return ""
}