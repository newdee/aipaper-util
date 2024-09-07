package log

import (
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logTmFmtWithMS = "2006-01-02 15:04:05.000"

// Level 日志等级，可以动态修改
var Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

func initZapLogger() *zap.Logger {

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format(logTmFmtWithMS) + "]")
	}
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + os.Getenv("SERVER_NAME") + "]") // server name
		enc.AppendString("[" + caller.TrimmedPath() + "]")
		shortFuncName := caller.Function
		fullName := strings.Split(caller.Function, "/")
		if len(fullName) > 0 {
			shortFuncName = fullName[len(fullName)-1]
		}
		enc.AppendString("[" + shortFuncName + "]")
		enc.AppendString("[" + os.Getenv("ENV") + "]")
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:        "caller_line", // 打印文件名和行数
		LevelKey:         "level_name",
		MessageKey:       "msg",
		TimeKey:          "ts",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeTime:       customTimeEncoder,   // 自定义时间格式
		EncodeLevel:      customLevelEncoder,  // 小写编码器
		EncodeCaller:     customCallerEncoder, // 全路径编码器
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: " ",
	}

	encoder := zapcore.NewConsoleEncoder(encoderConf)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/mix-paper.log",
		MaxSize:    10, // 单位为M
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   false,
	}
	writerSyncer := zapcore.AddSync(lumberJackLogger)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writerSyncer, Level),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), Level),
	)
	logger := zap.New(core, zap.AddCallerSkip(1), zap.AddCaller())
	return logger
}
