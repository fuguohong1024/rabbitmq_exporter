package main

import (
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
	"rabbitmq_exporter/api"
	"rabbitmq_exporter/etc"
)

func main() {
	var (
		Name = "rabbit_mq_exporter"

		listenAddress = kingpin.Flag("web.listen.address", "web listen address").
				Default(":41690").Envar("Listen_ADD").String()

		mqurl = kingpin.Flag("mq.url",
			"the url of rabbit_mq").
			Default("http://127.0.0.1:15672").Envar("MQ_URL").String()

		mquser = kingpin.Flag("mq.user", "the user of rabbit_mq").
			Default("").String()

		mqpasswd = kingpin.Flag("mq.passwd", "the password of rabbit_mq").
				Default("").String()

		timeInterval = kingpin.Flag("collect.interval",
			"collect mq info  interval ").
			Default("30s").Envar("COLLECT_INTERVAL").Duration()

		logLevel = kingpin.Flag("log.level",
			"Sets the loglevel. Valid levels are debug, info, warn, error").
			Default("info").Envar("LOG_LEVEL").String()

		logFormat = kingpin.Flag("log.format",
			"Sets the log format. Valid formats are json and logfmt").
			Default("logfmt").Envar("LOG_FMT").String()

		logOutput = kingpin.Flag("log.output",
			"Sets the log output. Valid outputs are stdout and stderr").
			Default("stdout").Envar("LOG_OUTPUT").String()
	)

	kingpin.Version(version.Print(Name))
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	logger := etc.GetLogger(*logLevel, *logOutput, *logFormat)

	api.Check(*mquser, *mqpasswd, logger)

	httpClient := &http.Client{Timeout: *timeInterval}

	collect := api.Newcolloect(*mqurl, *mquser, *mqpasswd, logger, httpClient, *timeInterval)

	prometheus.MustRegister(collect)
	collect.Updatemetrics()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		_ = level.Error(logger).Log(
			"msg", "http server quit",
			"err", err,
		)
		os.Exit(1)

	} else if err == nil {
		_ = level.Info(logger).Log("starting  rabbit_mq_exporter at port", *listenAddress)
	}

}
