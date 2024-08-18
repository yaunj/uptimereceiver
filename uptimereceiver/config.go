package uptimereceiver

import (
	"fmt"
	"time"
)

// Config represents the config settings for the uptimereceiver
type Config struct {
	Interval string `mapstructure:"interval"`
}

// Validate checks if the config is valid
func (cfg *Config) validate() error {
	interval, _ := time.ParseDuration(cfg.Interval)
	if interval.Seconds() < 1 {
		return fmt.Errorf("the interval needs to be more than 1s")
	}

	return nil
}
