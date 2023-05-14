package logs

import (
	"Blog/core/config"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logs *zap.Logger

func zapEncoder(config *config.Config) zapcore.Encoder {
	// 新建一个配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "Time",
		LevelKey:      "Level",
		NameKey:       "Logger",
		CallerKey:     "Caller",
		MessageKey:    "Message",
		StacktraceKey: "StackTrace",
		EncodeCaller:  zapcore.ShortCallerEncoder,
		LineEnding:    zapcore.DefaultLineEnding,
		FunctionKey:   zapcore.OmitKey,
	}
	// 自定义时间格式
	encoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	// 日志级别大写
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 秒级时间间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 简短的调用者输出
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 完整的序列化logger名称
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	// 最终的日志编码 json或者console
	switch config.Zap.Encode {
	case "json":
		return zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	// 默认console
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriteSyncer(cfg *config.Config) zapcore.WriteSyncer {
	syncers := make([]zapcore.WriteSyncer, 0, 2)
	// 如果开启了日志控制台输出，就加入控制台书写器
	if cfg.Zap.Writer == "both" || cfg.Zap.Writer == "console" {
		syncers = append(syncers, zapcore.AddSync(os.Stdout))
	}

	// 如果开启了日志文件存储，就根据文件路径切片加入书写器
	if cfg.Zap.Writer == "both" || cfg.Zap.Writer == "file" {
		// 添加日志输出器
		logger := &lumberjack.Logger{
			Filename:   cfg.Zap.LogFile.Output,   //文件路径
			MaxSize:    cfg.Zap.LogFile.MaxSize,  //分割文件的大小
			MaxBackups: cfg.Zap.LogFile.Backups,  //备份次数
			Compress:   cfg.Zap.LogFile.Compress, // 是否压缩
			LocalTime:  true,                     //使用本地时间
		}
		syncers = append(syncers, zapcore.Lock(zapcore.AddSync(logger)))
	}
	return zap.CombineWriteSyncers(syncers...)
}

// 日志级别
func zapLevelEnabler(cfg *config.Config) zapcore.LevelEnabler {
	switch cfg.Zap.Level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	}
	// 默认Debug级别
	return zap.DebugLevel
}

func InitZap(config *config.Config) *zap.Logger {
	// 构建编码器
	encoder := zapEncoder(config)
	// 构建日志级别
	levelEnabler := zapLevelEnabler(config)
	// 最后获得Core和Options
	subCore, options := tee(config, encoder, levelEnabler)
	// 创建Logger
	return zap.New(subCore, options...)
}

// 将所有合并
func tee(cfg *config.Config, encoder zapcore.Encoder, levelEnabler zapcore.LevelEnabler) (core zapcore.Core, options []zap.Option) {
	sink := zapWriteSyncer(cfg)
	return zapcore.NewCore(encoder, sink, levelEnabler), nil
}

// 构建Option
func buildOptions(cfg *config.Config, levelEnabler zapcore.LevelEnabler) (options []zap.Option) {
	if cfg.Zap.Caller {
		options = append(options, zap.AddCaller())
	}

	if cfg.Zap.StackTrace {
		options = append(options, zap.AddStacktrace(levelEnabler))
	}
	return
}
