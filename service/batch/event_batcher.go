package batch

import (
	"sync"
	"time"
	"sync/atomic"
)

type EventHandler interface{
	Handle(interface{})
}

type EventBatcher struct {
	Period   time.Duration
	lock     *sync.Mutex
	stopFlag int32
	buff     []*batchedMessage
	handler  EventHandler
	deleting bool
}

type batchedMessage struct {
	data           interface{}
	iterationsLeft int
}

//batcher T시간 간격으로 handler에게 메세지를 전달해준다
//deleting option에 따라서 전달한 message를 지울껀지 아니면 계속 남겨둘지를 설정한다.

//buff: protos queue
//lock: sync
//period: T time
//stopflag: to stop batcher
//handler: messaging handler
func NewGRPCMessageBatcher(period time.Duration, handler EventHandler, deleting bool) *EventBatcher{

	gb := &EventBatcher{
		buff:     make([]*batchedMessage, 0),
		lock:     &sync.Mutex{},
		Period:   period,
		stopFlag: int32(0),
		handler:  handler,
		deleting: deleting,
	}

	go gb.periodicEmit()

	return gb
}

//tested
func (gb *EventBatcher)Add(message interface{}){

	gb.lock.Lock()
	defer gb.lock.Unlock()

	iteration := len(gb.buff)
	gb.buff = append(gb.buff, &batchedMessage{data: message, iterationsLeft: iteration})
}

//tested
func (gb *EventBatcher)Stop(){
	atomic.StoreInt32(&(gb.stopFlag), int32(1))
}

//tested
func (gb *EventBatcher)Size() int{

	gb.lock.Lock()
	defer gb.lock.Unlock()
	return len(gb.buff)
}

//tested
func (gb *EventBatcher) toDie() bool {
	return atomic.LoadInt32(&(gb.stopFlag)) == int32(1)
}

//tested
func (gb *EventBatcher) periodicEmit() {
	for !gb.toDie() {
		time.Sleep(gb.Period)
		gb.lock.Lock()
		gb.emit()
		gb.lock.Unlock()
	}
}

//tested
func (gb *EventBatcher) emit() {

	if gb.toDie(){
		return
	}

	if len(gb.buff) == 0{
		return
	}

	for _, message := range gb.buff{
		gb.handler.Handle(message.data)
	}

	gb.vacate()
}

//test
func (gb *EventBatcher) vacate() {

	if gb.deleting{
		gb.buff = gb.buff[0:0]
	}
}