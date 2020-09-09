package api


//   json /api/overview
type overview struct {
	ManagementVersion string `json:"management_version"`
	RatesMode string `json:"rates_mode"`
	SampleRetentionPolicies SampleRetentionPolicies `json:"sample_retention_policies"`
	ExchangeTypes []ExchangeTypes `json:"exchange_types"`
	ProductVersion string `json:"product_version"`
	ProductName string `json:"product_name"`
	RabbitmqVersion string `json:"rabbitmq_version"`
	ClusterName string `json:"cluster_name"`
	ErlangVersion string `json:"erlang_version"`
	ErlangFullVersion string `json:"erlang_full_version"`
	DisableStats bool `json:"disable_stats"`
	EnableQueueTotals bool `json:"enable_queue_totals"`
	MessageStats MessageStats `json:"message_stats"`
	ChurnRates ChurnRates `json:"churn_rates"`
	QueueTotals QueueTotals `json:"queue_totals"`
	ObjectTotals ObjectTotals `json:"object_totals"`
	StatisticsDbEventQueue int `json:"statistics_db_event_queue"`
	Node string `json:"node"`
	Listeners []Listeners `json:"listeners"`
	Contexts []Contexts `json:"contexts"`
}
type SampleRetentionPolicies struct {
	Global []int `json:"global"`
	Basic []int `json:"basic"`
	Detailed []int `json:"detailed"`
}
type ExchangeTypes struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Enabled bool `json:"enabled"`
}
type MessageStats struct {
	Ack int `json:"ack"`
	AckDetails struct {
		Rate float64 `json:"rate"`
	} `json:"ack_details"`
	Confirm int `json:"confirm"`
	ConfirmDetails struct {
		Rate float64 `json:"rate"`
	} `json:"confirm_details"`
	Deliver int `json:"deliver"`
	DeliverDetails struct {
		Rate float64 `json:"rate"`
	} `json:"deliver_details"`
	DeliverGet int `json:"deliver_get"`
	DeliverGetDetails struct {
		Rate float64 `json:"rate"`
	} `json:"deliver_get_details"`
	DeliverNoAck int `json:"deliver_no_ack"`
	DeliverNoAckDetails struct {
		Rate float64 `json:"rate"`
	} `json:"deliver_no_ack_details"`
	DiskReads int `json:"disk_reads"`
	DiskReadsDetails struct {
		Rate float64 `json:"rate"`
	} `json:"disk_reads_details"`
	DiskWrites int `json:"disk_writes"`
	DiskWritesDetails struct {
		Rate float64 `json:"rate"`
	} `json:"disk_writes_details"`
	DropUnroutable int `json:"drop_unroutable"`
	DropUnroutableDetails struct {
		Rate float64 `json:"rate"`
	} `json:"drop_unroutable_details"`
	Get int `json:"get"`
	GetDetails struct {
		Rate float64 `json:"rate"`
	} `json:"get_details"`
	GetEmpty int `json:"get_empty"`
	GetEmptyDetails struct {
		Rate float64 `json:"rate"`
	} `json:"get_empty_details"`
	GetNoAck int `json:"get_no_ack"`
	GetNoAckDetails struct {
		Rate float64 `json:"rate"`
	} `json:"get_no_ack_details"`
	Publish int `json:"publish"`
	PublishDetails struct {
		Rate float64 `json:"rate"`
	} `json:"publish_details"`
	Redeliver int `json:"redeliver"`
	RedeliverDetails struct {
		Rate float64 `json:"rate"`
	} `json:"redeliver_details"`
	ReturnUnroutable int `json:"return_unroutable"`
	ReturnUnroutableDetails struct {
		Rate float64 `json:"rate"`
	} `json:"return_unroutable_details"`
}
type ChannelClosedDetails struct {
	Rate float64 `json:"rate"`
}
/*type ChannelCreatedDetails struct {
	Rate float64 `json:"rate"`
}*/
type ConnectionClosedDetails struct {
	Rate float64 `json:"rate"`
}
type ConnectionCreatedDetails struct {
	Rate float64 `json:"rate"`
}
type QueueCreatedDetails struct {
	Rate float64 `json:"rate"`
}
type QueueDeclaredDetails struct {
	Rate float64 `json:"rate"`
}
type QueueDeletedDetails struct {
	Rate float64 `json:"rate"`
}
type ChurnRates struct {
	ChannelClosed int `json:"channel_closed"`
	ChannelClosedDetails ChannelClosedDetails `json:"channel_closed_details"`
	ChannelCreated int `json:"channel_created"`
	ChannelCreatedDetails interface{} `json:"channel_created_details"`
	ConnectionClosed int `json:"connection_closed"`
	ConnectionClosedDetails ConnectionClosedDetails `json:"connection_closed_details"`
	ConnectionCreated int `json:"connection_created"`
	ConnectionCreatedDetails ConnectionCreatedDetails `json:"connection_created_details"`
	QueueCreated int `json:"queue_created"`
	QueueCreatedDetails QueueCreatedDetails `json:"queue_created_details"`
	QueueDeclared int `json:"queue_declared"`
	QueueDeclaredDetails QueueDeclaredDetails `json:"queue_declared_details"`
	QueueDeleted int `json:"queue_deleted"`
	QueueDeletedDetails QueueDeletedDetails `json:"queue_deleted_details"`
}
type QueueTotals struct {
	Messages int `json:"messages"`
	MessagesDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_details"`
	MessagesReady int `json:"messages_ready"`
	MessagesReadyDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_ready_details"`
	MessagesUnacknowledged int `json:"messages_unacknowledged"`
	MessagesUnacknowledgedDetails struct {
		Rate float64 `json:"rate"`
	} `json:"messages_unacknowledged_details"`
}


type ObjectTotals struct {
	Channels int `json:"channels"`
	Connections int `json:"connections"`
	Consumers int `json:"consumers"`
	Exchanges int `json:"exchanges"`
	Queues int `json:"queues"`
}
//type SocketOpts struct {
//	Backlog int `json:"backlog"`
//	Nodelay bool `json:"nodelay"`
//	Linger []interface{} `json:"linger"`
//	ExitOnClose bool `json:"exit_on_close"`
//}
type CowboyOpts struct {
	Sendfile bool `json:"sendfile"`
}
//type SocketOpts struct {
//	CowboyOpts CowboyOpts `json:"cowboy_opts"`
//	IP string `json:"ip"`
//	Port int `json:"port"`
//}
type Listeners  map[string]interface{}

type Contexts map[string]interface{}
