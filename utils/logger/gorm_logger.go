package logger

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const (
	// Silent SilentMode
	Silent gormlogger.LogLevel = iota + 1
	// Error ErrorLevel
	Error
	// Warn WarnLevel
	Warn
	// Info InfoLevel
	Info
)

// GormZapLogger Gorm対応のZapLogger
type GormZapLogger struct {
	LogLevel      gormlogger.LogLevel
	SlowThreshold time.Duration
}

// Default デフォルト設定 (スロークエリは1秒以上)
var Default = NewGormZapLogger(Info, 1000*time.Millisecond)

// NewGormZapLogger コンストラクタ
func NewGormZapLogger(logLevel gormlogger.LogLevel, slowThreshold time.Duration) *GormZapLogger {
	return &GormZapLogger{
		LogLevel:      logLevel,
		SlowThreshold: slowThreshold,
	}
}

// LogMode ログモード
func (l *GormZapLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info Infoログ出力
func (l *GormZapLogger) Info(ctx context.Context, s string, args ...interface{}) {
	if l.LogLevel < Info {
		Logger.Info(s, zap.Any("data", args))
	}
}

// Warn Warnログ出力
func (l *GormZapLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	if l.LogLevel >= Warn {
		Logger.Warn(s, zap.Any("data", args))
	}
}

// Error Errorログ出力
func (l *GormZapLogger) Error(ctx context.Context, s string, args ...interface{}) {
	if l.LogLevel >= Error {
		Logger.Error(s, zap.Any("data", args))
	}
}

// Trace Traceログ出力
func (l *GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > Silent {
		elapsed := time.Since(begin)

		sql, rows := fc()

		if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
			Logger.Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
			return
		}

		if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
			Logger.Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
			return
		}

		Logger.Info("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		return
	}
}
