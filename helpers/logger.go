package helpers

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type PrependEncoder struct {
	zapcore.Encoder
	Pool buffer.Pool
}

func Log() *zap.Logger {
	production := true
	cfg := ZapGetConfig(production)

	encoder := &PrependEncoder{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		Pool:    buffer.NewPool(),
	}

	logger := zap.New(
		zapcore.NewCore(
			encoder,
			os.Stdout,
			zapcore.DebugLevel,
		),

		zap.ErrorOutput(os.Stderr),
		zap.AddCaller(),
	)

	return logger
}

func ZapGetConfig(production bool) zap.Config {
	var config zap.Config

	if production {
		config = zap.NewProductionConfig()
		config.Encoding = "console"
		config.EncoderConfig.TimeKey = "timestamp"
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.DisableStacktrace = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{"stdout"}
	timeEncoder := func(t time.Time, e zapcore.PrimitiveArrayEncoder) {
		e.AppendString(time.Now().Format("2006-01-02 15:04:05"))
	}
	config.EncoderConfig.EncodeTime = timeEncoder

	return config
}

func (e *PrependEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf := e.Pool.Get()
	buf.AppendString(e.toJournaldPrefix(entry.Level))
	buf.AppendString(" ")

	consolebuf, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(consolebuf.Bytes())
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (e *PrependEncoder) toJournaldPrefix(lvl zapcore.Level) string {
	switch lvl {
	case zapcore.DebugLevel:
		return "<-1>"
	case zapcore.InfoLevel:
		return "<0>"
	case zapcore.WarnLevel:
		return "<1>"
	case zapcore.ErrorLevel:
		return "<2>"
	case zapcore.DPanicLevel:
		return "<3>"
	case zapcore.PanicLevel:
		return "<4>"
	case zapcore.FatalLevel:
		return "<5>"
	}
	return ""
}
