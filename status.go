package main

import (
	"encoding/json"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/pkg/errors"
)

// A special key to read the status json. https://apple.github.io/foundationdb/mr-status.html
var jsonKey = append([]byte{255, 255}, []byte("/status/json")...)

func getMetrics(db fdb.Database) (FdbStatus, error) {
	type ret struct {
		json []byte
		rv   int64
	}
	rv, err := db.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		rv := tr.GetReadVersion().MustGet()
		json := tr.Get(fdb.Key(jsonKey)).MustGet()
		return &ret{json: json, rv: rv}, nil
	})
	if err != nil {
		return FdbStatus{}, errors.Wrap(err, "cannot get status")
	}

	r := rv.(*ret)
	var status FdbStatus
	if err := json.Unmarshal(r.json, &status); err != nil {
		return FdbStatus{}, errors.Wrap(err, "cannot decode json")
	}
	status.ReadVersion = r.rv

	return status, nil
}

const (
	StorageRoleMetrics = "storage"
	LogRoleMetrics     = "log"
)

// These structs were principally generated by goland from the status json.
// The json format is described here:
// https://apple.github.io/foundationdb/mr-status.html#json-format
type FdbMachine struct {
	Id                  string
	Address             string `json:"address"`
	ContributingWorkers int64  `json:"contributing_workers"`
	Cpu                 struct {
		LogicalCoreUtilization float64 `json:"logical_core_utilization"`
	} `json:"cpu"`
	Excluded bool `json:"excluded"`
	Locality struct {
		DataHall  string `json:"data_hall"`
		Machineid string `json:"machineid"`
		Processid string `json:"processid"`
		Zoneid    string `json:"zoneid"`
	} `json:"locality"`
	MachineId string `json:"machine_id"`
	Memory    struct {
		CommittedBytes int64 `json:"committed_bytes"`
		FreeBytes      int64 `json:"free_bytes"`
		TotalBytes     int64 `json:"total_bytes"`
	} `json:"memory"`
	Network struct {
		MegabitsReceived struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_received"`
		MegabitsSent struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_sent"`
		TcpSegmentsRetransmitted struct {
			Hz float64 `json:"hz"`
		} `json:"tcp_segments_retransmitted"`
	} `json:"network"`
}

type FdbRole struct {
	Id           string `json:"id,omitempty"`
	Role         string `json:"role"`
	BytesQueried struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"bytes_queried,omitempty"`
	DataLag struct {
		Seconds  float64 `json:"seconds"`
		Versions int64   `json:"versions"`
	} `json:"data_lag,omitempty"`
	DataVersion   int64 `json:"data_version,omitempty"`
	DurabilityLag struct {
		Seconds  float64 `json:"seconds"`
		Versions int64   `json:"versions"`
	} `json:"durability_lag,omitempty"`
	DurableBytes struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"durable_bytes,omitempty"`
	DurableVersion  int64 `json:"durable_version,omitempty"`
	FetchedVersions struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"fetched_versions,omitempty"`
	FetchesFromLogs struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"fetches_from_logs,omitempty"`
	FinishedQueries struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"finished_queries,omitempty"`
	InputBytes struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"input_bytes,omitempty"`
	KeysQueried struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"keys_queried,omitempty"`
	KvstoreAvailableBytes int64 `json:"kvstore_available_bytes,omitempty"`
	KvstoreFreeBytes      int64 `json:"kvstore_free_bytes,omitempty"`
	KvstoreInlineKeys     int64 `json:"kvstore_inline_keys,omitempty"`
	KvstoreTotalBytes     int64 `json:"kvstore_total_bytes,omitempty"`
	KvstoreTotalNodes     int64 `json:"kvstore_total_nodes,omitempty"`
	KvstoreTotalSize      int64 `json:"kvstore_total_size,omitempty"`
	KvstoreUsedBytes      int64 `json:"kvstore_used_bytes,omitempty"`
	LocalRate             int64 `json:"local_rate,omitempty"`
	LowPriorityQueries    struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"low_priority_queries,omitempty"`
	MutationBytes struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"mutation_bytes,omitempty"`
	Mutations struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"mutations,omitempty"`
	QueryQueueMax         int64 `json:"query_queue_max,omitempty"`
	ReadLatencyStatistics struct {
		Count  int64   `json:"count"`
		Max    float64 `json:"max"`
		Mean   float64 `json:"mean"`
		Median float64 `json:"median"`
		Min    float64 `json:"min"`
		P25    float64 `json:"p25"`
		P90    float64 `json:"p90"`
		P95    float64 `json:"p95"`
		P99    float64 `json:"p99"`
		P999   float64 `json:"p99.9"`
	} `json:"read_latency_statistics,omitempty"`
	StorageMetadata struct {
		CreatedTimeDatetime  string  `json:"created_time_datetime"`
		CreatedTimeTimestamp float64 `json:"created_time_timestamp"`
	} `json:"storage_metadata,omitempty"`
	StoredBytes  int64 `json:"stored_bytes,omitempty"`
	TotalQueries struct {
		Counter   int64   `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"total_queries,omitempty"`
	QueueDiskAvailableBytes int64 `json:"queue_disk_available_bytes"`
	QueueDiskFreeBytes      int64 `json:"queue_disk_free_bytes"`
	QueueDiskTotalBytes     int64 `json:"queue_disk_total_bytes"`
	QueueDiskUsedBytes      int64 `json:"queue_disk_used_bytes"`
}

type FdbProcess struct {
	Address     string `json:"address"`
	ClassSource string `json:"class_source"`
	ClassType   string `json:"class_type"`
	CommandLine string `json:"command_line"`
	Cpu         struct {
		UsageCores float64 `json:"usage_cores"`
	} `json:"cpu"`
	Disk struct {
		Busy      float64 `json:"busy"`
		FreeBytes int64   `json:"free_bytes"`
		Reads     struct {
			Counter int64   `json:"counter"`
			Hz      float64 `json:"hz"`
			Sectors float64 `json:"sectors"`
		} `json:"reads"`
		TotalBytes int64 `json:"total_bytes"`
		Writes     struct {
			Counter int64   `json:"counter"`
			Hz      float64 `json:"hz"`
			Sectors float64 `json:"sectors"`
		} `json:"writes"`
	} `json:"disk"`
	Excluded    bool   `json:"excluded"`
	FaultDomain string `json:"fault_domain"`
	Locality    struct {
		DataHall  string `json:"data_hall"`
		Machineid string `json:"machineid"`
		Processid string `json:"processid"`
		Zoneid    string `json:"zoneid"`
	} `json:"locality"`
	MachineId string `json:"machine_id"`
	Memory    struct {
		AvailableBytes        int64 `json:"available_bytes"`
		LimitBytes            int64 `json:"limit_bytes"`
		RssBytes              int64 `json:"rss_bytes"`
		UnusedAllocatedMemory int64 `json:"unused_allocated_memory"`
		UsedBytes             int64 `json:"used_bytes"`
	} `json:"memory"`
	Messages []interface{} `json:"messages"`
	Network  struct {
		ConnectionErrors struct {
			Hz float64 `json:"hz"`
		} `json:"connection_errors"`
		ConnectionsClosed struct {
			Hz float64 `json:"hz"`
		} `json:"connections_closed"`
		ConnectionsEstablished struct {
			Hz float64 `json:"hz"`
		} `json:"connections_established"`
		CurrentConnections int64 `json:"current_connections"`
		MegabitsReceived   struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_received"`
		MegabitsSent struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_sent"`
		TlsPolicyFailures struct {
			Hz float64 `json:"hz"`
		} `json:"tls_policy_failures"`
	} `json:"network"`
	Roles         []FdbRole `json:"roles"`
	RunLoopBusy   float64   `json:"run_loop_busy"`
	UptimeSeconds float64   `json:"uptime_seconds"`
	Version       string    `json:"version"`
}

// FdbStatus was generated by goland from the status json.
// The json format is described here:
// https://apple.github.io/foundationdb/mr-status.html#json-format
type FdbStatus struct {
	ReadVersion int64
	Client      struct {
		ClusterFile struct {
			Path     string `json:"path"`
			UpToDate bool   `json:"up_to_date"`
		} `json:"cluster_file"`
		Coordinators struct {
			Coordinators []struct {
				Address   string `json:"address"`
				Protocol  string `json:"protocol"`
				Reachable bool   `json:"reachable"`
			} `json:"coordinators"`
			QuorumReachable bool `json:"quorum_reachable"`
		} `json:"coordinators"`
		DatabaseStatus struct {
			Available bool `json:"available"`
			Healthy   bool `json:"healthy"`
		} `json:"database_status"`
		Messages  []interface{} `json:"messages"`
		Timestamp int64         `json:"timestamp"`
	} `json:"client"`
	Cluster struct {
		ActivePrimaryDc string `json:"active_primary_dc"`
		ActiveTssCount  int64  `json:"active_tss_count"`
		BounceImpact    struct {
			CanCleanBounce bool `json:"can_clean_bounce"`
		} `json:"bounce_impact"`
		Clients struct {
			Count             int64 `json:"count"`
			SupportedVersions []struct {
				ClientVersion    string `json:"client_version"`
				ConnectedClients []struct {
					Address  string `json:"address"`
					LogGroup string `json:"log_group"`
				} `json:"connected_clients"`
				Count              int64 `json:"count"`
				MaxProtocolClients []struct {
					Address  string `json:"address"`
					LogGroup string `json:"log_group"`
				} `json:"max_protocol_clients"`
				MaxProtocolCount int64  `json:"max_protocol_count"`
				ProtocolVersion  string `json:"protocol_version"`
				SourceVersion    string `json:"source_version"`
			} `json:"supported_versions"`
		} `json:"clients"`
		ClusterControllerTimestamp int64 `json:"cluster_controller_timestamp"`
		Configuration              struct {
			BackupWorkerEnabled int64 `json:"backup_worker_enabled"`
			BlobGranulesEnabled int64 `json:"blob_granules_enabled"`
			CoordinatorsCount   int64 `json:"coordinators_count"`
			ExcludedServers     []struct {
				Address string `json:"address"`
			} `json:"excluded_servers"`
			LogSpill                       int64  `json:"log_spill"`
			PerpetualStorageWiggle         int64  `json:"perpetual_storage_wiggle"`
			PerpetualStorageWiggleEngine   string `json:"perpetual_storage_wiggle_engine"`
			PerpetualStorageWiggleLocality string `json:"perpetual_storage_wiggle_locality"`
			RedundancyMode                 string `json:"redundancy_mode"`
			StorageEngine                  string `json:"storage_engine"`
			StorageMigrationType           string `json:"storage_migration_type"`
			TenantMode                     string `json:"tenant_mode"`
			UsableRegions                  int64  `json:"usable_regions"`
		} `json:"configuration"`
		ConnectionString string `json:"connection_string"`
		Data             struct {
			AveragePartitionSizeBytes             int64 `json:"average_partition_size_bytes"`
			LeastOperatingSpaceBytesLogServer     int64 `json:"least_operating_space_bytes_log_server"`
			LeastOperatingSpaceBytesStorageServer int64 `json:"least_operating_space_bytes_storage_server"`
			MovingData                            struct {
				HighestPriority   int64 `json:"highest_priority"`
				InFlightBytes     int64 `json:"in_flight_bytes"`
				InQueueBytes      int64 `json:"in_queue_bytes"`
				TotalWrittenBytes int64 `json:"total_written_bytes"`
			} `json:"moving_data"`
			PartitionsCount int64 `json:"partitions_count"`
			State           struct {
				Healthy              bool   `json:"healthy"`
				MinReplicasRemaining int64  `json:"min_replicas_remaining"`
				Name                 string `json:"name"`
			} `json:"state"`
			SystemKvSizeBytes int64 `json:"system_kv_size_bytes"`
			TeamTrackers      []struct {
				InFlightBytes int64 `json:"in_flight_bytes"`
				Primary       bool  `json:"primary"`
				State         struct {
					Healthy              bool   `json:"healthy"`
					MinReplicasRemaining int64  `json:"min_replicas_remaining"`
					Name                 string `json:"name"`
				} `json:"state"`
				UnhealthyServers int64 `json:"unhealthy_servers"`
			} `json:"team_trackers"`
			TotalDiskUsedBytes int64 `json:"total_disk_used_bytes"`
			TotalKvSizeBytes   int64 `json:"total_kv_size_bytes"`
		} `json:"data"`
		DatabaseAvailable bool `json:"database_available"`
		DatabaseLockState struct {
			Locked bool `json:"locked"`
		} `json:"database_lock_state"`
		DatacenterLag struct {
			Seconds  float64 `json:"seconds"`
			Versions int64   `json:"versions"`
		} `json:"datacenter_lag"`
		DegradedProcesses int64 `json:"degraded_processes"`
		FaultTolerance    struct {
			MaxZoneFailuresWithoutLosingAvailability int64 `json:"max_zone_failures_without_losing_availability"`
			MaxZoneFailuresWithoutLosingData         int64 `json:"max_zone_failures_without_losing_data"`
		} `json:"fault_tolerance"`
		FullReplication         bool          `json:"full_replication"`
		Generation              int64         `json:"generation"`
		IncompatibleConnections []interface{} `json:"incompatible_connections"`
		LatencyProbe            struct {
			BatchPriorityTransactionStartSeconds     float64 `json:"batch_priority_transaction_start_seconds"`
			CommitSeconds                            float64 `json:"commit_seconds"`
			ImmediatePriorityTransactionStartSeconds float64 `json:"immediate_priority_transaction_start_seconds"`
			ReadSeconds                              float64 `json:"read_seconds"`
			TransactionStartSeconds                  float64 `json:"transaction_start_seconds"`
		} `json:"latency_probe"`
		Layers struct {
			Valid  bool `json:"_valid"`
			Backup struct {
				BlobRecentIo struct {
					BytesPerSecond     float64 `json:"bytes_per_second"`
					BytesSent          int64   `json:"bytes_sent"`
					RequestsFailed     int64   `json:"requests_failed"`
					RequestsSuccessful int64   `json:"requests_successful"`
				} `json:"blob_recent_io"`
				Instances map[string]struct {
					BlobStats struct {
						Recent struct {
							BytesPerSecond     float64 `json:"bytes_per_second"`
							BytesSent          int64   `json:"bytes_sent"`
							RequestsFailed     int64   `json:"requests_failed"`
							RequestsSuccessful int64   `json:"requests_successful"`
						} `json:"recent"`
						Total struct {
							BytesSent          int64 `json:"bytes_sent"`
							RequestsFailed     int64 `json:"requests_failed"`
							RequestsSuccessful int64 `json:"requests_successful"`
						} `json:"total"`
					} `json:"blob_stats"`
					ConfiguredWorkers    int64   `json:"configured_workers"`
					Id                   string  `json:"id"`
					LastUpdated          float64 `json:"last_updated"`
					MainThreadCpuSeconds float64 `json:"main_thread_cpu_seconds"`
					MemoryUsage          int64   `json:"memory_usage"`
					ProcessCpuSeconds    float64 `json:"process_cpu_seconds"`
					ResidentSize         int64   `json:"resident_size"`
					Version              string  `json:"version"`
				} `json:"instances"`
				InstancesRunning int64   `json:"instances_running"`
				LastUpdated      float64 `json:"last_updated"`
				Paused           bool    `json:"paused"`
				Tags             struct {
				} `json:"tags"`
				TotalWorkers int64 `json:"total_workers"`
			} `json:"backup"`
		} `json:"layers"`
		Logs []struct {
			BeginVersion      int64 `json:"begin_version"`
			Current           bool  `json:"current"`
			Epoch             int64 `json:"epoch"`
			LogFaultTolerance int64 `json:"log_fault_tolerance"`
			LogInterfaces     []struct {
				Address string `json:"address"`
				Healthy bool   `json:"healthy"`
				Id      string `json:"id"`
			} `json:"log_interfaces"`
			LogReplicationFactor int64 `json:"log_replication_factor"`
			LogWriteAntiQuorum   int64 `json:"log_write_anti_quorum"`
			PossiblyLosingData   bool  `json:"possibly_losing_data"`
		} `json:"logs"`
		Machines  map[string]FdbMachine `json:"machines"`
		Messages  []interface{}         `json:"messages"`
		PageCache struct {
			LogHitRate     float64 `json:"log_hit_rate"`
			StorageHitRate float64 `json:"storage_hit_rate"`
		} `json:"page_cache"`
		Processes       map[string]FdbProcess `json:"processes"`
		ProtocolVersion string                `json:"protocol_version"`
		Qos             struct {
			BatchPerformanceLimitedBy struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				ReasonId    int64  `json:"reason_id"`
			} `json:"batch_performance_limited_by"`
			BatchReleasedTransactionsPerSecond float64 `json:"batch_released_transactions_per_second"`
			BatchTransactionsPerSecondLimit    float64 `json:"batch_transactions_per_second_limit"`
			LimitingDataLagStorageServer       struct {
				Seconds  float64 `json:"seconds"`
				Versions int64   `json:"versions"`
			} `json:"limiting_data_lag_storage_server"`
			LimitingDurabilityLagStorageServer struct {
				Seconds  float64 `json:"seconds"`
				Versions int64   `json:"versions"`
			} `json:"limiting_durability_lag_storage_server"`
			LimitingQueueBytesStorageServer int64 `json:"limiting_queue_bytes_storage_server"`
			PerformanceLimitedBy            struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				ReasonId    int64  `json:"reason_id"`
			} `json:"performance_limited_by"`
			ReleasedTransactionsPerSecond float64 `json:"released_transactions_per_second"`
			ThrottledTags                 struct {
				Auto struct {
					BusyRead        int64 `json:"busy_read"`
					BusyWrite       int64 `json:"busy_write"`
					Count           int64 `json:"count"`
					RecommendedOnly int64 `json:"recommended_only"`
				} `json:"auto"`
				Manual struct {
					Count int64 `json:"count"`
				} `json:"manual"`
			} `json:"throttled_tags"`
			TransactionsPerSecondLimit float64 `json:"transactions_per_second_limit"`
			WorstDataLagStorageServer  struct {
				Seconds  float64 `json:"seconds"`
				Versions int64   `json:"versions"`
			} `json:"worst_data_lag_storage_server"`
			WorstDurabilityLagStorageServer struct {
				Seconds  float64 `json:"seconds"`
				Versions int64   `json:"versions"`
			} `json:"worst_durability_lag_storage_server"`
			WorstQueueBytesLogServer     int64 `json:"worst_queue_bytes_log_server"`
			WorstQueueBytesStorageServer int64 `json:"worst_queue_bytes_storage_server"`
		} `json:"qos"`
		RecoveryState struct {
			ActiveGenerations         int64   `json:"active_generations"`
			Description               string  `json:"description"`
			Name                      string  `json:"name"`
			SecondsSinceLastRecovered float64 `json:"seconds_since_last_recovered"`
		} `json:"recovery_state"`
		Workload struct {
			Bytes struct {
				Read struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"read"`
				Written struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"written"`
			} `json:"bytes"`
			Keys struct {
				Read struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"read"`
			} `json:"keys"`
			Operations struct {
				LocationRequests struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"location_requests"`
				LowPriorityReads struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"low_priority_reads"`
				MemoryErrors struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"memory_errors"`
				ReadRequests struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"read_requests"`
				Reads struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"reads"`
				Writes struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"writes"`
			} `json:"operations"`
			Transactions struct {
				Committed struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"committed"`
				Conflicted struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"conflicted"`
				RejectedForQueuedTooLong struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"rejected_for_queued_too_long"`
				Started struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"started"`
				StartedBatchPriority struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"started_batch_priority"`
				StartedDefaultPriority struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"started_default_priority"`
				StartedImmediatePriority struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"started_immediate_priority"`
			} `json:"transactions"`
		} `json:"workload"`
	} `json:"cluster"`
}
