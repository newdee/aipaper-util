package log

import (
	"fmt"

	"go.uber.org/zap"
)

// Logger 日志插件
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

var (
	loggerIns Logger = &defaultLog{logger: initZapLogger()}

	Debug  = loggerIns.Debug
	Debugf = loggerIns.Debugf

	Info  = loggerIns.Info
	Infof = loggerIns.Infof

	Error  = loggerIns.Error
	Errorf = loggerIns.Errorf
)

type defaultLog struct {
	logger *zap.Logger
}

func (d *defaultLog) Debug(v ...interface{}) {
	d.logger.Debug(getLogMsg(v...))
}

func (d *defaultLog) Debugf(format string, v ...interface{}) {
	d.logger.Debug(getLogMsgf(format, v...))
}

func (d *defaultLog) Info(v ...interface{}) {
	d.logger.Info(getLogMsg(v...))
}

func (d *defaultLog) Infof(format string, v ...interface{}) {
	d.logger.Info(getLogMsgf(format, v...))
}

func (d *defaultLog) Error(v ...interface{}) {
	d.logger.Error(getLogMsg(v...))
}

func (d *defaultLog) Errorf(format string, v ...interface{}) {
	d.logger.Error(getLogMsgf(format, v...))
}

func msgEncoder(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	msg = "[" + msg + "]"
	return msg
}

func getLogMsg(args ...interface{}) string {
	msg := fmt.Sprint(args...)
	msg = "[" + msg + "]"
	return msg
}

func getLogMsgf(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	msg = "[" + msg + "]"
	return msg
}
