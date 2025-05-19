package runner

import (
	"context"
	"fmt"
	"strings"

	"github.com/morikuni/aec"
)

var (
    ProgressStyleTitle = aec.CyanF
    ProgressStyleMessage = aec.Color8BitF(aec.NewRGB8Bit(132, 132, 132))
    ProgressStyleError = aec.RedF
)

func (t *Task) Progress(ctx context.Context, level int) Lines {
    var lines Lines

    collapse := t.state == TaskStateCompleted && t.Collapse
    if !t.Hidden {
        lines.Add(level, fmt.Sprintf("%s %s %s", t.Spinner(), ProgressStyleTitle.Apply(t.Title), t.Status()))
        if t.OutputLines > 0 && t.state == TaskStateProgress && !collapse {
            output := strings.Trim(t.Output.String(), "\n")
            if output != "" {
                splits := strings.Split(output, "\n")
                if len(splits) > t.OutputLines {
                    splits = splits[(len(splits) - t.OutputLines):]
                }
                for _, s := range splits {
                    lines.Add(level + 3, ProgressStyleMessage.Apply(s))
                }
            }
        }
    }

    if !collapse {
        for _, subtask := range t.Subtasks {
            lines = append(lines, subtask.Progress(ctx, level + 1)...)
        }
    }

    return lines
}