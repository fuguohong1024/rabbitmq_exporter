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

//GET /api/overview
func (c *Colloect) getoverview() {
	for {
		var lock sync.Mutex
		api := "/api/overview"
		path := c.url + api
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			level.Error(c.logger).Log("msg", err)
			reason := fmt.Sprintf("%s", err)
			lock.Lock()
			c.lastgetErrorTs.WithLabelValues(strconv.FormatInt(time.Now().Unix(), 10), reason, api).Inc()
			lock.Unlock()
		}
		// auth
		req.SetBasicAuth(c.user, c.passwd)
		resp, _ := c.client.Do(req)
		if false == httpnotok(resp, c.logger) {
			break
		} else {
			data, err := ioutil.ReadAll(resp.Body)
			var t overview
			err = json.Unmarshal(data, &t)
			if err != nil {
				level.Error(c.logger).Log("msg", err, "url", path)
				reason := fmt.Sprintf("%s", err)
				lock.Lock()
				c.lastgetErrorTs.WithLabelValues(strconv.FormatInt(time.Now().Unix(), 10), reason, api).Inc()
				lock.Unlock()
			} else {
				level.Info(c.logger).Log("msg", "get overview metrics success!")
			}
			c.connectionsTotal.WithLabelValues(t.ClusterName).Set(float64(t.ObjectTotals.Connections))
			c.channelsTotal.WithLabelValues(t.ClusterName).Set(float64(t.ObjectTotals.Channels))
			c.queuesTotal.WithLabelValues(t.ClusterName).Set(float64(t.ObjectTotals.Queues))
			c.consumersTotal.WithLabelValues(t.ClusterName).Set(float64(t.ObjectTotals.Consumers))
			c.exchangesTotal.WithLabelValues(t.ClusterName).Set(float64(t.ObjectTotals.Exchanges))
			c.queuetotalsmessages.WithLabelValues(t.ClusterName).Set(float64(t.QueueTotals.Messages))
			c.queuereadymessages.WithLabelValues(t.ClusterName).Set(float64(t.QueueTotals.MessagesReady))
			c.queueunackmessages.WithLabelValues(t.ClusterName).Set(float64(t.QueueTotals.MessagesUnacknowledged))
			c.messagestatspublish.WithLabelValues(t.ClusterName).Set(float64(t.MessageStats.Publish))
			c.messagestatspublishrate.WithLabelValues(t.ClusterName).Set(float64(t.MessageStats.PublishDetails.Rate))
			c.messagestatsdeliverget.WithLabelValues(t.ClusterName).Set(float64(t.MessageStats.DeliverGet))
			c.messagestatsdelivergetrate.WithLabelValues(t.ClusterName).Set(float64(t.MessageStats.DeliverGetDetails.Rate))

			time.Sleep(c.timeInterval)
		}
	}
}
