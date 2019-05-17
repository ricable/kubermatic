package log

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	ctrlruntimelog "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

// Options exports a options struct to be used by cmd's
type Options struct {
	// Enable debug logs
	Debug bool
	// Log format (JSON or plain text)
	Format string
}

func (o *Options) Validate() error {
	if !AvailableFormats.Contains(o.Format) {
		return fmt.Errorf("invalid log-format specified %q; available: %s", o.Format, AvailableFormats.String())
	}
	return nil
}

type Format string

type Formats []Format

const (
	FormatJSON    Format = "JSON"
	FormatConsole Format = "Console"
)

var (
	AvailableFormats = Formats{FormatJSON, FormatConsole}
)

func (f Formats) String() string {
	const separator = ", "
	var s string
	for _, format := range f {
		s = s + separator + string(format)
	}
	return strings.TrimPrefix(s, separator)
}

func (f Formats) Contains(s string) bool {
	for _, format := range f {
		if s == string(format) {
			return true
		}
	}
	return false
}

func New(debug bool, format Format) *zap.Logger {
	// this basically mimics New<type>Config, but with a custom sink
	sink := zapcore.AddSync(os.Stderr)

	// Level - We only support setting Info+ or Debug+
	lvl := zap.NewAtomicLevelAt(zap.InfoLevel)
	if debug {
		lvl = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	encCfg := zap.NewProductionEncoderConfig()
	// Having a dateformat makes it more easy to look at logs outside of something like Kibana
	encCfg.TimeKey = "time"
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var enc zapcore.Encoder
	if format == FormatJSON {
		enc = zapcore.NewJSONEncoder(encCfg)
	} else if format == FormatConsole {
		enc = zapcore.NewConsoleEncoder(encCfg)
	}

	opts := []zap.Option{
		zap.AddCaller(),
		zap.ErrorOutput(sink),
	}

	coreLog := zapcore.NewCore(&ctrlruntimelog.KubeAwareEncoder{Encoder: enc}, sink, lvl)
	return zap.New(coreLog, opts...)
}