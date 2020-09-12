package api

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

// 定义常量
const (
	name      = "RabbitMq_exporter"
	namespace = "rabbitmq"
	subsystem = "exporter"
)

func Name() string {
	return name
}

// 定义prometheus.Desc
var (
	scrapeDurationDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "collector_duration_seconds"),
		"Collector time duration.",
		[]string{"collector"}, nil,
	)
)

// 定义基础metrics
type Metrics struct {
	TotalScrapes prometheus.Counter
	ScrapeErrors *prometheus.CounterVec
	Error        prometheus.Gauge
	MqUp         prometheus.Gauge
}

// New返回Metrics
func NewMetrics() Metrics {
	return Metrics{
		TotalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "scrapes_total",
			Help:      "Total number of times harbor was scraped for metrics.",
		}),
		ScrapeErrors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "scrape_errors_total",
			Help:      "Total number of times an error occurred scraping a harbor.",
		}, []string{"collector"}),
		Error: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "last_scrape_error",
			Help:      "Whether the last scrape of metrics from harbor resulted in an error (1 for error, 0 for success).",
		}),
		MqUp: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Whether the RabbitMq is up.",
		}),
	}
}

// 定义Exporter
type Exporter struct {
	//ctx      context.Context  //http timeout will work, don't need this
	client   *MqClient
	scrapers []Scraper
	metrics  Metrics
}

// New返回Exporter
func NewExporter(Mq *MqOpts, metrics Metrics, scrapers []Scraper) (*Exporter, error) {
	uri := Mq.Url
	// url如果不包含http://或者https:// 不全
	if !strings.Contains(uri, "://") {
		uri = "http://" + uri
	}
	// 解析url是否合格
	u, err := url.Parse(uri)
	if err != nil {
		// 输出错误类型及详情
		log.Errorf("invalid Mq URL:%v", err)
	}
	if u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
		log.Errorf("invalid Mq URL:%v", err)
	}

	// 环境变量设置账号密码
	user := os.Getenv("MQ_USER")
	if user != "" {
		Mq.Username = user
	}

	pass := os.Getenv("MQ_PWD")
	if pass != "" {
		Mq.Password = pass
	}

	mq := &MqClient{
		Opts:   Mq,
		Client: &http.Client{Timeout: Mq.Timeout},
	}

	return &Exporter{
		client:   mq,
		metrics:  metrics,
		scrapers: scrapers,
	}, nil
}

// 实现colletor接口
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.metrics.TotalScrapes.Desc()
	ch <- e.metrics.Error.Desc()
	e.metrics.ScrapeErrors.Describe(ch)
	ch <- e.metrics.MqUp.Desc()
}

// 实现colletor接口
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// 使用自定义抓取数据接口,抓取中有Describe
	e.scrape(ch)

	ch <- e.metrics.TotalScrapes
	ch <- e.metrics.Error
	e.metrics.ScrapeErrors.Collect(ch)
	ch <- e.metrics.MqUp
}

// 抓取数据,在Collect中调用
func (e *Exporter) scrape(ch chan<- prometheus.Metric) {
	// 总抓取次数
	e.metrics.TotalScrapes.Inc()
	// 抓取时间
	scrapeTime := time.Now()
	// 验证账号密码及url可用
	if pong, err := e.client.Ping(); pong != true || err != nil {
		log.Errorf("ping api err:%v", err)
		e.metrics.MqUp.Set(0)
		e.metrics.Error.Set(1)
	}
	e.metrics.MqUp.Set(1)
	e.metrics.Error.Set(0)
	// 抓取时间
	ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(scrapeTime).Seconds(), "ping")
	// 等待抓取协程结束
	var wg sync.WaitGroup
	defer wg.Wait()
	for _, scraper := range e.scrapers {
		// 协程间通信安全
		wg.Add(1)
		go func(scraper Scraper) {
			defer wg.Done()
			label := scraper.Name()
			scrapeTime := time.Now()
			if err := scraper.Scrape(e.client, ch); err != nil {
				reason := fmt.Sprintf("Scrape %v err,%v", scraper.Name(), err)
				log.Error(reason)
				// 记录抓取失败的Scraper
				e.metrics.ScrapeErrors.WithLabelValues(label).Inc()
				e.metrics.Error.Set(1)
			}
			ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(scrapeTime).Seconds(), label)
		}(scraper)
	}
}
