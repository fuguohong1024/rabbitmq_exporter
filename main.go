package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"net/http"
	"os"
	"rabbitmq_exporter/api"
	"runtime"
	"time"
)

func main() {

	listenAddress := flag.String("web.listen.address", ":49107", "Address to listen on for web interface and telemetry.")
	metricsPath := flag.String("web.telemetry.path", "/metrics", "Path under which to expose metrics.")
	logLevel := flag.String("log-level", "info", "The logging level:[debug, info, warn, error, fatal]")
	logFile := flag.String("log-output", "", "the file which log to, default stdout")
	v := flag.Bool("version", false, "print version info!")

	opts := &api.MqOpts{}
	// 复制opts
	opts.AddFlag()

	// 根据默认设置或flag生成抓取清单
	scraperFlags := map[api.Scraper]*bool{}
	for scraper, enabledByDefault := range api.Scrapers {
		defaultOn := false
		if enabledByDefault {
			defaultOn = true
		}
		f := flag.Bool("collect."+scraper.Name(), defaultOn, scraper.Help())
		scraperFlags[scraper] = f
	}
	flag.Parse()

	if *v {
		fmt.Print(versionPrint())
		return
	}

	//抓取清单
	enabledScrapers := []api.Scraper{}
	for scraper, enabled := range scraperFlags {
		if *enabled {
			log.Infof("Scraper enabled %s", scraper.Name())
			enabledScrapers = append(enabledScrapers, scraper)
		}
	}

	if err := LogInit(*logLevel, *logFile); err != nil {
		log.Fatal(errors.Wrap(err, "set log level error"))
	}

	// NewExporter
	exporter, err := api.NewExporter(opts, api.NewMetrics(), enabledScrapers)
	if err != nil {
		log.Infof("New Exporter err:%s", err)
	}

	// 注册
	prometheus.MustRegister(exporter)

	// 设置路由函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Rabbitmq_exporter</title></head>
             <body>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	http.Handle(*metricsPath, promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer,
		promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{
				ErrorLog: log.StandardLogger(),
			},
		),
	),
	)

	// 检测服务状态路由
	http.HandleFunc("/-/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})
	log.Infof("Listening on address %s", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Infof("http listen err:%s", err)
	}

}

var (
	builduser = "fgh"
	Version   = "v2"
	branch    = "master"                        // state of git tree, either "clean" or "dirty"
	buildDate = time.Now().Format("2006-01-02") // build date, output of $(date +'%Y-%m-%dT%H:%M:%S')
)

func versionPrint() string {
	return fmt.Sprintf(`Name: %s
Version: %s
build user: %s
build date: %s
branch: %s
go version: %s
Compiler: %s
Platform: %s/%s
`, api.Name(), Version, builduser, buildDate, branch, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
}

func LogInit(level, file string) error {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	le, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetLevel(le)

	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		log.SetOutput(f)
	}

	return nil
}
