package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"runtime"
	"time"
)

type ErrorInfo struct {
	//EnvID         string `json:"env_id"`
	TracebackInfo string `json:"traceback_info"`
	AppName       string `json:"app_name"`
	Filename      string `json:"filename"`
	LineNo        int    `json:"line_no"`
	FuncName      string `json:"func_name"`
	//TraceID       string `json:"trace_id,omitempty"`
	Error string `json:"error"`
}

// NewErrorInfo initializes an ErrorInfo with default environment values.
func NewErrorInfo(tracebackInfo, appName, filename, funcName, err string, lineNo int) *ErrorInfo {
	//envID := os.Getenv("ENV_ID")
	return &ErrorInfo{
		//EnvID:         envID,
		TracebackInfo: tracebackInfo,
		AppName:       appName,
		Filename:      filename,
		LineNo:        lineNo,
		FuncName:      funcName,
		//TraceID:       traceId,
		Error: err,
	}
}

// Logger 日志插件
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	ErrorfFeishuAlert(format string, v ...interface{}) // 带飞书告警的错误日志记录
}

var (
	loggerIns Logger = &defaultLog{logger: initZapLogger()}

	Debug  = loggerIns.Debug
	Debugf = loggerIns.Debugf

	Info  = loggerIns.Info
	Infof = loggerIns.Infof

	Error  = loggerIns.Error
	Errorf = loggerIns.Errorf

	ErrorfAlert = loggerIns.ErrorfFeishuAlert // Expose the new method
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

func (d *defaultLog) ErrorfFeishuAlert(format string, v ...interface{}) {
	msg := getLogMsgf(format, v...)
	d.logger.Error(msg)
	// 飞书告警
	feishuAlert(msg)
}

func feishuAlert(msg string) bool {
	// 获取堆栈信息
	pc, file, line, ok := runtime.Caller(1)
	funcName := ""
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	stackBuf := make([]byte, 1024)
	stackLen := runtime.Stack(stackBuf, false)
	stackTrace := string(stackBuf[:stackLen])
	// 构建错误信息并发送到飞书
	errorInfo := NewErrorInfo(stackTrace, "ai-paper", file, funcName, msg, line)
	feishuAlertURL := os.Getenv("FEISHU_ALERT_URL")
	if feishuAlertURL == "" {
		feishuAlertURL = "https://www.feishu.cn/flow/api/trigger-webhook/7b6bc4fae0628c365a436cadd8ad9799"
	}
	data, _ := json.Marshal(errorInfo)
	client := &http.Client{Timeout: 300 * time.Second}
	resp, _ := client.Post(feishuAlertURL, "application/json", bytes.NewBuffer(data))
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
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
