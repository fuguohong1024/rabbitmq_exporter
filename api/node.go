package api

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
)

// 定义desc,调用MustNewConstMetric生产指标并通过channel传递数据
var (
	nodehealth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "health"),
		"api status(0 for error, 1 for success).",
		[]string{"api"}, nil)
	nodeuptime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "up_time"),
		"node uptime.",
		[]string{"node"}, nil)
	fdtotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "fd_total"),
		"File descriptors available.",
		[]string{"node"}, nil)
	fdused = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "fd_used"),
		"File descriptors used.",
		[]string{"node"}, nil)
	memused = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "mem_used"),
		"Total amount of memory used.",
		[]string{"node"}, nil)
	memlimit = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "memlimit"),
		"Memory usage high watermark.",
		[]string{"node"}, nil)
	socketstotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "sockets_total"),
		"Sockets available.",
		[]string{"node"}, nil)
	socketsused = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "sockets_used"),
		"Sockets used.",
		[]string{"node"}, nil)
	ioreadavgtime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "io_read_avg_time"),
		"io read avg time.",
		[]string{"node"}, nil)
	iowriteavgtime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "io_write_avg_time"),
		"io write avg time.",
		[]string{"node"}, nil)
	ioseekavgtime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "io_seek_avg_time"),
		"io seek avg time.",
		[]string{"node"}, nil)
	iosyncavgtime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "io_sync_avg_time"),
		"io sync avg time.",
		[]string{"node"}, nil)
	gcnum = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "gc_num"),
		"GC runs num.",
		[]string{"node"}, nil)
	gcbytesreclaimed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "gc_bytes_reclaimed"),
		"Bytes reclaimed by GC.",
		[]string{"node"}, nil)
	proctotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "proc_total"),
		"Erlang process limit.",
		[]string{"node"}, nil)
	procused = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "proc_used"),
		"Erlang processes used.",
		[]string{"node"}, nil)
	runqueue = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "run_queue"),
		"Runtime run queue.",
		[]string{"node"}, nil)
)

// var name  type value,判断结构体是否满足接口
var _ Scraper = ScrapeNode{}

type ScrapeNode struct{}

func (ScrapeNode) Name() string {
	return "node"
}

func (ScrapeNode) Help() string {
	return "collect node metrics,true or false"
}
func (ScrapeNode) Scrape(client *MqClient, ch chan<- prometheus.Metric) error {
	endpoint := "/nodes"
	data, err := client.Request(endpoint)
	var t node
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	// 设定值
	ch <- prometheus.MustNewConstMetric(nodehealth, prometheus.GaugeValue, 1.0, endpoint)
	for _, v := range t {
		// MustNewConstMetric(desc *Desc, valueType ValueType, value float64, labelValues ...string) Metric
		ch <- prometheus.MustNewConstMetric(nodeuptime, prometheus.GaugeValue, float64(v.Uptime), v.Name)
		ch <- prometheus.MustNewConstMetric(fdtotal, prometheus.GaugeValue, float64(v.FdTotal), v.Name)
		ch <- prometheus.MustNewConstMetric(fdused, prometheus.GaugeValue, float64(v.FdUsed), v.Name)
		ch <- prometheus.MustNewConstMetric(memlimit, prometheus.GaugeValue, float64(v.MemLimit), v.Name)
		ch <- prometheus.MustNewConstMetric(socketstotal, prometheus.GaugeValue, float64(v.SocketsTotal), v.Name)
		ch <- prometheus.MustNewConstMetric(socketsused, prometheus.GaugeValue, float64(v.SocketsUsed), v.Name)
		ch <- prometheus.MustNewConstMetric(ioreadavgtime, prometheus.GaugeValue, float64(v.IoReadAvgTime), v.Name)
		ch <- prometheus.MustNewConstMetric(iowriteavgtime, prometheus.GaugeValue, float64(v.IoWriteAvgTime), v.Name)
		ch <- prometheus.MustNewConstMetric(ioseekavgtime, prometheus.GaugeValue, float64(v.IoSeekAvgTime), v.Name)
		ch <- prometheus.MustNewConstMetric(iosyncavgtime, prometheus.GaugeValue, float64(v.IoSyncAvgTime), v.Name)
		ch <- prometheus.MustNewConstMetric(gcnum, prometheus.GaugeValue, float64(v.GcNum), v.Name)
		ch <- prometheus.MustNewConstMetric(gcbytesreclaimed, prometheus.GaugeValue, float64(v.GcBytesReclaimed), v.Name)
		ch <- prometheus.MustNewConstMetric(proctotal, prometheus.GaugeValue, float64(v.ProcTotal), v.Name)
		ch <- prometheus.MustNewConstMetric(procused, prometheus.GaugeValue, float64(v.ProcUsed), v.Name)
		ch <- prometheus.MustNewConstMetric(runqueue, prometheus.GaugeValue, float64(v.RunQueue), v.Name)
	}

	return nil
}
