package runner

import "fmt"

func (t *Task) Status() string {
	if t.err != nil {
		return ProgressStyleError.Apply(fmt.Sprintf("-> %s", t.err.Error()))
	}
	return ""
}