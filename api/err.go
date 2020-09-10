package api

import (
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
	"os"
)

type httperr struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

// http.statuscode!=200
func httpnotok(resp *http.Response, logger log.Logger) bool {
	if resp.StatusCode != http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			level.Error(logger).Log("msg", err)
		}
		var t httperr
		err = json.Unmarshal(data, &t)
		level.Error(logger).Log("err", t.Error, "msg", t.Reason)
		return false
	}
	return true
}

func Check(user, passwd string, logger log.Logger) {
	if user == "" {
		level.Error(logger).Log("err", "mq user is null")
		os.Exit(1)
	} else if passwd == "" {
		level.Error(logger).Log("err", "mq passwd is null")
		os.Exit(1)
	}
}
