package main

import (
	"fmt"
	"time"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/log_framework/internal/infrastructure/logger"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/log_framework/internal/pkg"
)

func main() {
	root := logger.NewRoot("root", pkg.LevelDebug, pkg.NewStandardLayout(), pkg.NewConsoleExporter())

	gameLogger := logger.New(root, "app.game", pkg.LevelInfo, pkg.NewStandardLayout(), pkg.NewCompositeExporter(
		pkg.NewConsoleExporter(),
		pkg.NewCompositeExporter(
			pkg.NewFileExporter("game.log"),
			pkg.NewFileExporter("game.backup.log"),
		),
	))

	aiLogger := logger.New(gameLogger, "app.game.ai", pkg.LevelTrace, pkg.NewStandardLayout(), nil)

	gameLogger.Info("The game begins.")
	for i := range 4 {
		gameLogger.Trace(fmt.Sprintf("The player %d begins his turn.", i+1))
		aiLogger.Trace(fmt.Sprintf("The player %d starts making decisions...", i+1))
		aiLogger.Warn(fmt.Sprintf("The player %d  decides to give up.", i+1))
		aiLogger.Error("Something goes wrong when AI gives up.")
		aiLogger.Trace(fmt.Sprintf("The player %d  completes its decision.", i+1))
		gameLogger.Trace(fmt.Sprintf("The player %d finishes his turn.", i+1))
	}

	gameLogger.Debug("Game ends")

	time.Sleep(2 * time.Second)
}
