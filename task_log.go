package runner

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (t *Task) Log() *zerolog.Logger {
    return &log.Logger
}