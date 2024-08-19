package uptimereceiver

import (
	"fmt"
	"time"

	"github.com/yaunj/uptimereceiver/uptimereceiver/internal/metadata"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

// Config represents the config settings for the uptimereceiver
type Config struct {
	// Interval is not used anymore, but temporarily serves as an example of how to handle extra config values
	Interval string `mapstructure:"interval"`

	// MetricsBuilderConfig to enable/disable specific metrics (default: all enabled)
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
	// ControllerConfig to configure scraping interval (default: every second?)
	scraperhelper.ControllerConfig `mapstructure:",squash"`
}

// Validate checks if the config is valid
func (cfg *Config) validate() error {
	interval, _ := time.ParseDuration(cfg.Interval)
	if interval.Seconds() < 1 {
		return fmt.Errorf("the interval needs to be more than 1s")
	}

	return nil
}
