package uptimereceiver

import (
	"context"
	"fmt"
	"time"

	"github.com/yaunj/uptimereceiver/uptimereceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"
)

const (
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	return &Config{
		Interval: string(defaultInterval),
		ControllerConfig: scraperhelper.NewDefaultControllerConfig(),
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
	}
}

func createMetricsReceiver(_ context.Context, settings receiver.Settings, baseCfg component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	logger := settings.Logger
	uptimereceiverConfig, ok := baseCfg.(*Config)
	if !ok {
		logger.Error("failed to cast receiver config")
		return nil, fmt.Errorf("failed to cast receiver config")
	}

	ns := newScraper(uptimereceiverConfig, settings)
	scraper, err := scraperhelper.NewScraper(metadata.Type.String(), ns.scrape)
	if err != nil {
		logger.Error("failed to create scraper", zap.Error(err))
		return nil, err
	}

	return scraperhelper.NewScraperControllerReceiver(
		&uptimereceiverConfig.ControllerConfig,
		settings,
		consumer,
		scraperhelper.AddScraper(scraper),
	)
}

// NewFactory() creates a factory for the uptimereceiver
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelAlpha),
	)
}
