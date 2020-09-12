package api

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"net/http"
	"time"
)

// 是否抓取
var (
	Scrapers = map[Scraper]bool{
		ScrapeNode{}:     true,
		ScrapeOverview{}: true,
		ScrapeQueue{}:    true,
	}
)

// 定义mq实体
type MqOpts struct {
	Url      string
	Username string
	Password string
	Timeout  time.Duration
}

// 定义mq连接结构体
// 定义client属性及日志属性
// MqClient方法AddFlag,request和ping
type MqClient struct {
	Client *http.Client
	Opts   *MqOpts
}

// 赋值MqOpts
func (m *MqOpts) AddFlag() {
	flag.StringVar(&m.Url, "mq.url", "http://testmojing.ts/api", "HTTP API address of a mq server or agent. (prefix with https:// to connect over HTTPS)")
	flag.StringVar(&m.Username, "mq.user", "user", "mq username")
	flag.StringVar(&m.Password, "mq.pwd", "2Jv4v3Qjrx", "mq password")
	//flag.StringVar(&m.UA, "harbor-ua", "harbor_exporter", "user agent of the harbor http client")
	flag.DurationVar(&m.Timeout, "timeout", time.Second*5, "Timeout on HTTP requests to the harbor API.")
	//flag.BoolVar(&m.Insecure, "insecure", true, "Disable TLS host verification.")
}

// 根据不同的url请求,返回[]byte
func (m *MqClient) Request(endpoint string) ([]byte, error) {
	url := m.Opts.Url + endpoint
	// GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("http GET err:%v", err)
		// 退出
		return nil, err
	}

	req.SetBasicAuth(m.Opts.Username, m.Opts.Password)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := m.Client.Do(req)
	if err != nil {
		log.Errorf("GET Response err:%v", err)
		return nil, err
	}
	// 延迟关闭连接
	defer resp.Body.Close()
	// httpstatuscode判断
	if resp.StatusCode != http.StatusOK {
		reason := fmt.Sprintf("error handling request for %s http-statuscode: %s", endpoint, resp.Status)
		log.Errorf("httpStatusCode != 200,err %v", reason)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("IO READ err:%v", err)
		return nil, err
	}

	return body, nil
}

// 验证账号密码及url可用
func (m *MqClient) Ping() (bool, error) {
	req, err := http.NewRequest("GET", m.Opts.Url+"/whoami", nil)
	if err != nil {
		log.Errorf("GET whoami api err:%v", err)
		return false, err
	}
	req.SetBasicAuth(m.Opts.Username, m.Opts.Password)

	resp, err := m.Client.Do(req)
	if err != nil {
		return false, err
	}

	resp.Body.Close()

	switch {
	case resp.StatusCode == http.StatusOK:
		return true, nil
	case resp.StatusCode == http.StatusUnauthorized:
		return false, fmt.Errorf("username or password incorrect, http-statuscode: %s", resp.Status)
	default:
		return false, fmt.Errorf("error handling request, http-statuscode: %s", resp.Status)
	}
}
