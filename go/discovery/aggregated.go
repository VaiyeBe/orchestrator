/*
<<<<<<< HEAD
   Copyright 2016 Simon J Mudd
=======
   Copyright 2017 Simon J Mudd
>>>>>>> refs/remotes/github/master

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package discovery

import (
	"time"

	"github.com/montanaflynn/stats"
<<<<<<< HEAD
)

// this contains routines for cooked DiscoveryMetric values
// which can be called from api/discovery-metrics-aggregated/:seconds

type AggregatedDiscoveryMetrics struct {
	FirstSeen                       time.Time // timestamp of the first data seen
	LastSeen                        time.Time // timestamp of the last data seen
	CountDistinctInstanceKeys       int       // number of distinct Instances seen
=======

	"github.com/github/orchestrator/go/collection"
)

// AggregatedDiscoveryMetrics contains aggregated metrics for instance discovery.
// Called from api/discovery-metrics-aggregated/:seconds
type AggregatedDiscoveryMetrics struct {
	FirstSeen                       time.Time // timestamp of the first data seen
	LastSeen                        time.Time // timestamp of the last data seen
	CountDistinctInstanceKeys       int       // number of distinct Instances seen (note: this may not be true: distinct = succeeded + failed)
	CountDistinctOkInstanceKeys     int       // number of distinct Instances which succeeded
>>>>>>> refs/remotes/github/master
	CountDistinctFailedInstanceKeys int       // number of distinct Instances which failed
	FailedDiscoveries               uint64    // number of failed discoveries
	SuccessfulDiscoveries           uint64    // number of successful discoveries
	MeanTotalSeconds                float64
	MeanBackendSeconds              float64
	MeanInstanceSeconds             float64
	FailedMeanTotalSeconds          float64
	FailedMeanBackendSeconds        float64
	FailedMeanInstanceSeconds       float64
	MaxTotalSeconds                 float64
	MaxBackendSeconds               float64
	MaxInstanceSeconds              float64
	FailedMaxTotalSeconds           float64
	FailedMaxBackendSeconds         float64
	FailedMaxInstanceSeconds        float64
	MedianTotalSeconds              float64
	MedianBackendSeconds            float64
	MedianInstanceSeconds           float64
	FailedMedianTotalSeconds        float64
	FailedMedianBackendSeconds      float64
	FailedMedianInstanceSeconds     float64
	P95TotalSeconds                 float64
	P95BackendSeconds               float64
	P95InstanceSeconds              float64
	FailedP95TotalSeconds           float64
	FailedP95BackendSeconds         float64
	FailedP95InstanceSeconds        float64
}

<<<<<<< HEAD
// internal routine to return the average value or 0
func mean(values stats.Float64Data) float64 {
	s, err := stats.Mean(values)
	if err != nil {
		return 0
	}
	return s
}

// internal routine to return the requested percentile valeu or 0
func percentile(values stats.Float64Data, percent float64) float64 {
	s, err := stats.Percentile(values, percent)
	if err != nil {
		return 0
	}
	return s
}

// internal routine to return the maximum value or 0
func max(values stats.Float64Data) float64 {
	s, err := stats.Max(values)
	if err != nil {
		return 0
	}
	return s
}

// internal routine to return the median or 0
func median(values stats.Float64Data) float64 {
	s, err := stats.Median(values)
	if err != nil {
		return 0
	}
	return s
}

// AggregatedSince returns a large number of aggregated metrics
// based on the raw metrics collected since the given time.
func (mc *MetricCollection) AggregatedSince(t time.Time) (AggregatedDiscoveryMetrics, error) {
	results, err := mc.Since(t)
	if err != nil {
		return AggregatedDiscoveryMetrics{}, err
	}

	return aggregate(results), nil
}

func aggregate(results [](*Metric)) AggregatedDiscoveryMetrics {
=======
// aggregate returns the aggregate values of the given metrics (assumed to be Metric)
func aggregate(results []collection.Metric) AggregatedDiscoveryMetrics {
>>>>>>> refs/remotes/github/master
	if len(results) == 0 {
		return AggregatedDiscoveryMetrics{}
	}

	var (
		first time.Time
		last  time.Time
	)

<<<<<<< HEAD
	counters := make(map[string]uint64)             // map of string based counters
	names := make(map[string](map[string]int))      // map of string based names (using a map)
	timings := make(map[string](stats.Float64Data)) // map of string based float64 values

	// initialise counters
	for _, v := range []string{"FailedDiscoveries", "Discoveries"} {
		counters[v] = 0
	}
	// initialise names
	for _, v := range []string{"InstanceKeys", "FailedInstanceKeys"} {
		names[v] = make(map[string]int)
	}
	// initialise timers
	for _, v := range []string{"TotalSeconds", "BackendSeconds", "InstanceSeconds", "FailedTotalSeconds", "FailedBackendSeconds", "FailedInstanceSeconds"} {
=======
	type counterKey string
	type hostKey string
	type timerKey string
	const (
		FailedDiscoveries     counterKey = "FailedDiscoveries"
		Discoveries                      = "Discoveries"
		InstanceKeys          hostKey    = "InstanceKeys"
		OkInstanceKeys                   = "OkInstanceKeys"
		FailedInstanceKeys               = "FailedInstanceKeys"
		TotalSeconds          timerKey   = "TotalSeconds"
		BackendSeconds                   = "BackendSeconds"
		InstanceSeconds                  = "InstanceSeconds"
		FailedTotalSeconds               = "FailedTotalSeconds"
		FailedBackendSeconds             = "FailedBackendSeconds"
		FailedInstanceSeconds            = "FailedInstanceSeconds"
	)

	counters := make(map[counterKey]uint64)           // map of string based counters
	names := make(map[hostKey](map[string]int))       // map of string based names (using a map)
	timings := make(map[timerKey](stats.Float64Data)) // map of string based float64 values

	// initialise counters
	for _, v := range []counterKey{FailedDiscoveries, Discoveries} {
		counters[v] = 0
	}
	// initialise names
	for _, v := range []hostKey{InstanceKeys, FailedInstanceKeys, OkInstanceKeys} {
		names[v] = make(map[string]int)
	}
	// initialise timers
	for _, v := range []timerKey{TotalSeconds, BackendSeconds, InstanceSeconds, FailedTotalSeconds, FailedBackendSeconds, FailedInstanceSeconds} {
>>>>>>> refs/remotes/github/master
		timings[v] = nil
	}

	// iterate over results storing required values
<<<<<<< HEAD
	for _, v := range results {
=======
	for _, v2 := range results {
		v := v2.(*Metric) // convert to the right type

>>>>>>> refs/remotes/github/master
		// first and last
		if first.IsZero() || first.After(v.Timestamp) {
			first = v.Timestamp
		}
		if last.Before(v.Timestamp) {
			last = v.Timestamp
		}

<<<<<<< HEAD
		// successful names
		x := names["InstanceKeys"]
		x[v.InstanceKey.String()] = 1 // Value doesn't matter
		names["InstanceKeys"] = x
		// failed names
		if v.Err != nil {
			x := names["FailedInstanceKeys"]
			x[v.InstanceKey.String()] = 1 // Value doesn't matter
			names["FailedInstanceKeys"] = x
		}

		// discoveries
		counters["Discoveries"]++
		if v.Err != nil {
			counters["FailedDiscoveries"]++
		}

		// All timings
		timings["TotalSeconds"] = append(timings["TotalSeconds"], v.TotalLatency.Seconds())
		timings["BackendSeconds"] = append(timings["BackendSeconds"], v.BackendLatency.Seconds())
		timings["InstanceSeconds"] = append(timings["InstanceSeconds"], v.InstanceLatency.Seconds())

		// Failed timings
		if v.Err != nil {
			timings["FailedTotalSeconds"] = append(timings["FailedTotalSeconds"], v.TotalLatency.Seconds())
			timings["FailedBackendSeconds"] = append(timings["FailedBackendSeconds"], v.BackendLatency.Seconds())
			timings["FailedInstanceSeconds"] = append(timings["FailedInstanceSeconds"], v.InstanceLatency.Seconds())
=======
		// different names
		x := names[InstanceKeys]
		x[v.InstanceKey.String()] = 1 // Value doesn't matter
		names[InstanceKeys] = x

		if v.Err == nil {
			// ok names
			x := names[OkInstanceKeys]
			x[v.InstanceKey.String()] = 1 // Value doesn't matter
			names[OkInstanceKeys] = x
		} else {
			// failed names
			x := names[FailedInstanceKeys]
			x[v.InstanceKey.String()] = 1 // Value doesn't matter
			names[FailedInstanceKeys] = x
		}

		// discoveries
		counters[Discoveries]++
		if v.Err != nil {
			counters[FailedDiscoveries]++
		}

		// All timings
		timings[TotalSeconds] = append(timings[TotalSeconds], v.TotalLatency.Seconds())
		timings[BackendSeconds] = append(timings[BackendSeconds], v.BackendLatency.Seconds())
		timings[InstanceSeconds] = append(timings[InstanceSeconds], v.InstanceLatency.Seconds())

		// Failed timings
		if v.Err != nil {
			timings[FailedTotalSeconds] = append(timings[FailedTotalSeconds], v.TotalLatency.Seconds())
			timings[FailedBackendSeconds] = append(timings[FailedBackendSeconds], v.BackendLatency.Seconds())
			timings[FailedInstanceSeconds] = append(timings[FailedInstanceSeconds], v.InstanceLatency.Seconds())
>>>>>>> refs/remotes/github/master
		}
	}

	return AggregatedDiscoveryMetrics{
		FirstSeen:                       first,
		LastSeen:                        last,
<<<<<<< HEAD
		CountDistinctInstanceKeys:       len(names["InstanceKeys"]),
		CountDistinctFailedInstanceKeys: len(names["FailedInstanceKeys"]),
		FailedDiscoveries:               counters["FailedDiscoveries"],
		SuccessfulDiscoveries:           counters["Discoveries"],
		MeanTotalSeconds:                mean(timings["TotalSeconds"]),
		MeanBackendSeconds:              mean(timings["BackendSeconds"]),
		MeanInstanceSeconds:             mean(timings["InstanceSeconds"]),
		FailedMeanTotalSeconds:          mean(timings["FailedTotalSeconds"]),
		FailedMeanBackendSeconds:        mean(timings["FailedBackendSeconds"]),
		FailedMeanInstanceSeconds:       mean(timings["FailedInstanceSeconds"]),
		MaxTotalSeconds:                 max(timings["TotalSeconds"]),
		MaxBackendSeconds:               max(timings["BackendSeconds"]),
		MaxInstanceSeconds:              max(timings["InstanceSeconds"]),
		FailedMaxTotalSeconds:           max(timings["FailedTotalSeconds"]),
		FailedMaxBackendSeconds:         max(timings["FailedBackendSeconds"]),
		FailedMaxInstanceSeconds:        max(timings["FailedInstanceSeconds"]),
		MedianTotalSeconds:              median(timings["TotalSeconds"]),
		MedianBackendSeconds:            median(timings["BackendSeconds"]),
		MedianInstanceSeconds:           median(timings["InstanceSeconds"]),
		FailedMedianTotalSeconds:        median(timings["FailedTotalSeconds"]),
		FailedMedianBackendSeconds:      median(timings["FailedBackendSeconds"]),
		FailedMedianInstanceSeconds:     median(timings["FailedInstanceSeconds"]),
		P95TotalSeconds:                 percentile(timings["TotalSeconds"], 95),
		P95BackendSeconds:               percentile(timings["BackendSeconds"], 95),
		P95InstanceSeconds:              percentile(timings["InstanceSeconds"], 95),
		FailedP95TotalSeconds:           percentile(timings["FailedTotalSeconds"], 95),
		FailedP95BackendSeconds:         percentile(timings["FailedBackendSeconds"], 95),
		FailedP95InstanceSeconds:        percentile(timings["FailedInstanceSeconds"], 95),
	}
=======
		CountDistinctInstanceKeys:       len(names[InstanceKeys]),
		CountDistinctFailedInstanceKeys: len(names[FailedInstanceKeys]),
		FailedDiscoveries:               counters[FailedDiscoveries],
		SuccessfulDiscoveries:           counters[Discoveries],
		MeanTotalSeconds:                mean(timings[TotalSeconds]),
		MeanBackendSeconds:              mean(timings[BackendSeconds]),
		MeanInstanceSeconds:             mean(timings[InstanceSeconds]),
		FailedMeanTotalSeconds:          mean(timings[FailedTotalSeconds]),
		FailedMeanBackendSeconds:        mean(timings[FailedBackendSeconds]),
		FailedMeanInstanceSeconds:       mean(timings[FailedInstanceSeconds]),
		MaxTotalSeconds:                 max(timings[TotalSeconds]),
		MaxBackendSeconds:               max(timings[BackendSeconds]),
		MaxInstanceSeconds:              max(timings[InstanceSeconds]),
		FailedMaxTotalSeconds:           max(timings[FailedTotalSeconds]),
		FailedMaxBackendSeconds:         max(timings[FailedBackendSeconds]),
		FailedMaxInstanceSeconds:        max(timings[FailedInstanceSeconds]),
		MedianTotalSeconds:              median(timings[TotalSeconds]),
		MedianBackendSeconds:            median(timings[BackendSeconds]),
		MedianInstanceSeconds:           median(timings[InstanceSeconds]),
		FailedMedianTotalSeconds:        median(timings[FailedTotalSeconds]),
		FailedMedianBackendSeconds:      median(timings[FailedBackendSeconds]),
		FailedMedianInstanceSeconds:     median(timings[FailedInstanceSeconds]),
		P95TotalSeconds:                 percentile(timings[TotalSeconds], 95),
		P95BackendSeconds:               percentile(timings[BackendSeconds], 95),
		P95InstanceSeconds:              percentile(timings[InstanceSeconds], 95),
		FailedP95TotalSeconds:           percentile(timings[FailedTotalSeconds], 95),
		FailedP95BackendSeconds:         percentile(timings[FailedBackendSeconds], 95),
		FailedP95InstanceSeconds:        percentile(timings[FailedInstanceSeconds], 95),
	}
}

// AggregatedSince returns a large number of aggregated metrics
// based on the raw metrics collected since the given time.
func AggregatedSince(c *collection.Collection, t time.Time) (AggregatedDiscoveryMetrics, error) {
	results, err := c.Since(t)
	if err != nil {
		return AggregatedDiscoveryMetrics{}, err
	}

	return aggregate(results), nil
>>>>>>> refs/remotes/github/master
}
