package runner

import (
	"context"
	"strings"
)

func (t *Task) Progress(ctx context.Context, level int) Lines {
    var lines Lines

    collapse := t.state == TaskStateCompleted && t.Collapse
    if !t.Hidden {
        lines.Add(level, t.Spinner(), ApplyStyle(StyleTitle, t.Title))
        if t.OutputLines > 0 && t.state == TaskStateProgress && !collapse {
            output := strings.Trim(t.Output.String(), "\n")
            if output != "" {
                splits := strings.Split(output, "\n")
                if len(splits) > t.OutputLines {
                    splits = splits[(len(splits) - t.OutputLines):]
                }
                for _, s := range splits {
                    lines.Add(level + 2, "", ApplyStyle(StyleMessage, s))
                }
            }
        }
        if t.err != nil {
            lines.Add(level + 2, "", ApplyStyle(StyleError, t.err.Error()))
        }
    }

    if !collapse {
        for _, subtask := range t.Subtasks {
            lines = append(lines, subtask.Progress(ctx, level + 1)...)
        }
    }

    return lines
}