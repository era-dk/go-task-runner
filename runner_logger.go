package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (r *Runner) setupLogger() error {
    zerolog.SetGlobalLevel(zerolog.Disabled)
    if r.logfile != "" {
        resolvedLogPath, err := filepath.Abs(r.logfile)
        if err != nil {
            return err
        }

        logStream, err := os.OpenFile(resolvedLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
        if err != nil {
            return err
        }

        zerolog.SetGlobalLevel(zerolog.DebugLevel)
        log.Logger = log.Output(zerolog.ConsoleWriter{
            Out: logStream,
            NoColor: true,
            TimeFormat: time.RFC3339,
            FormatLevel: func(i interface{}) string {
                return strings.ToUpper(fmt.Sprintf("[%s]", i))
            },
        })
    }
    return nil
}