package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"strconv"
	"sync"
	"net/http"
	"time"
)

type health struct {
	Status string `json:"status"`
}

func(c *Colloect)Checkhealth(){
	for {
		var lock sync.Mutex
		api := "/api/healthchecks/node"
		path := c.url + api
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			level.Info(c.logger).Log("msg", err)
			reason := fmt.Sprintf("%s", err)
			lock.Lock()
			c.lastgetErrorTs.WithLabelValues(strconv.FormatInt(time.Now().Unix(),10),reason,api).Inc()
			lock.Unlock()
		}
		// auth
		req.SetBasicAuth(c.user, c.passwd)
		resp, _ := c.client.Do(req)
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			level.Info(c.logger).Log("err",err)
		}
		var t health
		err = json.Unmarshal(data, &t)
		if err != nil {
			level.Error(c.logger).Log("msg", err)
			reason := fmt.Sprintf("%s", err)
			lock.Lock()
			c.lastgetErrorTs.WithLabelValues(strconv.FormatInt(time.Now().Unix(),10),reason,api).Inc()
			lock.Unlock()
		} else {
			level.Info(c.logger).Log("msg", "check  health success!")
		}
		if t.Status == "ok"{
			c.healthstatus.WithLabelValues(t.Status).Set(1.0)
		}else {
			c.healthstatus.WithLabelValues(t.Status).Set(0.0)
		}
		time.Sleep(c.timeInterval)
	}
}