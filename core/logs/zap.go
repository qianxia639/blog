package logs

import (
	"Blog/core/config"
	"io"
	"os"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logs *zap.Logger

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

func initZap(conf config.Zap) *zap.Logger {
	// 新建一个配置
	encoderConfig := zapEncoderConfig()

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	topicDebugging := zapcore.AddSync(logSegmentation(conf))
	topicErrors := zapcore.AddSync(logSegmentation(conf))

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(jsonEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	return zap.New(core)
}

func logSegmentation(conf config.Zap) io.Writer {
	return &lumberjack.Logger{
		Filename:   conf.LogFile.Output,   //文件路径
		MaxSize:    conf.LogFile.MaxSize,  //分割文件的大小
		MaxAge:     conf.LogFile.MaxAge,   // 保存天数
		MaxBackups: conf.LogFile.Backups,  //最大保留数
		Compress:   conf.LogFile.Compress, // 是否压缩
		LocalTime:  true,                  //使用本地时间
	}
}

var once sync.Once

func GetInstance(conf config.Zap) *zap.Logger {
	once.Do(func() {
		Logs = initZap(conf)
	})
	return Logs
}
