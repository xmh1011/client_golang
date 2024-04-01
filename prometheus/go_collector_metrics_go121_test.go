// Code generated by gen_go_collector_metrics_set.go; DO NOT EDIT.
//go:generate go run gen_go_collector_metrics_set.go go1.21

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package prometheus

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