package main

import (
	"os"

	"github.com/Rindrics/expect-label-prefix/application"
	"github.com/Rindrics/expect-label-prefix/domain"
	"github.com/Rindrics/expect-label-prefix/infra"
)

func main() {
	logger := infra.ParseLogLevel()

	logger.Info("Loading webhook event")
	event, err := infra.LoadEventFromEnv()
	if err != nil {
		logger.Error("Failed to load event from environment", "error", err)
		return
	}

	logger.Info("Parsing webhook event")
	eventInfo := infra.ParseEvent(event, logger)

	config := application.NewConfig()
	logger.Debug("config", "loaded", config)

	logger.Debug("event info", "number", eventInfo.Number, "labels", eventInfo.Labels)

	rl := domain.RequiredLabel{
		Prefix:    config.Prefix,
		Separator: config.Separator,
	}

	found := rl.DoExist(eventInfo.Labels)
	if found {
		logger.Info("found label with required prefix")
		os.Exit(0)
	}

	logger.Info("label with required prefix not found")
	client := infra.NewGitHubClient(config.Token, logger)

	app := application.New(eventInfo, client, *config, logger)
	err = app.Run()
	if err != nil {
		logger.Error("Error running application", "error", err)
		os.Exit(1)
	}
}
