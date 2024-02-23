package simulation

import (
	"fmt"
	"github.com/clouddea/devs-go/modeling"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Root struct {
	processor Processor
	t         uint64
	inited    bool
}

type RootTimeArg struct {
	T uint64
}

type RootMessageArg struct {
	Message modeling.Message
	T       uint64
}

func NewRoot(processor Processor) *Root {
	return &Root{
		processor: processor,
		t:         0,
		inited:    false,
	}
}

func (this *Root) Init(input *RootTimeArg, output *RootTimeArg) error {
	this.processor.Init(input.T)
	return nil
}

func (this *Root) Advance(input *RootTimeArg, output *RootTimeArg) error {
	this.processor.Advance(input.T)
	return nil
}

func (this *Root) ComputeOutput(input *RootTimeArg, ouput *RootTimeArg) error {
	this.processor.ComputeOutput(input.T)
	return nil
}

func (this *Root) PutMessage(input *RootMessageArg, output *RootTimeArg) error {
	this.processor.PutMessage(input.Message, input.T)
	return nil
}

func (this *Root) GetTN(input *RootTimeArg, output *RootTimeArg) error {
	output.T = this.processor.GetTN()
	return nil
}

func (receiver *Root) Setup() {
	receiver.processor.Init(0)
}

func (receiver *Root) Step() uint64 {
	if receiver.t < modeling.INFINITE {
		tn := receiver.processor.GetTN()
		if tn >= modeling.INFINITE {
			return tn
		}
		receiver.processor.ComputeOutput(tn)
		receiver.processor.Advance(tn)
		receiver.t = tn
	}
	return receiver.t
}

/** 启动仿真 */
func (receiver *Root) Simulate(delay time.Duration, stepsCallback func(t uint64)) {
	receiver.processor.Init(0)
	for receiver.t < modeling.INFINITE {
		tn := receiver.processor.GetTN()
		if tn >= modeling.INFINITE {
			break
		}
		receiver.processor.ComputeOutput(tn)
		receiver.processor.Advance(tn)
		receiver.t = tn
		if stepsCallback != nil {
			stepsCallback(receiver.t)
		}
		time.Sleep(delay)
	}
}

/** 启动服务 */
func (receiver *Root) Serve(endpoint string) {
	err := rpc.Register(receiver)
	if err != nil {
		log.Fatal("register service error:", err)
	}
	rpc.HandleHTTP()
	lister, e := net.Listen("tcp", endpoint) // :1234
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("served on " + endpoint + " ...")
	err = http.Serve(lister, nil)
	if err != nil {
		log.Fatal("service serve error:", err)
	}
}
