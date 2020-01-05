package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/xhminc/xhm-framework/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var (
	log             *zap.Logger
	timestampFormat = "2006-01-02 15:04:05.000"
)

func GetLogger() *zap.Logger {
	return InitLogger(nil)
}

func InitLogger(c *config.YAMLConfig) *zap.Logger {

	if log != nil {
		return log
	}

	encoderConfig := newEncoderConfig()
	cfg := newLoggerConfig(c, encoderConfig)
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	debugLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})

	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel
	})

	var infoHook io.Writer
	var warnHook io.Writer

	if strings.HasSuffix(c.Logging.FilePath, "/") {
		infoHook = getHook(c.Logging.FilePath + c.Logging.FileName)
		warnHook = getHook(c.Logging.FilePath + errorFilename(c.Logging.FileName))
	} else {
		infoHook = getHook(c.Logging.FilePath + "/" + c.Logging.FileName)
		warnHook = getHook(c.Logging.FilePath + "/" + errorFilename(c.Logging.FileName))
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), debugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnHook), warnLevel),
	)

	log, _ = cfg.Build(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return core
	}))

	return log
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeCaller:  zapcore.FullCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(timestampFormat))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		//EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		//	enc.AppendInt64(int64(d) / 1000000)
		//},
	}
}

func newLoggerConfig(c *config.YAMLConfig, encoderConfig zapcore.EncoderConfig) zap.Config {

	var level zap.AtomicLevel
	var development bool
	var disableCaller bool

	if c.Application.Profile == "dev" || c.Application.Profile == "test" {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
		development = true
		disableCaller = false
	} else {
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
		development = false
		disableCaller = true
	}

	cfg := zap.Config{
		Level:             level,
		Development:       development,
		DisableCaller:     disableCaller,
		DisableStacktrace: false,
		Encoding:          "console",
		EncoderConfig:     encoderConfig,
	}

	return cfg
}

func getHook(filename string) io.Writer {

	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}

	return hook
}

func errorFilename(filename string) string {
	pos := strings.LastIndex(filename, ".")
	if pos != -1 {
		return filename[0:pos] + "_error" + filename[pos:]
	} else {
		return filename + "_error"
	}
}
