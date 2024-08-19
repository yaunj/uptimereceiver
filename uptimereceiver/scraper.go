package uptimereceiver

import (
	"context"
	"time"

	"github.com/yaunj/uptimereceiver/uptimereceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
)

type scraper struct {
	logger          *zap.Logger
	metricsBuilder  *metadata.MetricsBuilder // MetricsBuilder to build metrics
	attributeReader *attributeReader         // gets system attributes
}

func newScraper(cfg *Config, settings receiver.Settings) *scraper {
	return &scraper{
		logger:          settings.TelemetrySettings.Logger,
		metricsBuilder:  metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
		attributeReader: newAttributeReader(settings.TelemetrySettings.Logger),
	}
}

func (s *scraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	s.logger.Debug("Getting uptime")

	uptime, err := getUptime()
	if err != nil {
		s.logger.Error("Failed to get uptime", zap.Error(err))
		return pmetric.Metrics{}, err
	}
	s.logger.Info("Got uptime", zap.Duration("uptime", uptime))

	attrs := s.attributeReader.getAttributes()
	s.recordUptime(uptime, attrs)
	return s.metricsBuilder.Emit(), nil
}

func (s *scraper) recordUptime(uptime time.Duration, attrs *attributes) {
	now := pcommon.NewTimestampFromTime(time.Now())
	s.metricsBuilder.RecordUptimeDataPoint(now, int64(uptime.Seconds()), attrs.hostname, attrs.os, attrs.arch, "uptime")
	s.logger.Debug(
		"Added stuff to metricsBuilder",
		zap.String("hostname", attrs.hostname),
		zap.String("os", attrs.os),
		zap.String("arch", attrs.arch),
		zap.Duration("uptime", uptime),
	)
}
