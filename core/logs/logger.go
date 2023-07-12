package logs

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	Logger = initLogger()
}

func zapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "Time",                           // 日志时间对应的key
		LevelKey:      "Level",                          // 日志级别对应的key
		NameKey:       "Logger",                         //
		CallerKey:     "Caller",                         // 日志所在文件对应的key
		MessageKey:    "Message",                        // 结构化(json)输出：msg的key
		StacktraceKey: "StackTrace",                     //
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, // 日志级别大写
		EncodeCaller:  zapcore.ShortCallerEncoder,       // 相对路径编码输出
		LineEnding:    zapcore.DefaultLineEnding,
		FunctionKey:   zapcore.OmitKey,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) { // 自定义时间格式
			pae.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, // 秒级时间间隔
	}
}

func initLogger() *zap.Logger {
	// 新建一个配置
	encoderConfig := zapEncoderConfig()

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	topicDebugging := zapcore.AddSync(logSegmentation())
	topicErrors := zapcore.AddSync(logSegmentation())

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(fileEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	return zap.New(core, zap.AddCaller())
}

func logSegmentation() io.Writer {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("./log/%s.log", time.Now().Format("2006-01-02")), //文件路径
		MaxSize:    20,                                                           //分割文件的大小
		MaxAge:     7,                                                            // 保存天数
		MaxBackups: 5,                                                            //最大保留数
		Compress:   true,                                                         // 是否压缩
		LocalTime:  true,                                                         //使用本地时间
	}
}
