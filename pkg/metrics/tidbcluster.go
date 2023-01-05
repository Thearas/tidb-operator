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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ClusterSpecReplicas = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "tidb_operator",
			Subsystem: "cluster",
			Name:      "spec_replicas",
			Help:      "Desired replicas of each component in TidbCluster",
		}, []string{LabelNamespace, LabelName, LabelComponent})

	KubeRequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "kube_request_latency_seconds",
			Buckets: prometheus.ExponentialBuckets(0.005, 2, 12),
		}, []string{"path", "method"})

	KubeRateLimiterLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "kube_request_ratelimit_seconds",
			Buckets: prometheus.ExponentialBuckets(0.005, 2, 12),
		}, []string{"path", "method"})
)
