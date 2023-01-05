// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"net/url"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/tools/metrics"
)

// RegisterMetrics registers all metrics of tidb-operator.
func RegisterMetrics() {
	prometheus.MustRegister(ClusterSpecReplicas)
	metrics.Register(metrics.RegisterOpts{
		RequestLatency:     &KubeLatencyMetric{latencySeconds: KubeRequestLatency},
		RateLimiterLatency: &KubeLatencyMetric{latencySeconds: KubeRateLimiterLatency},
	})
}

type KubeLatencyMetric struct {
	latencySeconds *prometheus.HistogramVec
}

func (k *KubeLatencyMetric) Observe(verb string, u url.URL, latency time.Duration) {
	k.latencySeconds.WithLabelValues(u.Path, verb).Observe(latency.Seconds())
}

// Label constants.
const (
	LabelNamespace = "namespace"
	LabelName      = "name"
	LabelComponent = "component"
)
