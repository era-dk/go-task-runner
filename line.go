package runner

import (
	"fmt"
	"strings"
)

type Lines []string

func (l *Lines) Add(offset int, text string) {
    *l = append(*l, fmt.Sprintf("%s%s", strings.Repeat(" ", offset), text))
}