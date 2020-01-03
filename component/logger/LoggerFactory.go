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

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		TimeKey:        "ts",
		EncodeTime:     func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { enc.AppendString(t.Format(timestampFormat)) },
		CallerKey:      "file",
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) { enc.AppendInt64(int64(d) / 1000000) },
	})

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

	log = zap.New(core, zap.AddCaller())
	return log
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
