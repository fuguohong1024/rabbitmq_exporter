// Copyright 2014 The Prometheus Authors
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

package prometheus

// Collector is the interface implemented by anything that can be used by
// Prometheus to collect metrics. A Collector has to be registered for
// collection. See Registerer.Register.
//
// The stock metrics provided by this package (Gauge, Counter, Summary,
// Histogram, Untyped) are also Collectors (which only ever collect one metric,
// namely itself). An implementer of Collector may, however, collect multiple
// metrics in a coordinated fashion and/or create metrics on the fly. Examples
// for collectors already implemented in this library are the metric vectors
// (i.e. collection of multiple instances of the same Metric but with different
// label values) like GaugeVec or SummaryVec, and the ExpvarCollector.
type Collector interface {
	// Describe将此收集器收集的指标的所有可能描述符的超集发送到提供的通道，并在发送完最后一个描述符后返回。
	//发送的描述符满足Desc文档中描述的一致性和唯一性要求。
	// 如果一个收集器和同一收集器发送重复的描述符，则该方法有效。 这些重复项将被忽略。
	//但是，两个不同的收集器一定不能发送重复的描述符。
	// 完全不发送任何描述符会将收集器标记为“未检查”，即在注册时将不执行检查，
	//并且收集器可以在其Collect方法中产生它认为合适的任何度量标准。
	//此方法在收集器的整个生命周期中均等地发送相同的描述符。 它可以被同时调用，因此必须以并发安全的方式实现。
	//如果收集器在执行此方法时遇到错误，则它必须发送一个无效的描述符（使用NewInvalidDesc创建），以将错误信号通知注册表。
	Describe(chan<- *Desc)
	//收集指标时，Prometheus注册表会调用“收集”。 该实现通过提供的通道发送每个收集的度量，并在发送完最后一个度量后返回。
	//每个发送的指标的描述符是Describe返回的指标之一（除非未选中收集器，请参见上文）。
	//共享相同描述符的返回指标必须在其可变标签值上有所不同。
	//可以同时调用此方法，因此必须以并发安全的方式实现。
	//发生阻塞会以呈现所有已注册指标的总体性能为代价。 理想情况下，收集器实现支持并发读取器。
	Collect(chan<- Metric)
}

// DescribeByCollect is a helper to implement the Describe method of a custom
// Collector. It collects the metrics from the provided Collector and sends
// their descriptors to the provided channel.
//
// If a Collector collects the same metrics throughout its lifetime, its
// Describe method can simply be implemented as:
//
//   func (c customCollector) Describe(ch chan<- *Desc) {
//   	DescribeByCollect(c, ch)
//   }
//
// However, this will not work if the metrics collected change dynamically over
// the lifetime of the Collector in a way that their combined set of descriptors
// changes as well. The shortcut implementation will then violate the contract
// of the Describe method. If a Collector sometimes collects no metrics at all
// (for example vectors like CounterVec, GaugeVec, etc., which only collect
// metrics after a metric with a fully specified label set has been accessed),
// it might even get registered as an unchecked Collector (cf. the Register
// method of the Registerer interface). Hence, only use this shortcut
// implementation of Describe if you are certain to fulfill the contract.
//
// The Collector example demonstrates a use of DescribeByCollect.
func DescribeByCollect(c Collector, descs chan<- *Desc) {
	metrics := make(chan Metric)
	go func() {
		c.Collect(metrics)
		close(metrics)
	}()
	for m := range metrics {
		descs <- m.Desc()
	}
}

// selfCollector implements Collector for a single Metric so that the Metric
// collects itself. Add it as an anonymous field to a struct that implements
// Metric, and call init with the Metric itself as an argument.
type selfCollector struct {
	self Metric
}

// init provides the selfCollector with a reference to the metric it is supposed
// to collect. It is usually called within the factory function to create a
// metric. See example.
func (c *selfCollector) init(self Metric) {
	c.self = self
}

// Describe implements Collector.
func (c *selfCollector) Describe(ch chan<- *Desc) {
	ch <- c.self.Desc()
}

// Collect implements Collector.
func (c *selfCollector) Collect(ch chan<- Metric) {
	ch <- c.self
}
