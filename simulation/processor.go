package simulation

import (
	"github.com/clouddea/devs-go/modeling"
	"log"
	"net/rpc"
	"sync"
)

type Processor interface {
	Init(t uint64)
	Advance(t uint64)
	ComputeOutput(t uint64)
	PutMessage(message modeling.Message, t uint64)
	GetTN() uint64
}

type ProcessorStub struct {
	client   *rpc.Client
	endpoint string
	lock     sync.Mutex
}

func NewProcessorStub(endpoint string) *ProcessorStub {
	return &ProcessorStub{
		client:   nil,
		endpoint: endpoint,
	}
}

func (this *ProcessorStub) insureClient() {
	this.lock.Lock()
	if this.client == nil {
		client, err := rpc.DialHTTP("tcp", this.endpoint)
		if err != nil {
			log.Fatal("create rpc client error :", err)
		}
		this.client = client
	}
	this.lock.Unlock()
}

func (this *ProcessorStub) Init(t uint64) {
	this.insureClient()
	var input RootTimeArg = RootTimeArg{T: t}
	var reply RootTimeArg
	err := this.client.Call("Root.Init", &input, &reply)
	if err != nil {
		log.Fatal("call Init() error:", err)
	}
}

func (this *ProcessorStub) Advance(t uint64) {
	this.insureClient()
	var input RootTimeArg = RootTimeArg{T: t}
	var reply RootTimeArg
	err := this.client.Call("Root.Advance", &input, &reply)
	if err != nil {
		log.Fatal("call Advance() error:", err)
	}
}

func (this *ProcessorStub) ComputeOutput(t uint64) {
	this.insureClient()
	var input RootTimeArg = RootTimeArg{T: t}
	var reply RootTimeArg
	err := this.client.Call("Root.ComputeOutput", &input, &reply)
	if err != nil {
		log.Fatal("call ComputeOutput() error:", err)
	}
}

func (this *ProcessorStub) PutMessage(message modeling.Message, t uint64) {
	this.insureClient()
	var input RootMessageArg
	var reply RootTimeArg
	input.Message = message
	input.T = t
	err := this.client.Call("Root.PutMessage", &input, &reply)
	if err != nil {
		log.Fatal("call PutMessage() error:", err)
	}
}

func (this *ProcessorStub) GetTN() uint64 {
	this.insureClient()
	var input RootTimeArg
	var reply RootTimeArg
	err := this.client.Call("Root.GetTN", &input, &reply)
	if err != nil {
		log.Fatal("call GetTN() error:", err)
	}
	return reply.T
}
