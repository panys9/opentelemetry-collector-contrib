// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for filesystem metrics.
type MetricsConfig struct {
	SystemFilesystemInodesUsage MetricConfig `mapstructure:"system.filesystem.inodes.usage"`
	SystemFilesystemUsage       MetricConfig `mapstructure:"system.filesystem.usage"`
	SystemFilesystemUtilization MetricConfig `mapstructure:"system.filesystem.utilization"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SystemFilesystemInodesUsage: MetricConfig{
			Enabled: true,
		},
		SystemFilesystemUsage: MetricConfig{
			Enabled: true,
		},
		SystemFilesystemUtilization: MetricConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for filesystem metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
