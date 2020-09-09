package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func(c *Colloect)getqueues() {
	for {
		var lock sync.Mutex
		api := "/api/queues"
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
		var t queue
		err = json.Unmarshal(data, &t)
		if err != nil {
			level.Error(c.logger).Log("msg", err,"url",path)
			reason := fmt.Sprintf("%s", err)
			lock.Lock()
			c.lastgetErrorTs.WithLabelValues(strconv.FormatInt(time.Now().Unix(),10),reason,api).Inc()
			lock.Unlock()
		} else {
			level.Info(c.logger).Log("msg", "get queues status success!")
		}
		for _,v := range t{
			c.memory.WithLabelValues(v.Node,v.Name).Set(float64(v.Memory))
			c.messagesready.WithLabelValues(v.Node,v.Name).Set(float64(v.MessagesReady))
			c.messagesunack.WithLabelValues(v.Node,v.Name).Set(float64(v.MessagesUnacknowledged))
			c.publish.WithLabelValues(v.Node,v.Name).Set(float64(v.MessageStats.Publish))
			c.publishrate.WithLabelValues(v.Node,v.Name).Set(float64(v.MessageStats.PublishDetails.Rate))
			c.deliverget.WithLabelValues(v.Node,v.Name).Set(float64(v.MessageStats.DeliverGet))
			c.delivergetrate.WithLabelValues(v.Node,v.Name).Set(float64(v.MessageStats.DeliverDetails.Rate))
		}

		time.Sleep(c.timeInterval)
	}
}