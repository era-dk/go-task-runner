package runner

import (
	"fmt"
	"strings"
)

type Lines []string

func (l *Lines) Add(offset int, spinner string, text string, status string) {
    *l = append(*l, fmt.Sprintf(
        "%s%s %s %s",
        strings.Repeat(" ", offset),
        spinner,
        text,
        status,
    ))
}