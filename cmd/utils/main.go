package utils

import (
	"runtime"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func HandleError(e error) {
	if e != nil {
		Logger.Error("Error", zap.Error(e))
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	Logger.Info("Alloc", zap.Uint64("Mem", bToMb(m.Alloc)))
	Logger.Info("TotalAlloc", zap.Uint64("Mem", bToMb(m.TotalAlloc)))
	Logger.Info("Sys", zap.Uint64("Mem", bToMb(m.Sys)))
	Logger.Info("NumGC", zap.Uint32("Mem", m.NumGC))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func Recovery() {
	if r := recover(); r != nil {
		Logger.Info("Recovered in f", zap.Any("r", r))
	}
}

func InitLogger() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
}
