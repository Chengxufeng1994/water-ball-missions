package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/log_framework/internal/infrastructure/logger"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/log_framework/internal/pkg"
)

func main() {
	// 初始化日誌器
	loggers := initializeLoggers()
	fmt.Println("Loggers initialized:", loggers)
	for _, logger := range loggers {
		fmt.Printf("Logger: %s, Level: %s\n", logger.GetName(), logger.GetLevel())
	}
}

func initializeLoggers() []pkg.Logger {
	data, _ := os.ReadFile("logger_config.json")

	config := Config{}
	_ = json.Unmarshal(data, &config)

	var loggers []pkg.Logger
	createLogger("root", config.Logger, nil, &loggers)

	return loggers
}

func createLogger(name string, loggerConfig Logger, parent pkg.Logger, loggers *[]pkg.Logger) pkg.Logger {
	var exporter pkg.Exporter
	switch loggerConfig.Exporter.Type {
	case "console":
		exporter = pkg.NewConsoleExporter()
	case "file":
		exporter = pkg.NewFileExporter(loggerConfig.Exporter.FileName)
	case "composite":
		var childExporters []pkg.Exporter
		for _, child := range loggerConfig.Exporter.Children {
			childExporters = append(childExporters, createExporter(child))
		}
		exporter = pkg.NewCompositeExporter(childExporters...)
	}

	var newLogger pkg.Logger
	if parent == nil {
		newLogger = logger.NewRoot(name, parseLevel(loggerConfig.LevelThreshold), nil, exporter)
	} else {
		newLogger = logger.New(parent, name, parseLevel(loggerConfig.LevelThreshold), nil, exporter)
	}

	*loggers = append(*loggers, newLogger)

	for childName, childConfig := range loggerConfig.Children {
		createLogger(childName, childConfig, newLogger, loggers)
	}

	return newLogger
}

func createExporter(exporterConfig Exporter) pkg.Exporter {
	switch exporterConfig.Type {
	case "console":
		return pkg.NewConsoleExporter()
	case "file":
		return pkg.NewFileExporter(exporterConfig.FileName)
	case "composite":
		var childExporters []pkg.Exporter
		for _, child := range exporterConfig.Children {
			childExporters = append(childExporters, createExporter(child))
		}
		return pkg.NewCompositeExporter(childExporters...)
	default:
		panic("Unknown exporter type: " + exporterConfig.Type)
	}
}

func parseLevel(level string) pkg.Level {
	switch level {
	case "TRACE":
		return pkg.LevelTrace
	case "DEBUG":
		return pkg.LevelDebug
	case "INFO":
		return pkg.LevelInfo
	case "WARN":
		return pkg.LevelWarn
	case "ERROR":
		return pkg.LevelError
	default:
		return pkg.LevelInfo
	}
}

// Config 是最外層的結構，包含 `loggers`
type Config struct {
	Logger Logger `json:"loggers"`
}

// Logger 是每個層級的結構，會包括 `levelThreshold`、`exporter` 和 `layout`，並且支持嵌套
type Logger struct {
	LevelThreshold string            `json:"levelThreshold,omitempty"`
	Exporter       Exporter          `json:"exporter,omitempty"`
	Layout         string            `json:"layout,omitempty"`   // 僅在根層出現
	Children       map[string]Logger `json:"children,omitempty"` // 嵌套的 logger 以 map 的形式存在
}

// Exporter 描述了輸出配置，可能包含嵌套的子 exporter
type Exporter struct {
	Type     string     `json:"type"`
	FileName string     `json:"fileName,omitempty"`
	Children []Exporter `json:"children,omitempty"`
}
