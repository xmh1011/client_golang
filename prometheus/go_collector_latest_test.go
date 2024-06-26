// Copyright 2021 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build go1.17
// +build go1.17

package prometheus

import (
	"math"
	"reflect"
	"runtime"
	"runtime/metrics"
	"sync"
	"testing"

	dto "github.com/prometheus/client_model/go"

	"github.com/prometheus/client_golang/prometheus/internal"
)

var expectedRuntimeMetrics = map[string]string{
	"/cgo/go-to-c-calls:calls":                                  "go_cgo_go_to_c_calls_calls_total",
	"/cpu/classes/gc/mark/assist:cpu-seconds":                   "go_cpu_classes_gc_mark_assist_cpu_seconds_total",
	"/cpu/classes/gc/mark/dedicated:cpu-seconds":                "go_cpu_classes_gc_mark_dedicated_cpu_seconds_total",
	"/cpu/classes/gc/mark/idle:cpu-seconds":                     "go_cpu_classes_gc_mark_idle_cpu_seconds_total",
	"/cpu/classes/gc/pause:cpu-seconds":                         "go_cpu_classes_gc_pause_cpu_seconds_total",
	"/cpu/classes/gc/total:cpu-seconds":                         "go_cpu_classes_gc_total_cpu_seconds_total",
	"/cpu/classes/idle:cpu-seconds":                             "go_cpu_classes_idle_cpu_seconds_total",
	"/cpu/classes/scavenge/assist:cpu-seconds":                  "go_cpu_classes_scavenge_assist_cpu_seconds_total",
	"/cpu/classes/scavenge/background:cpu-seconds":              "go_cpu_classes_scavenge_background_cpu_seconds_total",
	"/cpu/classes/scavenge/total:cpu-seconds":                   "go_cpu_classes_scavenge_total_cpu_seconds_total",
	"/cpu/classes/total:cpu-seconds":                            "go_cpu_classes_total_cpu_seconds_total",
	"/cpu/classes/user:cpu-seconds":                             "go_cpu_classes_user_cpu_seconds_total",
	"/gc/cycles/automatic:gc-cycles":                            "go_gc_cycles_automatic_gc_cycles_total",
	"/gc/cycles/forced:gc-cycles":                               "go_gc_cycles_forced_gc_cycles_total",
	"/gc/cycles/total:gc-cycles":                                "go_gc_cycles_total_gc_cycles_total",
	"/gc/gogc:percent":                                          "go_gc_gogc_percent",
	"/gc/gomemlimit:bytes":                                      "go_gc_gomemlimit_bytes",
	"/gc/heap/allocs-by-size:bytes":                             "go_gc_heap_allocs_by_size_bytes",
	"/gc/heap/allocs:bytes":                                     "go_gc_heap_allocs_bytes_total",
	"/gc/heap/allocs:objects":                                   "go_gc_heap_allocs_objects_total",
	"/gc/heap/frees-by-size:bytes":                              "go_gc_heap_frees_by_size_bytes",
	"/gc/heap/frees:bytes":                                      "go_gc_heap_frees_bytes_total",
	"/gc/heap/frees:objects":                                    "go_gc_heap_frees_objects_total",
	"/gc/heap/goal:bytes":                                       "go_gc_heap_goal_bytes",
	"/gc/heap/live:bytes":                                       "go_gc_heap_live_bytes",
	"/gc/heap/objects:objects":                                  "go_gc_heap_objects_objects",
	"/gc/heap/tiny/allocs:objects":                              "go_gc_heap_tiny_allocs_objects_total",
	"/gc/limiter/last-enabled:gc-cycle":                         "go_gc_limiter_last_enabled_gc_cycle",
	"/gc/pauses:seconds":                                        "go_gc_pauses_seconds",
	"/gc/scan/globals:bytes":                                    "go_gc_scan_globals_bytes",
	"/gc/scan/heap:bytes":                                       "go_gc_scan_heap_bytes",
	"/gc/scan/stack:bytes":                                      "go_gc_scan_stack_bytes",
	"/gc/scan/total:bytes":                                      "go_gc_scan_total_bytes",
	"/gc/stack/starting-size:bytes":                             "go_gc_stack_starting_size_bytes",
	"/godebug/non-default-behavior/execerrdot:events":           "go_godebug_non_default_behavior_execerrdot_events_total",
	"/godebug/non-default-behavior/gocachehash:events":          "go_godebug_non_default_behavior_gocachehash_events_total",
	"/godebug/non-default-behavior/gocachetest:events":          "go_godebug_non_default_behavior_gocachetest_events_total",
	"/godebug/non-default-behavior/gocacheverify:events":        "go_godebug_non_default_behavior_gocacheverify_events_total",
	"/godebug/non-default-behavior/http2client:events":          "go_godebug_non_default_behavior_http2client_events_total",
	"/godebug/non-default-behavior/http2server:events":          "go_godebug_non_default_behavior_http2server_events_total",
	"/godebug/non-default-behavior/installgoroot:events":        "go_godebug_non_default_behavior_installgoroot_events_total",
	"/godebug/non-default-behavior/jstmpllitinterp:events":      "go_godebug_non_default_behavior_jstmpllitinterp_events_total",
	"/godebug/non-default-behavior/multipartmaxheaders:events":  "go_godebug_non_default_behavior_multipartmaxheaders_events_total",
	"/godebug/non-default-behavior/multipartmaxparts:events":    "go_godebug_non_default_behavior_multipartmaxparts_events_total",
	"/godebug/non-default-behavior/multipathtcp:events":         "go_godebug_non_default_behavior_multipathtcp_events_total",
	"/godebug/non-default-behavior/panicnil:events":             "go_godebug_non_default_behavior_panicnil_events_total",
	"/godebug/non-default-behavior/randautoseed:events":         "go_godebug_non_default_behavior_randautoseed_events_total",
	"/godebug/non-default-behavior/tarinsecurepath:events":      "go_godebug_non_default_behavior_tarinsecurepath_events_total",
	"/godebug/non-default-behavior/tlsmaxrsasize:events":        "go_godebug_non_default_behavior_tlsmaxrsasize_events_total",
	"/godebug/non-default-behavior/x509sha1:events":             "go_godebug_non_default_behavior_x509sha1_events_total",
	"/godebug/non-default-behavior/x509usefallbackroots:events": "go_godebug_non_default_behavior_x509usefallbackroots_events_total",
	"/godebug/non-default-behavior/zipinsecurepath:events":      "go_godebug_non_default_behavior_zipinsecurepath_events_total",
	"/memory/classes/heap/free:bytes":                           "go_memory_classes_heap_free_bytes",
	"/memory/classes/heap/objects:bytes":                        "go_memory_classes_heap_objects_bytes",
	"/memory/classes/heap/released:bytes":                       "go_memory_classes_heap_released_bytes",
	"/memory/classes/heap/stacks:bytes":                         "go_memory_classes_heap_stacks_bytes",
	"/memory/classes/heap/unused:bytes":                         "go_memory_classes_heap_unused_bytes",
	"/memory/classes/metadata/mcache/free:bytes":                "go_memory_classes_metadata_mcache_free_bytes",
	"/memory/classes/metadata/mcache/inuse:bytes":               "go_memory_classes_metadata_mcache_inuse_bytes",
	"/memory/classes/metadata/mspan/free:bytes":                 "go_memory_classes_metadata_mspan_free_bytes",
	"/memory/classes/metadata/mspan/inuse:bytes":                "go_memory_classes_metadata_mspan_inuse_bytes",
	"/memory/classes/metadata/other:bytes":                      "go_memory_classes_metadata_other_bytes",
	"/memory/classes/os-stacks:bytes":                           "go_memory_classes_os_stacks_bytes",
	"/memory/classes/other:bytes":                               "go_memory_classes_other_bytes",
	"/memory/classes/profiling/buckets:bytes":                   "go_memory_classes_profiling_buckets_bytes",
	"/memory/classes/total:bytes":                               "go_memory_classes_total_bytes",
	"/sched/gomaxprocs:threads":                                 "go_sched_gomaxprocs_threads",
	"/sched/goroutines:goroutines":                              "go_sched_goroutines_goroutines",
	"/sched/latencies:seconds":                                  "go_sched_latencies_seconds",
	"/sync/mutex/wait/total:seconds":                            "go_sync_mutex_wait_total_seconds_total",
}

const expectedRuntimeMetricsCardinality = 114

func TestRmForMemStats(t *testing.T) {
	if got, want := len(bestEffortLookupRM(rmForMemStats)), len(rmForMemStats); got != want {
		t.Errorf("got %d, want %d metrics", got, want)
	}
}

func expectedBaseMetrics() map[string]struct{} {
	metrics := map[string]struct{}{}
	b := newBaseGoCollector()
	for _, m := range []string{
		b.gcDesc.fqName,
		b.goInfoDesc.fqName,
		b.goroutinesDesc.fqName,
		b.gcLastTimeDesc.fqName,
		b.threadsDesc.fqName,
	} {
		metrics[m] = struct{}{}
	}
	return metrics
}

func addExpectedRuntimeMemStats(metrics map[string]struct{}) map[string]struct{} {
	for _, m := range goRuntimeMemStats() {
		metrics[m.desc.fqName] = struct{}{}
	}
	return metrics
}

func addExpectedRuntimeMetrics(metrics map[string]struct{}) map[string]struct{} {
	for _, m := range expectedRuntimeMetrics {
		metrics[m] = struct{}{}
	}
	return metrics
}

func TestGoCollector(t *testing.T) {
	for _, tcase := range []struct {
		collections       uint32
		expectedFQNameSet map[string]struct{}
	}{
		{
			collections:       0,
			expectedFQNameSet: expectedBaseMetrics(),
		},
		{
			collections:       goRuntimeMemStatsCollection,
			expectedFQNameSet: addExpectedRuntimeMemStats(expectedBaseMetrics()),
		},
		{
			collections:       goRuntimeMetricsCollection,
			expectedFQNameSet: addExpectedRuntimeMetrics(expectedBaseMetrics()),
		},
		{
			collections:       goRuntimeMemStatsCollection | goRuntimeMetricsCollection,
			expectedFQNameSet: addExpectedRuntimeMemStats(addExpectedRuntimeMetrics(expectedBaseMetrics())),
		},
	} {
		if ok := t.Run("", func(t *testing.T) {
			goMetrics := collectGoMetrics(t, tcase.collections)
			goMetricSet := make(map[string]Metric)
			for _, m := range goMetrics {
				goMetricSet[m.Desc().fqName] = m
			}

			for i := range goMetrics {
				name := goMetrics[i].Desc().fqName

				if _, ok := tcase.expectedFQNameSet[name]; !ok {
					t.Errorf("found unpexpected metric %s", name)
					continue
				}
			}

			// Now iterate over the expected metrics and look for removals.
			for expectedName := range tcase.expectedFQNameSet {
				if _, ok := goMetricSet[expectedName]; !ok {
					t.Errorf("missing expected metric %s in collection", expectedName)
					continue
				}
			}
		}); !ok {
			return
		}
	}
}

var sink interface{}

func TestBatchHistogram(t *testing.T) {
	goMetrics := collectGoMetrics(t, goRuntimeMetricsCollection)

	var mhist Metric
	for _, m := range goMetrics {
		if m.Desc().fqName == "go_gc_heap_allocs_by_size_bytes" {
			mhist = m
			break
		}
	}
	if mhist == nil {
		t.Fatal("failed to find metric to test")
	}
	hist, ok := mhist.(*batchHistogram)
	if !ok {
		t.Fatal("found metric is not a runtime/metrics histogram")
	}

	// Make a bunch of allocations then do another collection.
	//
	// The runtime/metrics API tries to reuse memory where possible,
	// so make sure that we didn't hang on to any of that memory in
	// hist.
	countsCopy := make([]uint64, len(hist.counts))
	copy(countsCopy, hist.counts)
	for i := 0; i < 100; i++ {
		sink = make([]byte, 128)
	}
	collectGoMetrics(t, defaultGoCollections)
	for i, v := range hist.counts {
		if v != countsCopy[i] {
			t.Error("counts changed during new collection")
			break
		}
	}

	// Get the runtime/metrics copy.
	s := []metrics.Sample{
		{Name: "/gc/heap/allocs-by-size:bytes"},
	}
	metrics.Read(s)
	rmHist := s[0].Value.Float64Histogram()
	wantBuckets := internal.RuntimeMetricsBucketsForUnit(rmHist.Buckets, "bytes")
	// runtime/metrics histograms always have a +Inf bucket and are lower
	// bound inclusive. In contrast, we have an implicit +Inf bucket and
	// are upper bound inclusive, so we can chop off the first bucket
	// (since the conversion to upper bound inclusive will shift all buckets
	// down one index) and the +Inf for the last bucket.
	wantBuckets = wantBuckets[1 : len(wantBuckets)-1]

	// Check to make sure the output proto makes sense.
	pb := &dto.Metric{}
	hist.Write(pb)

	if math.IsInf(pb.Histogram.Bucket[len(pb.Histogram.Bucket)-1].GetUpperBound(), +1) {
		t.Errorf("found +Inf bucket")
	}
	if got := len(pb.Histogram.Bucket); got != len(wantBuckets) {
		t.Errorf("got %d buckets in protobuf, want %d", got, len(wantBuckets))
	}
	for i, bucket := range pb.Histogram.Bucket {
		// runtime/metrics histograms are lower-bound inclusive, but we're
		// upper-bound inclusive. So just make sure the new inclusive upper
		// bound is somewhere close by (in some cases it's equal).
		wantBound := wantBuckets[i]
		if gotBound := *bucket.UpperBound; (wantBound-gotBound)/wantBound > 0.001 {
			t.Errorf("got bound %f, want within 0.1%% of %f", gotBound, wantBound)
		}
		// Make sure counts are cumulative. Because of the consistency guarantees
		// made by the runtime/metrics package, we're really not guaranteed to get
		// anything even remotely the same here.
		if i > 0 && *bucket.CumulativeCount < *pb.Histogram.Bucket[i-1].CumulativeCount {
			t.Error("cumulative counts are non-monotonic")
		}
	}
}

func collectGoMetrics(t *testing.T, enabledCollections uint32) []Metric {
	t.Helper()

	c := NewGoCollector(func(o *GoCollectorOptions) {
		o.EnabledCollections = enabledCollections
	}).(*goCollector)

	// Collect all metrics.
	ch := make(chan Metric)
	var wg sync.WaitGroup
	var metrics []Metric
	wg.Add(1)
	go func() {
		defer wg.Done()
		for metric := range ch {
			metrics = append(metrics, metric)
		}
	}()
	c.Collect(ch)
	close(ch)

	wg.Wait()

	return metrics
}

func TestMemStatsEquivalence(t *testing.T) {
	var msReal, msFake runtime.MemStats
	descs := bestEffortLookupRM(rmForMemStats)

	samples := make([]metrics.Sample, len(descs))
	samplesMap := make(map[string]*metrics.Sample)
	for i := range descs {
		samples[i].Name = descs[i].Name
		samplesMap[descs[i].Name] = &samples[i]
	}

	// Force a GC cycle to try to reach a clean slate.
	runtime.GC()

	// Populate msReal.
	runtime.ReadMemStats(&msReal)
	// Populate msFake and hope that no GC happened in between (:
	metrics.Read(samples)

	memStatsFromRM(&msFake, samplesMap)

	// Iterate over them and make sure they're somewhat close.
	msRealValue := reflect.ValueOf(msReal)
	msFakeValue := reflect.ValueOf(msFake)

	typ := msRealValue.Type()
	for i := 0; i < msRealValue.NumField(); i++ {
		fr := msRealValue.Field(i)
		ff := msFakeValue.Field(i)

		if typ.Field(i).Name == "PauseTotalNs" || typ.Field(i).Name == "LastGC" {
			// We don't use those fields for metrics,
			// thus we are not interested in having this filled.
			continue
		}
		switch fr.Kind() {
		// Fields which we are interested in are all uint64s.
		// The only float64 field GCCPUFraction is by design omitted.
		case reflect.Uint64:
			vr := fr.Interface().(uint64)
			vf := ff.Interface().(uint64)
			if float64(vr-vf)/float64(vf) > 0.05 {
				t.Errorf("wrong value for %s: got %d, want %d", typ.Field(i).Name, vf, vr)
			}
		}
	}
}

func TestExpectedRuntimeMetrics(t *testing.T) {
	goMetrics := collectGoMetrics(t, goRuntimeMetricsCollection)
	goMetricSet := make(map[string]Metric)
	for _, m := range goMetrics {
		goMetricSet[m.Desc().fqName] = m
	}

	descs := metrics.All()
	rmSet := make(map[string]struct{})
	// Iterate over runtime-reported descriptions to find new metrics.
	for i := range descs {
		rmName := descs[i].Name
		rmSet[rmName] = struct{}{}

		// expectedRuntimeMetrics depends on Go version.
		expFQName, ok := expectedRuntimeMetrics[rmName]
		if !ok {
			t.Errorf("found new runtime/metrics metric %s", rmName)
			_, _, _, ok := internal.RuntimeMetricsToProm(&descs[i])
			if !ok {
				t.Errorf("new metric has name that can't be converted, or has an unsupported Kind")
			}
			continue
		}
		_, ok = goMetricSet[expFQName]
		if !ok {
			t.Errorf("existing runtime/metrics metric %s (expected fq name %s) not collected", rmName, expFQName)
			continue
		}
	}

	// Now iterate over the expected metrics and look for removals.
	cardinality := 0
	for rmName, fqName := range expectedRuntimeMetrics {
		if _, ok := rmSet[rmName]; !ok {
			t.Errorf("runtime/metrics metric %s removed", rmName)
			continue
		}
		if _, ok := goMetricSet[fqName]; !ok {
			t.Errorf("runtime/metrics metric %s not appearing under expected name %s", rmName, fqName)
			continue
		}

		// While we're at it, check to make sure expected cardinality lines
		// up, but at the point of the protobuf write to get as close to the
		// real deal as possible.
		//
		// Note that we filter out non-runtime/metrics metrics here, because
		// those are manually managed.
		var m dto.Metric
		if err := goMetricSet[fqName].Write(&m); err != nil {
			t.Errorf("writing metric %s: %v", fqName, err)
			continue
		}
		// N.B. These are the only fields populated by runtime/metrics metrics specifically.
		// Other fields are populated by e.g. GCStats metrics.
		switch {
		case m.Counter != nil:
			fallthrough
		case m.Gauge != nil:
			cardinality++
		case m.Histogram != nil:
			cardinality += len(m.Histogram.Bucket) + 3 // + sum, count, and +inf
		default:
			t.Errorf("unexpected protobuf structure for metric %s", fqName)
		}
	}

	if t.Failed() {
		t.Log("a new Go version may have been detected, please run")
		t.Log("\tgo run gen_go_collector_metrics_set.go go1.X")
		t.Log("where X is the Go version you are currently using")
	}

	expectCardinality := expectedRuntimeMetricsCardinality
	if cardinality != expectCardinality {
		t.Errorf("unexpected cardinality for runtime/metrics metrics: got %d, want %d", cardinality, expectCardinality)
	}
}

func TestGoCollectorConcurrency(t *testing.T) {
	c := NewGoCollector().(*goCollector)

	// Set up multiple goroutines to Collect from the
	// same GoCollector. In race mode with GOMAXPROCS > 1,
	// this test should fail often if Collect is not
	// concurrent-safe.
	for i := 0; i < 4; i++ {
		go func() {
			ch := make(chan Metric)
			go func() {
				// Drain all metrics received until the
				// channel is closed.
				for range ch {
				}
			}()
			c.Collect(ch)
			close(ch)
		}()
	}
}
