package api

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	queuehealth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "queue", "health"),
		"api status(0 for error, 1 for success).",
		[]string{"api"}, nil)
	memory = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "queue", "memory"),
		"Memory.",
		[]string{"node", "queue_name"}, nil)
	messagesready = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "queue", "messages_ready"),
		"Number of messages ready for delivery.",
		[]string{"node", "queue_name"}, nil)
	messagesunack = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "queue", "messages_unack"),
		"Number of unacknowledged messages.",
		[]string{"node", "queue_name"}, nil)
	publish = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "queue", "message_stats_publish"),
		"Messages published recently.",
		[]string{"node", "queue_name"}, nil)
	deliverget = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "queue", "message_stats_publish_rate"),
		"Messages delivered recently.",
		[]string{"node", "queue_name"}, nil)
	publishrate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "queue", "message_stats_deliver_get"),
		"Message published rate.",
		[]string{"node", "queue_name"}, nil)
	delivergetrate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "queue", "message_stats_deliver_get_rate"),
		"Message delivery rate.",
		[]string{"node", "queue_name"}, nil)
)

// var name  type value,判断结构体是否满足接口
var _ Scraper = ScrapeQueue{}

type ScrapeQueue struct{}

func (ScrapeQueue) Name() string {
	return "queue"
}

func (ScrapeQueue) Help() string {
	return "collect queue metrics,true or false"
}
func (ScrapeQueue) Scrape(client *MqClient, ch chan<- prometheus.Metric) error {
	endpoint := "/queues"
	data, err := client.Request(endpoint)
	var t queue
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	ch <- prometheus.MustNewConstMetric(queuehealth, prometheus.GaugeValue, 1.0, endpoint)
	for _, v := range t {
		ch <- prometheus.MustNewConstMetric(memory, prometheus.GaugeValue, float64(v.Memory), v.Node, v.Name)
		ch <- prometheus.MustNewConstMetric(messagesready, prometheus.GaugeValue, float64(v.MessagesReady), v.Node, v.Name)
		ch <- prometheus.MustNewConstMetric(messagesunack, prometheus.GaugeValue, float64(v.MessagesUnacknowledged), v.Node, v.Name)
		ch <- prometheus.MustNewConstMetric(publish, prometheus.GaugeValue, float64(v.MessageStats.Publish), v.Node, v.Name)
		ch <- prometheus.MustNewConstMetric(deliverget, prometheus.GaugeValue, float64(v.MessageStats.PublishDetails.Rate), v.Node, v.Name)
		ch <- prometheus.MustNewConstMetric(publishrate, prometheus.GaugeValue, float64(v.MessageStats.DeliverGet), v.Node, v.Name)
		ch <- prometheus.MustNewConstMetric(delivergetrate, prometheus.GaugeValue, float64(v.MessageStats.DeliverDetails.Rate), v.Node, v.Name)
	}

	return nil
}
