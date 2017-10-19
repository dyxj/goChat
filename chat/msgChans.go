package chat

import (
	"log"
	"sync"
)

type MsgData struct {
	Msg  string `json:"msg"`
	User string `json:"user"`
}

type MsgChans struct {
	sync.RWMutex
	mapMsgChanel map[string]chan MsgData
}

func NewMsgChans() *MsgChans {
	return &MsgChans{
		mapMsgChanel: make(map[string]chan MsgData),
	}
}

func (ms *MsgChans) get(roomName string) chan MsgData {
	msgChan, ok := ms.mapMsgChanel[roomName]
	if !ok {
		log.Println("Creating broadcast channel")
		msgChan = make(chan MsgData, 20)
		ms.mapMsgChanel[roomName] = msgChan
	}
	return msgChan
}

func (ms *MsgChans) Get(roomName string) chan MsgData {
	ms.RLock()
	defer ms.RUnlock()
	return ms.get(roomName)
}

func (ms *MsgChans) delete(roomName string) {
	delete(ms.mapMsgChanel, roomName)
}

func (ms *MsgChans) Delete(roomName string) {
	ms.Lock()
	defer ms.Unlock()
	ms.delete(roomName)
}
