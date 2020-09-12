package api

import (
	"github.com/prometheus/client_golang/prometheus"
)

// 每个抓取接口实现自己的方法来抓取数据

// 自定义抓取接口,方法由结构体具体定义
type Scraper interface {
	// Name of the Scraper. Should be unique.
	Name() string

	// Help describes the role of the Scraper.
	// Example: "Collect from SHOW ENGINE INNODB STATUS"
	Help() string

	// Scrape收集数据，并将prometheus.Metric通过channel传输
	Scrape(client *MqClient, ch chan<- prometheus.Metric) error
}
