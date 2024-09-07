package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestDebugf(t *testing.T) {
	for i := 0; i < 10; i++ {
		Debug("log i")
		Debugf("log i:%v", i)
	}
}

func TestZapLog(t *testing.T) {
	encoderConf := zapcore.EncoderConfig{
		CallerKey:        "caller_line", // 打印文件名和行数
		LevelKey:         "level_name",
		MessageKey:       "msg",
		TimeKey:          "ts",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeTime:       zapcore.ISO8601TimeEncoder,    // 自定义时间格式
		EncodeLevel:      zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeCaller:     zapcore.FullCallerEncoder,     // 全路径编码器
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: " ",
	}

	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConf
	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()
	if err != nil {
		t.Logf("config build failed, err:%v", err)
		return
	}
	logger.Info("Hello, Zap!")
}

func Test_msgEncoder(t *testing.T) {
	msg := msgEncoder("%v", 1)
	t.Logf(msg)
}
