package uptimereceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

type uptimeReceiver struct {
	host         component.Host
	cancel       context.CancelFunc
	logger       *zap.Logger
	nextConsumer consumer.Metrics
	config       *Config
}

func (u uptimeReceiver) Start(ctx context.Context, host component.Host) error {
	u.host = host
	ctx = context.Background()
	ctx, u.cancel = context.WithCancel(ctx)

	u.logger.Info("Starting uptimeReceiver")

	interval, _ := time.ParseDuration(u.config.Interval)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				u.logger.Info("Getting uptime")
				uptime, err := getUptime()
				if err != nil {
					u.logger.Error("Failed to get uptime", zap.Error(err))
				} else {
					u.logger.Info("Got uptime", zap.Duration("uptime", uptime))
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	u.logger.Info("Started uptimeReceiver")
	return nil
}

func (u uptimeReceiver) Shutdown(ctx context.Context) error {
	u.logger.Info("Stopping uptimereceiver")
	if u.cancel != nil {
		u.cancel()
	}
	return nil
}
