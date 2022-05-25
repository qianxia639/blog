package utils

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	infoPath  = "./log/info." + time.Now().Format("2006-01-02") + ".log"
	errorPath = "./log/error." + time.Now().Format("2006-01-02") + ".log"
)

// 初始化日志
func Zap() *zap.SugaredLogger {
	config := zapcore.EncoderConfig{
		MessageKey:   "message",                   // 结构化(json)输出: msg的key
		LevelKey:     "level",                     // 日志级别的key
		TimeKey:      "time",                      // 时间的key
		CallerKey:    "file",                      // 打印日志的文件对应的key
		EncodeLevel:  zapcore.CapitalLevelEncoder, // 将日志级别转换成大写(INFO,WARN,ERROR等)
		EncodeCaller: zapcore.FullCallerEncoder,   // 采用完整文件路径编码输出 (d://.../test/main.go:14)
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: func(d time.Duration, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendInt64(int64(d) / 1000000)
		},
	}

	// 自定义日志级别 Info
	infoLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zapcore.WarnLevel && l >= zap.InfoLevel
	})

	// 自定义日志级别 warn
	warnLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.WarnLevel && l >= zap.InfoLevel
	})

	// 获取io.writer的实现
	infoWriter := getWriter(infoPath)
	warnWriter := getWriter(errorPath)

	core := zapcore.NewTee(
		// 将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(infoWriter), infoLevel),
		// 将warn及以上写入errPath
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(warnWriter), warnLevel),
		// 同时将日志输入到控制台，NewJSONEncoder是结构化输出
		// zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), infoLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	return logger.Sugar()
}

// 日志切割配置
func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename, // 日志存放文件名
		MaxSize:    10,       // 最大内存占用(单位：M),超过则切割
		MaxBackups: 10,       // 最大文件保留数，超过会删除最老的日志文件
		MaxAge:     30,       // 保存天数
		Compress:   true,     // 是否压缩
	}
}
