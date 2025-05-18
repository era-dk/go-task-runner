package runner

import (
	"fmt"
)

func (t *Task) Msg(s string) {
    t.Output.WriteString(fmt.Sprintf("%s\n", s))
}