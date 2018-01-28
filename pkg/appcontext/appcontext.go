package appcontext

import (
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

type AppContext struct {
	config *config.Config
	logger *logger.Logger
}

func NewAppContext(config *config.Config, logger *logger.Logger) *AppContext {
	return &AppContext{
		config: config,
		logger: logger,
	}
}

func (a *AppContext) GetConfig() *config.Config {
	return a.config
}

func (a *AppContext) GetLogger() *logger.Logger {
	return a.logger
}