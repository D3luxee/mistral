/*-
 * Copyright © 2017, Jörg Pernfuß <code.jpe@gmail.com>
 * All rights reserved.
 *
 * Use of this source code is governed by a 2-clause BSD license
 * that can be found in the LICENSE file.
 */

package mistral // import "github.com/solnx/mistral/internal/mistral"

import (
	"fmt"
	"os"

	"github.com/solnx/legacy"
	metrics "github.com/rcrowley/go-metrics"
)

// FormatMetrics is the formatting function to export Mistral metrics
// via legacy.MetricSocket, implementing legacy.Formatter
func FormatMetrics(batch *legacy.PluginMetricBatch) func(string, interface{}) {
	return func(metric string, v interface{}) {
		switch v.(type) {
		case *metrics.StandardMeter:
			value := v.(*metrics.StandardMeter)
			batch.Metrics = append(batch.Metrics, legacy.PluginMetric{
				Type:   `float`,
				Metric: fmt.Sprintf("%s/avg/rate/1min", metric),
				Value: legacy.MetricValue{
					FlpVal: value.Rate1(),
				},
			})
		}
	}
}

// DebugFormatMetrics is the formatting function to print Mistral metrics
// on STDERR
func DebugFormatMetrics(_ *legacy.PluginMetricBatch) func(string, interface{}) {
	return func(metric string, v interface{}) {
		switch v.(type) {
		case *metrics.StandardMeter:
			value := v.(*metrics.StandardMeter)
			fmt.Fprintf(os.Stderr, "%s/avg/rate/1min: %f\n",
				metric, value.Rate1())
		}
	}
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
