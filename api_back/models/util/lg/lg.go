package lg

import (
	"dataxWeb/models/util/tool"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var lg *zap.Logger

func init() {
	writeSyncer := getLogWriter("log/dataxWeb.log", 10, 5)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte("debug"))
	if err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	// zap.AddCallerSkip(1) 向前跳几层（由于把 logger封装成函数时调用中使用的）
	lg = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	//return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{ // 日志切割
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		Compress:   true, // 压缩备份的日志
	}
	ws := io.MultiWriter(lumberJackLogger, os.Stderr) // 打印到终端，并写入日志文件
	return zapcore.AddSync(ws)
}

// SqlLogDebug sql语句输出
func SqlLogDebug(sql string, parameter ...interface{}) {
	lg.Debug(fmt.Sprintf("[SQL] \t %v \nsql参数===>: %v", sql, parameter))
}

// Error error 日志输出
func Error(err error, s ...string) {
	if s != nil {
		lg.Error(fmt.Sprintf("[ERROR] \t %v <=== \t %v ", s, err.Error()))
	} else {
		lg.Error(err.Error())
	}
}

// ErrorStr
//
//	@Description: 字符串error日志输出
//	@param err
//	@param s
func ErrorStr(err string, s ...string) {
	lg.Error(fmt.Sprintf("[ERROR] \t %v <=== \t %v ", s, err))
}

// Info Info 日志输出
func Info(s interface{}) {
	lg.Info("\t" + tool.ToString(s))
}

// Debug Debug 日志输出
func Debug(s string) {
	lg.Debug(s)
}

// Warn Warn 日志
func Warn(s string) {
	lg.Warn(s)
}
