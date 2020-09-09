package api


type node []struct {
		Partitions    []interface{} `json:"partitions"`
		OsPid         string        `json:"os_pid"`
		FdTotal       int           `json:"fd_total"`
		SocketsTotal  int           `json:"sockets_total"`
		MemLimit      int64         `json:"mem_limit"`
		MemAlarm      bool          `json:"mem_alarm"`
		DiskFreeLimit int           `json:"disk_free_limit"`
		DiskFreeAlarm bool          `json:"disk_free_alarm"`
		ProcTotal     int           `json:"proc_total"`
		RatesMode     string        `json:"rates_mode"`
		Uptime        int64         `json:"uptime"`
		RunQueue      int           `json:"run_queue"`
		Processors    int           `json:"processors"`
		ExchangeTypes []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Enabled     bool   `json:"enabled"`
		} `json:"exchange_types"`
		AuthMechanisms []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Enabled     bool   `json:"enabled"`
		} `json:"auth_mechanisms"`
		Applications []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Version     string `json:"version"`
		} `json:"applications"`
		Contexts []struct {
			Description string `json:"description"`
			Path        string `json:"path"`
			CowboyOpts  string `json:"cowboy_opts"`
			Port        string `json:"port"`
		} `json:"contexts"`
		LogFiles               []string      `json:"log_files"`
		DbDir                  string        `json:"db_dir"`
		ConfigFiles            []interface{} `json:"config_files"`
		NetTicktime            int           `json:"net_ticktime"`
		EnabledPlugins         []string      `json:"enabled_plugins"`
		MemCalculationStrategy string        `json:"mem_calculation_strategy"`
		RaOpenFileMetrics      struct {
			RaLogWal           int `json:"ra_log_wal"`
			RaLogSegmentWriter int `json:"ra_log_segment_writer"`
		} `json:"ra_open_file_metrics"`
		Name           string `json:"name"`
		Type           string `json:"type"`
		Running        bool   `json:"running"`
		MemUsed        int    `json:"mem_used"`
		MemUsedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"mem_used_details"`
		FdUsed        int `json:"fd_used"`
		FdUsedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"fd_used_details"`
		SocketsUsed        int `json:"sockets_used"`
		SocketsUsedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"sockets_used_details"`
		ProcUsed        int `json:"proc_used"`
		ProcUsedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"proc_used_details"`
		DiskFree        int64 `json:"disk_free"`
		DiskFreeDetails struct {
			Rate float64 `json:"rate"`
		} `json:"disk_free_details"`
		GcNum        int `json:"gc_num"`
		GcNumDetails struct {
			Rate float64 `json:"rate"`
		} `json:"gc_num_details"`
		GcBytesReclaimed        int64 `json:"gc_bytes_reclaimed"`
		GcBytesReclaimedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"gc_bytes_reclaimed_details"`
		ContextSwitches        int `json:"context_switches"`
		ContextSwitchesDetails struct {
			Rate float64 `json:"rate"`
		} `json:"context_switches_details"`
		IoReadCount        int `json:"io_read_count"`
		IoReadCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_read_count_details"`
		IoReadBytes        int `json:"io_read_bytes"`
		IoReadBytesDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_read_bytes_details"`
		IoReadAvgTime        float64 `json:"io_read_avg_time"`
		IoReadAvgTimeDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_read_avg_time_details"`
		IoWriteCount        int `json:"io_write_count"`
		IoWriteCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_write_count_details"`
		IoWriteBytes        int `json:"io_write_bytes"`
		IoWriteBytesDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_write_bytes_details"`
		IoWriteAvgTime        float64 `json:"io_write_avg_time"`
		IoWriteAvgTimeDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_write_avg_time_details"`
		IoSyncCount        int `json:"io_sync_count"`
		IoSyncCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_sync_count_details"`
		IoSyncAvgTime        float64 `json:"io_sync_avg_time"`
		IoSyncAvgTimeDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_sync_avg_time_details"`
		IoSeekCount        int `json:"io_seek_count"`
		IoSeekCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_seek_count_details"`
		IoSeekAvgTime        float64 `json:"io_seek_avg_time"`
		IoSeekAvgTimeDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_seek_avg_time_details"`
		IoReopenCount        int `json:"io_reopen_count"`
		IoReopenCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_reopen_count_details"`
		MnesiaRAMTxCount        int `json:"mnesia_ram_tx_count"`
		MnesiaRAMTxCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"mnesia_ram_tx_count_details"`
		MnesiaDiskTxCount        int `json:"mnesia_disk_tx_count"`
		MnesiaDiskTxCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"mnesia_disk_tx_count_details"`
		MsgStoreReadCount        int `json:"msg_store_read_count"`
		MsgStoreReadCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"msg_store_read_count_details"`
		MsgStoreWriteCount        int `json:"msg_store_write_count"`
		MsgStoreWriteCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"msg_store_write_count_details"`
		QueueIndexJournalWriteCount        int `json:"queue_index_journal_write_count"`
		QueueIndexJournalWriteCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"queue_index_journal_write_count_details"`
		QueueIndexWriteCount        int `json:"queue_index_write_count"`
		QueueIndexWriteCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"queue_index_write_count_details"`
		QueueIndexReadCount        int `json:"queue_index_read_count"`
		QueueIndexReadCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"queue_index_read_count_details"`
		IoFileHandleOpenAttemptCount        int `json:"io_file_handle_open_attempt_count"`
		IoFileHandleOpenAttemptCountDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_file_handle_open_attempt_count_details"`
		IoFileHandleOpenAttemptAvgTime        float64 `json:"io_file_handle_open_attempt_avg_time"`
		IoFileHandleOpenAttemptAvgTimeDetails struct {
			Rate float64 `json:"rate"`
		} `json:"io_file_handle_open_attempt_avg_time_details"`
		ConnectionCreated        int `json:"connection_created"`
		ConnectionCreatedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"connection_created_details"`
		ConnectionClosed        int `json:"connection_closed"`
		ConnectionClosedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"connection_closed_details"`
		ChannelCreated        int `json:"channel_created"`
		ChannelCreatedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"channel_created_details"`
		ChannelClosed        int `json:"channel_closed"`
		ChannelClosedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"channel_closed_details"`
		QueueDeclared        int `json:"queue_declared"`
		QueueDeclaredDetails struct {
			Rate float64 `json:"rate"`
		} `json:"queue_declared_details"`
		QueueCreated        int `json:"queue_created"`
		QueueCreatedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"queue_created_details"`
		QueueDeleted        int `json:"queue_deleted"`
		QueueDeletedDetails struct {
			Rate float64 `json:"rate"`
		} `json:"queue_deleted_details"`
		ClusterLinks []struct {
			Stats struct {
				SendBytes        int64 `json:"send_bytes"`
				SendBytesDetails struct {
					Rate float64 `json:"rate"`
				} `json:"send_bytes_details"`
				RecvBytes        int64 `json:"recv_bytes"`
				RecvBytesDetails struct {
					Rate float64 `json:"rate"`
				} `json:"recv_bytes_details"`
			} `json:"stats"`
			Name      string `json:"name"`
			PeerAddr  string `json:"peer_addr"`
			PeerPort  int    `json:"peer_port"`
			SockAddr  string `json:"sock_addr"`
			SockPort  int    `json:"sock_port"`
			RecvBytes int64  `json:"recv_bytes"`
			SendBytes int64  `json:"send_bytes"`
		} `json:"cluster_links"`
		MetricsGcQueueLength struct {
			ConnectionClosed       int `json:"connection_closed"`
			ChannelClosed          int `json:"channel_closed"`
			ConsumerDeleted        int `json:"consumer_deleted"`
			ExchangeDeleted        int `json:"exchange_deleted"`
			QueueDeleted           int `json:"queue_deleted"`
			VhostDeleted           int `json:"vhost_deleted"`
			NodeNodeDeleted        int `json:"node_node_deleted"`
			ChannelConsumerDeleted int `json:"channel_consumer_deleted"`
		} `json:"metrics_gc_queue_length"`
	}