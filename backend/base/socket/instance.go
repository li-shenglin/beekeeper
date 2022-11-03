package socket

import (
	"backend/common"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

var chanPool = sync.Pool{
	New: func() interface{} {
		return make(chan *Message, 1)
	},
}

type Instance struct {
	model     InstanceModel
	Status    InstanceStatus
	processor *Processor
	handler   func(int, Param) (any, error)
	invokeMap sync.Map

	// client properties
	Address string

	// server properties
	LastHeartBeat time.Time
}

func NewServerInstance(conn net.Conn, f func(int, Param) (any, error)) *Instance {
	instance := &Instance{
		model:         SeverInstance,
		Status:        InstanceOnline,
		handler:       f,
		LastHeartBeat: time.Now(),
	}
	instance.processor = NewProcessor(conn, instance)
	return instance
}

func NewClientInstance(address string, f func(int, Param) (any, error)) *Instance {
	return &Instance{
		Address: address,
		model:   ClientInstance,
		Status:  InstanceUnderLine,
		handler: f,
	}
}

func (instance *Instance) Invoke(obj *Parameter) (Param, error) {
	if instance.Status != InstanceOnline {
		return nil, fmt.Errorf("instance is not health")
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	message := NewMessage(data)
	c := chanPool.Get().(chan *Message)
	defer chanPool.Put(c)
	seqID := string(message.SeqID)
	instance.invokeMap.Store(seqID, c)
	defer instance.invokeMap.Delete(seqID)
	instance.processor.Send(message)
	select {
	case r := <-c:
		return r.GetData()
	case <-time.After(time.Second * 2):
		return nil, fmt.Errorf("rpc timeout")
	}
}

func (instance *Instance) Start() {
	if instance.model == ClientInstance {
		conn, err := net.Dial("tcp", instance.Address)
		common.PanicNotNull(err)
		instance.processor = NewProcessor(conn, instance)

		go instance.ServerHeart()
	}

	instance.processor.Start()
	log.Infof("Instance[%v] start", instance.Address)
}

func (instance *Instance) ServerHeart() {
	for {
		time.Sleep(time.Second * 2)
		if instance.Status == InstanceUnHealth {
			conn, err := net.Dial("tcp", instance.Address)
			if err != nil {
				log.Infof("Conncet Server Failed: %v", err)
				continue
			}
			instance.processor = NewProcessor(conn, instance)
			instance.Status = InstanceOnline
			instance.processor.Start()
		}
	}
}

func (instance *Instance) Accept(message *Message, processor *Processor) {
	if message.Type == 1 { // result
		v, loaded := instance.invokeMap.LoadAndDelete(string(message.SeqID))
		if loaded {
			v.(chan *Message) <- message
		}
		return
	}
	data, err := message.GetData()
	if err != nil {
		data = data.SetErr(err)
	} else {
		res, err := instance.handler(int(data.Opt), data)
		if err != nil {
			data = data.SetErr(err)
		} else {
			data = data.SetData(res)
		}
	}
	marshal, _ := json.Marshal(data)
	returnMessage := NewReturnMessage(message.SeqID, marshal)
	processor.Send(returnMessage)
}
