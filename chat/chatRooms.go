package chat

import (
	"golang.org/x/net/websocket"
	"sync"
	"log"
)

// CRooms :- chat room
type CRooms struct {
	sync.RWMutex
	mapRooms map[string]map[*websocket.Conn]bool
}

func NewCRooms() *CRooms {
	return &CRooms{
		mapRooms: make(map[string]map[*websocket.Conn]bool),
	}
}

func (cr *CRooms) add(roomName string, ws *websocket.Conn) {
	// Get clients for room
	room, ok := cr.mapRooms[roomName]
	// Initialize ws pool
	if !ok {
		log.Println("Creating room for clients")
		room = make(map[*websocket.Conn]bool)
		cr.mapRooms[roomName] = room
	}
	// add websocket client to room
	room[ws] = true
}

func (cr *CRooms) Add(roomName string, ws *websocket.Conn) {
	cr.Lock()
	defer cr.Unlock()
	cr.add(roomName, ws)
}

func (cr *CRooms) delete(roomName string, ws *websocket.Conn) {
	// Get clients for room
	room, ok := cr.mapRooms[roomName]
	// Initialize ws pool
	if !ok {
		return
	}
	// add websocket client to room
	delete(room, ws)
}

func (cr *CRooms) Delete(roomName string, ws *websocket.Conn) {
	cr.Lock()
	defer cr.Unlock()
	cr.delete(roomName, ws)
}

func (cr *CRooms) get(roomName string) map[*websocket.Conn]bool{
	// Get clients for room
	room, ok := cr.mapRooms[roomName]
	// Initialize ws pool
	if !ok {
		log.Println("Creating room for clients")
		room = make(map[*websocket.Conn]bool)
		cr.mapRooms[roomName] = room
	}
	// add websocket client to room
	return room
}

func (cr *CRooms) Get(roomName string) map[*websocket.Conn]bool{
	cr.RLock()
	defer cr.RUnlock()
	return cr.get(roomName)
}

func (cr *CRooms) clientsInRoom(roomName string) int{
	room, ok := cr.mapRooms[roomName]
	if !ok {
		return 0
	}
	return len(room)
}

func (cr *CRooms) ClientsInRoom(roomName string) int{
	cr.RLock()
	defer cr.RUnlock()
	return cr.clientsInRoom(roomName)
}