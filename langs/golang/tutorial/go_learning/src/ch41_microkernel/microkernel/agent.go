package microkernel

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("can not take the operation in the current state")

type CollectorError struct {
	CollectorErrors []error
}

func (ce CollectorError) Error() string {
	var strs []string
	for _, err := range ce.CollectorErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ":")
}

type Event struct {
	Name, Content string
}

type EventReceiver interface {
	OnEvent(evt Event)
}

// 使用context可以取消掉不同的协程里执行的任务
// Init & Destroy可以比较安全的起到安装和卸载collector的功能
type Collector interface {
	Init(evtReceiver EventReceiver) error // 初始化自己要使用的资源
	Start(agtCtx context.Context) error   // 启动
	Stop() error                          // 停止
	Destroy() error                       // 释放掉自己使用的资源
}

type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      int // hold住agent当前所处的状态：启动着、停着、。。，类似状态机的功能，只有在执行正确的操作的时候才执行相应操作
}

// 接收collector发送过来的消息，存在buf中
// 10个消息以后输出一下
func (agt *Agent) EventProcessGoroutine() {
	var evtSeg [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuf:
			case <-agt.ctx.Done():
				return
			}
		}
		fmt.Println(evtSeg)
	}
}

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}
	return &agt
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {

	if agt.state != Waiting {
		return WrongStateError
	}
	agt.collectors[name] = collector
	return collector.Init(agt)
}

func (agt *Agent) startCollectors() error {
	var err error
	var errs CollectorError
	var mutex sync.Mutex
	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			err = collector.Start(ctx)
			mutex.Lock()
			if err != nil {
				errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
			}
		}(name, collector, agt.ctx)
	}
	return errs
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			e := errors.New(name + ":" + err.Error())
			errs.CollectorErrors = append(errs.CollectorErrors, e)
		}
	}
	return errs
}

func (agt *Agent) destroyCollectors() error {

	var err error
	var errs CollectorError
	for name, collector := range agt.collectors {
		if err = collector.Destroy(); err != nil {
			e := errors.New(name + ":" + err.Error())
			errs.CollectorErrors = append(errs.CollectorErrors, e)
		}
	}
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	go agt.EventProcessGoroutine()
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStateError
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent) Destroy() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	return agt.destroyCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}

//<<Kernel>> Agent
//
//>  Extension Point
//
//- <<Plugin>> FileCollector
//- <<Plugin>>ProcessCollector
//- ...
//- <<Plugin>>AppCollector
