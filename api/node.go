package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"strconv"
)

func(c *Colloect)getnodes(){
	for {
		var lock sync.Mutex
		api := "/api/nodes"
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
		var t node
		err = json.Unmarshal(data, &t)
		if err != nil {
			level.Error(c.logger).Log("msg", err,"url",path)
			reason := fmt.Sprintf("%s", err)
			lock.Lock()
			c.lastgetErrorTs.WithLabelValues(strconv.FormatInt(time.Now().Unix(),10),reason,api).Inc()
			lock.Unlock()
		} else {
			level.Info(c.logger).Log("msg", "get nodes status success!")
		}
		for _,v := range t{
			c.nodeuptime.WithLabelValues(v.Name).Set(float64(v.Uptime))
			pid, _ := strconv.Atoi(v.OsPid)
			c.pid.WithLabelValues(v.Name).Set(float64(pid))
			c.memused.WithLabelValues(v.Name).Set(float64(v.MemUsed))
			if v.MemAlarm == true{
				c.memalarm.WithLabelValues(v.Name).Set(1.0)
			}else {
				c.memalarm.WithLabelValues(v.Name).Set(0.0)
			}
			c.memlimit.WithLabelValues(v.Name).Set(float64(v.MemLimit))
			c.diskfreelimit.WithLabelValues(v.Name).Set(float64(v.DiskFreeLimit))
			if v.DiskFreeAlarm == true{
				c.diskfreealarm.WithLabelValues(v.Name).Set(1.0)
			}else {
				c.diskfreealarm.WithLabelValues(v.Name).Set(0.0)
			}
			c.fdtotal.WithLabelValues(v.Name).Set(float64(v.FdTotal))
			c.fdused.WithLabelValues(v.Name).Set(float64(v.FdUsed))
			c.io_file_handle_open_attempt_count.WithLabelValues(v.Name).Set(float64(v.IoFileHandleOpenAttemptCount))
			c.socketstotal.WithLabelValues(v.Name).Set(float64(v.SocketsTotal))
			c.socketsused.WithLabelValues(v.Name).Set(float64(v.SocketsUsed))
			c.ioreadavgtime.WithLabelValues(v.Name).Set(float64(v.IoReadAvgTime))
			c.iowriteavgtime.WithLabelValues(v.Name).Set(float64(v.IoWriteAvgTime))
			c.iosyncavgtime.WithLabelValues(v.Name).Set(float64(v.IoSyncAvgTime))
			c.ioseekavgtime.WithLabelValues(v.Name).Set(float64(v.IoSeekAvgTime))
			c.gcnum.WithLabelValues(v.Name).Set(float64(v.GcNum))
			c.proctotal.WithLabelValues(v.Name).Set(float64(v.ProcTotal))
			c.procused.WithLabelValues(v.Name).Set(float64(v.ProcUsed))
			c.runqueue.WithLabelValues(v.Name).Set(float64(v.RunQueue))
		}
		time.Sleep(c.timeInterval)

	}
}
