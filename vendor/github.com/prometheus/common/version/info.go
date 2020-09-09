// Copyright 2016 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"text/template"
	//"strconv"
	"time"
	"github.com/prometheus/client_golang/prometheus"
)

// Build information. Populated at build-time.
var (
	Version   = "v1"
	Remark   = "单节点mq有些数据没有"
	Branch    = "master"
	BuildUser = "fgh"
	BuildDate = time.Now().Format("2006-01-02")
	GoVersion = runtime.Version()
)

// NewCollector returns a collector that exports metrics about current version
// information.
func NewCollector(program string) prometheus.Collector {
	return prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: program,
			Name:      "build_info",
			Help: fmt.Sprintf(
				"A metric with a constant '1' value labeled by version, revision, branch, and goversion from which %s was built.",
				program,
			),
			ConstLabels: prometheus.Labels{
				"version":   Version,
				"remark":  Remark,
				"branch":    Branch,
				"goversion": GoVersion,
			},
		},
		func() float64 { return 1 },
	)
}

// versionInfoTmpl contains the template used by Info.
// del ", revision: {{.revision}}"
var versionInfoTmpl = `
{{.program}}, version {{.version}} (branch: {{.branch}})
  build user:       {{.buildUser}}
  build date:       {{.buildDate}}
  go version:       {{.goVersion}}
  remark:           {{.remark}}        
`

// Print returns version information.
func Print(program string) string {
	m := map[string]string{
		"program":   program,
		"version":   Version,
		"remark":  Remark,
		"branch":    Branch,
		"buildUser": BuildUser,
		"buildDate": BuildDate,
		"goVersion": GoVersion,
	}
	t := template.Must(template.New("version").Parse(versionInfoTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

// Info returns version, branch and revision information.
func Info() string {
	return fmt.Sprintf("(version=%s, branch=%s, remark=%s)", Version, Branch, Remark)
}

// BuildContext returns goVersion, buildUser and buildDate information.
func BuildContext() string {
	return fmt.Sprintf("(go=%s, user=%s, date=%s)", GoVersion, BuildUser, BuildDate)
}
