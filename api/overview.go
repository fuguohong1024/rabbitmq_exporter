package api

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
)

// 定义desc,调用MustNewConstMetric生产指标并通过channel传递数据
var (
	overviewhealth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "health"),
		"api status(0 for error, 1 for success).",
		[]string{"api"}, nil)
	queuetotalsmessages = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "queue_totals_messages"),
		"Total number of messages (ready plus unacknowledged).",
		[]string{"cluster_name"}, nil)
	queuereadymessages = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "queue_ready_messages"),
		"Number of messages ready for delivery.",
		[]string{"cluster_name"}, nil)
	queueunackmessages = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "queue_unack_messages"),
		"Number of unacknowledged messages.",
		[]string{"cluster_name"}, nil)
	messagestatspublish = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "message_stats_publish"),
		"Messages published recently.",
		[]string{"cluster_name"}, nil)
	messagestatspublishrate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "message_stats_publish_rate"),
		"Message publish rate.",
		[]string{"cluster_name"}, nil)
	messagestatsdeliverget = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "message_stats_deliver_get"),
		"Messages delivered to consumers recently.",
		[]string{"cluster_name"}, nil)
	messagestatsdelivergetrate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "message_stats_deliver_get_rate"),
		"message_stats_deliver_get_rate.",
		[]string{"cluster_name"}, nil)
	connectionsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "connections_total"),
		"Total number of open connections.",
		[]string{"cluster_name"}, nil)
	channelsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "channels_total"),
		"Total number of open channels.",
		[]string{"cluster_name"}, nil)
	queuesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "queues_total"),
		"Total number of queues in use.",
		[]string{"cluster_name"}, nil)
	consumersTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "consumers_total"),
		"Total number of message consumers.",
		[]string{"cluster_name"}, nil)
	exchangesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "overview", "exchanges_total"),
		"Total number of exchanges in use.",
		[]string{"cluster_name"}, nil)
)

var _ Scraper = ScrapeOverview{}

type ScrapeOverview struct{}

func (ScrapeOverview) Name() string {
	return "overview"
}

func (ScrapeOverview) Help() string {
	return "collect overview metrics,true or false"
}

func (ScrapeOverview) Scrape(client *MqClient, ch chan<- prometheus.Metric) error {
	endpoint := "/overview"
	data, err := client.Request(endpoint)
	var t overview
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	// 设定值
	ch <- prometheus.MustNewConstMetric(overviewhealth, prometheus.GaugeValue, 1.0, endpoint)
	ch <- prometheus.MustNewConstMetric(queuetotalsmessages, prometheus.GaugeValue, float64(t.QueueTotals.Messages), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(queuereadymessages, prometheus.GaugeValue, float64(t.QueueTotals.MessagesReady), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(queueunackmessages, prometheus.GaugeValue, float64(t.QueueTotals.MessagesUnacknowledged), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(messagestatspublish, prometheus.GaugeValue, float64(t.MessageStats.Publish), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(messagestatspublishrate, prometheus.GaugeValue, float64(t.MessageStats.PublishDetails.Rate), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(messagestatsdeliverget, prometheus.GaugeValue, float64(t.MessageStats.DeliverGet), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(messagestatsdelivergetrate, prometheus.GaugeValue, float64(t.MessageStats.DeliverGetDetails.Rate), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(connectionsTotal, prometheus.GaugeValue, float64(t.ObjectTotals.Connections), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(channelsTotal, prometheus.GaugeValue, float64(t.ObjectTotals.Channels), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(consumersTotal, prometheus.GaugeValue, float64(t.ObjectTotals.Consumers), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(queuesTotal, prometheus.GaugeValue, float64(t.ObjectTotals.Queues), t.ClusterName)
	ch <- prometheus.MustNewConstMetric(exchangesTotal, prometheus.GaugeValue, float64(t.ObjectTotals.Exchanges), t.ClusterName)

	return nil
}
