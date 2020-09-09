package api

import (
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

type Colloect struct{
	logger log.Logger
	url string
	user string
	passwd string
	client  *http.Client
	timeInterval time.Duration
	lastgetErrorTs *prometheus.CounterVec
	lastUnmarshalErrorts *prometheus.GaugeVec
	//overview metrics
	queuetotalsmessages *prometheus.GaugeVec
	queuereadymessages *prometheus.GaugeVec
	queueunackmessages *prometheus.GaugeVec
	messagestatspublish *prometheus.GaugeVec
	messagestatspublishrate *prometheus.GaugeVec
	messagestatsdeliverget *prometheus.GaugeVec
	messagestatsdelivergetrate *prometheus.GaugeVec
	connectionsTotal *prometheus.GaugeVec
	channelsTotal *prometheus.GaugeVec
	queuesTotal *prometheus.GaugeVec
	consumersTotal *prometheus.GaugeVec
	exchangesTotal *prometheus.GaugeVec
	//health check metrics
	healthstatus *prometheus.GaugeVec
	//nodes metrics
	nodeuptime *prometheus.GaugeVec
	pid *prometheus.GaugeVec
	fdtotal *prometheus.GaugeVec
	fdused *prometheus.GaugeVec
	memused *prometheus.GaugeVec
	memlimit *prometheus.GaugeVec
	memalarm *prometheus.GaugeVec
	diskfreelimit *prometheus.GaugeVec
	diskfreealarm *prometheus.GaugeVec
	io_file_handle_open_attempt_count *prometheus.GaugeVec
	socketstotal *prometheus.GaugeVec
	socketsused *prometheus.GaugeVec
	ioreadavgtime *prometheus.GaugeVec
	iowriteavgtime *prometheus.GaugeVec
	iosyncavgtime *prometheus.GaugeVec
	ioseekavgtime *prometheus.GaugeVec
	gcnum *prometheus.GaugeVec
	gcbytesreclaimed *prometheus.GaugeVec
	proctotal *prometheus.GaugeVec
	procused *prometheus.GaugeVec
	runqueue *prometheus.GaugeVec
	//queue metrics
	memory *prometheus.GaugeVec
	messagesready *prometheus.GaugeVec
	messagesunack *prometheus.GaugeVec
	publish *prometheus.GaugeVec
	deliverget *prometheus.GaugeVec
	publishrate *prometheus.GaugeVec
	delivergetrate *prometheus.GaugeVec

}

//实现Describe和Collect方法
func (c *Colloect)Describe(ch chan<- *prometheus.Desc)  {
	c.lastgetErrorTs.Describe(ch)
	c.lastUnmarshalErrorts.Describe(ch)
	c.queuetotalsmessages.Describe(ch)
	c.queuereadymessages.Describe(ch)
	c.queueunackmessages.Describe(ch)
	c.messagestatspublish.Describe(ch)
	c.messagestatspublishrate.Describe(ch)
	c.messagestatsdeliverget.Describe(ch)
	c.messagestatsdelivergetrate.Describe(ch)
	c.connectionsTotal.Describe(ch)
	c.channelsTotal.Describe(ch)
	c.queuesTotal.Describe(ch)
	c.consumersTotal.Describe(ch)
	c.exchangesTotal.Describe(ch)
	c.healthstatus.Describe(ch)
	c.nodeuptime.Describe(ch)
	c.pid.Describe(ch)
	c.fdtotal.Describe(ch)
	c.fdused.Describe(ch)
	c.memused.Describe(ch)
	c.memlimit.Describe(ch)
	c.memalarm.Describe(ch)
	c.diskfreelimit.Describe(ch)
	c.diskfreealarm.Describe(ch)
	c.io_file_handle_open_attempt_count.Describe(ch)
	c.socketstotal.Describe(ch)
	c.socketsused.Describe(ch)
	c.ioreadavgtime.Describe(ch)
	c.iowriteavgtime.Describe(ch)
	c.iosyncavgtime.Describe(ch)
	c.ioseekavgtime.Describe(ch)
	c.gcnum.Describe(ch)
	c.gcbytesreclaimed.Describe(ch)
	c.proctotal.Describe(ch)
	c.procused.Describe(ch)
	c.runqueue.Describe(ch)
	c.memory.Describe(ch)
	c.messagesready.Describe(ch)
	c.messagesunack.Describe(ch)
	c.publish.Describe(ch)
	c.deliverget.Describe(ch)
	c.publishrate.Describe(ch)
	c.delivergetrate.Describe(ch)
}


func (c *Colloect)Collect(ch chan<- prometheus.Metric)  {
	c.lastgetErrorTs.Collect(ch)
	c.lastUnmarshalErrorts.Collect(ch)
	c.queuetotalsmessages.Collect(ch)
	c.queuereadymessages.Collect(ch)
	c.queueunackmessages.Collect(ch)
	c.messagestatspublish.Collect(ch)
	c.messagestatspublishrate.Collect(ch)
	c.messagestatsdeliverget.Collect(ch)
	c.messagestatsdelivergetrate.Collect(ch)
	c.connectionsTotal.Collect(ch)
	c.channelsTotal.Collect(ch)
	c.queuesTotal.Collect(ch)
	c.consumersTotal.Collect(ch)
	c.exchangesTotal.Collect(ch)
	c.healthstatus.Collect(ch)
	c.nodeuptime.Collect(ch)
	c.pid.Collect(ch)
	c.fdtotal.Collect(ch)
	c.fdused.Collect(ch)
	c.memused.Collect(ch)
	c.memlimit.Collect(ch)
	c.memalarm.Collect(ch)
	c.diskfreelimit.Collect(ch)
	c.diskfreealarm.Collect(ch)
	c.io_file_handle_open_attempt_count.Collect(ch)
	c.socketstotal.Collect(ch)
	c.socketsused.Collect(ch)
	c.ioreadavgtime.Collect(ch)
	c.iowriteavgtime.Collect(ch)
	c.iosyncavgtime.Collect(ch)
	c.ioseekavgtime.Collect(ch)
	c.gcnum.Collect(ch)
	c.gcbytesreclaimed.Collect(ch)
	c.proctotal.Collect(ch)
	c.procused.Collect(ch)
	c.runqueue.Collect(ch)
	c.memory.Collect(ch)
	c.messagesready.Collect(ch)
	c.messagesunack.Collect(ch)
	c.publish.Collect(ch)
	c.deliverget.Collect(ch)
	c.publishrate.Collect(ch)
	c.delivergetrate.Collect(ch)
}

func(c *Colloect)Updatemetrics(){
	go c.getoverview()
	go c.getnodes()
	go c.getqueues()
	go c.Checkhealth()
}





const (
	namespace =  "rabbit_mq"
	subsystem = "queue"
)



//return new Colloect sturct

func Newcolloect(url,user,passwd string,logger log.Logger,client *http.Client,interval time.Duration)*Colloect{
	return &Colloect{
		logger: logger,
		url: url,
		user: user,
		passwd: passwd,
		client:   client,
		timeInterval :  interval ,
		queuetotalsmessages: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "queue_totals_messages",
				Help: "Total number of messages (ready plus unacknowledged)",
			},[]string{"cluster_name"}),

		queuereadymessages: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "queue_ready_messages",
				Help: "Number of messages ready for delivery",
			},[]string{"cluster_name"}),

		queueunackmessages: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "queue_unack_messages",
				Help: "Number of unacknowledged messages",
			},[]string{"cluster_name"}),

		messagestatspublish: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "message_stats_publish",
				Help: "Messages published recently",
			},[]string{"cluster_name"}),

		messagestatspublishrate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "message_stats_publish_rate",
				Help: "Message publish rate",
			},[]string{"cluster_name"}),

		messagestatsdeliverget: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "message_stats_deliver_get",
				Help: "Messages delivered to consumers recently",
			},[]string{"cluster_name"}),

		messagestatsdelivergetrate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "message_stats_deliver_get_rate",
				Help: "Message delivery rate",
			},[]string{"cluster_name"}),

		connectionsTotal : prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "connections_total",
				Help:      "Total number of open connections.",
			},
			[]string{
				// Which node was checked?
				"cluster_name",
			},
		),
		channelsTotal : prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "channels_total",
				Help:      "Total number of open channels.",
			},
			[]string{
				"cluster_name",
			},
		),
		queuesTotal : prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "queues_total",
				Help:      "Total number of queues in use.",
			},
			[]string{
				"cluster_name",
			},
		),
		consumersTotal : prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "consumers_total",
				Help:      "Total number of message consumers.",
			},
			[]string{
				"cluster_name",
			},
		),
		exchangesTotal : prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "exchanges_total",
				Help:      "Total number of exchanges in use.",
			},
			[]string{
				"cluster_name",
			},
		),

		healthstatus: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "health_status",
				Help: "health checks",
			},
			[]string{"status"},
		),

		nodeuptime: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "node_uptime",
				Help: "node uptime",
			},
			[]string{"node"},
		),

		lastgetErrorTs: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name: "last_get_errts",
				Help: "last get data err timeStamp",
			},
			[]string{"ts","err","api"},
		),

		pid: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "process_id",
				Help: "os process id",
			},
			[]string{"node"},
		),

		fdtotal: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "fd_total",
				Help: "File descriptors available",
			},
			[]string{"node"},
		),

		fdused: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "fd_used",
				Help: "File descriptors used",
			},
			[]string{"node"},
		),

		memused: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "mem_used",
				Help: "Total amount of memory used",
			},
			[]string{"node"},
		),

		memlimit: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "mem_limit",
				Help: "Memory usage high watermark",
			},
			[]string{"node"},
		),

		memalarm: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "mem_alarm",
				Help: "memory alarm in effect?true or false",
			},
			[]string{"node"},
		),

		diskfreelimit: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "disk_free_limit",
				Help: "Free disk space low watermark",
			},
			[]string{"node"},
		),

		diskfreealarm: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "disk_free_alarm",
				Help: "disk alarm in effect?true or false",
			},
			[]string{"node"},
		),

		io_file_handle_open_attempt_count: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "io_file_handle_open_attempt_count",
				Help: "File descriptor open attempts",
			},
			[]string{"node"},
		),

		socketstotal:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "sockets_total",
				Help: "Sockets available",
			},
			[]string{"node"},
		),

		socketsused:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "sockets_used",
				Help: "Sockets used",
			},
			[]string{"node"},
		),

		ioreadavgtime:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "io_read_avg_time",
				Help: "io read avg time",
			},
			[]string{"node"},
		),

		iowriteavgtime:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "io_write_avg_time",
				Help: "io write avg time",
			},
			[]string{"node"},
		),

		ioseekavgtime: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "io_seek_avg_time",
				Help: "io seek avg time",
			},[]string{"node"},
		),

		iosyncavgtime: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "io_sync_avg_time",
				Help: "io sync avg time",
			},[]string{"node"},
		),

		gcnum:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "gc_num",
				Help: "GC runs num",
			},
			[]string{"node"},
		),

		gcbytesreclaimed:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "gc_bytes_reclaimed",
				Help: "Bytes reclaimed by GC",
			},
			[]string{"node"},
		),

		proctotal:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "proc_total",
				Help: "Erlang process limit",
			},
			[]string{"node"},
		),

		procused:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "proc_used",
				Help: "Erlang processes used",
			},
			[]string{"node"},
		),

		runqueue:	prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "run_queue",
				Help: "Runtime run queue",
			},
			[]string{"node"},
		),

		memory: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name: "memory",
				Help: "Memory",
			},[]string{"node","queue_name"}),

		messagesready: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name: "messages_ready",
				Help: "Number of messages ready for delivery",
			},[]string{"node","queue_name"}),

		messagesunack: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name: "messages_unacknowledged",
				Help: "Number of unacknowledged messages",
			},[]string{"node","queue_name"}),

		publish: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name: "message_stats_publish",
				Help: "Messages published recently",
			},[]string{"node","queue_name"}),

		publishrate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name: "message_stats_publish_rate",
				Help: "",
			}, []string{"node","queue_name"}),

		deliverget: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name: "message_stats_deliver_get",
				Help: "Messages delivered recently",
			},[]string{"node","queue_name"}),

		delivergetrate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name: "message_stats_deliver_get_rate",
				Help: "Message delivery rate",
			}, []string{"node","queue_name"}),

		lastUnmarshalErrorts: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name: "last_Unmarshal_Error_ts",
				Help: "last Unmarshal json error ts",
			},[]string{"ts","err","api"}),
	}
}


