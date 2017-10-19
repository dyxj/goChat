package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"goChat/chat"
	"golang.org/x/net/websocket"
)

var (
	rooms  = chat.NewCRooms()
	mChans = chat.NewMsgChans()
)

func main() {
	// Routes
	rt := mux.Router{}
	fs := http.FileServer(http.Dir("./dist"))

	rt.StrictSlash(true).Handle("/cr/{chatroom}", websocket.Handler(newWebsocket))
	rt.PathPrefix("/").Handler(fs)

	// Start server
	if err := http.ListenAndServe(":1234", &rt); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func newWebsocket(ws *websocket.Conn) {
	// Close connection when function returns
	defer ws.Close()
	//log.Println(ws.Request())
	//log.Println(ws.Request().Header.Get("Sec-WebSocket-Protocol"))

	vars := mux.Vars(ws.Request())
	cRoom := vars["chatroom"]
	log.Println(cRoom)

	// Add client to a room
	rooms.Add(cRoom, ws)

	// Get mChans channel for room
	mBroad := mChans.Get(cRoom)

	// Start up message broadcaster for chat room
	go broadcastMessage(cRoom)

	// Setup receiver
	for {
		var msg chat.MsgData
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			log.Printf("Cant receive error: %v\n", err)
			rooms.Delete(cRoom, ws)
			// Close Broadcast Channel if there are no more clients in chatroom
			if rooms.ClientsInRoom(cRoom) <= 0 {
				log.Println("No clients in room")
				close(mBroad)
				mChans.Delete(cRoom)
			}
			break
		}
		log.Println("Received:", msg)
		mBroad <- msg
	} // end while

}

// Modify broadcastMessage
func broadcastMessage(roomName string) {
	for {
		msgChan := mChans.Get(roomName)
		// Get message from mChans channel
		msg, ok := <-msgChan
		if !ok {
			log.Println("Broadcast channel for", roomName, "closed")
			break
		}
		log.Printf("Broadcast %v: %v\n", roomName, msg)
		// Get map of clients in room
		roomClients := rooms.Get(roomName)
		// Send to connected clients
		for client := range roomClients {
			go func(client *websocket.Conn){
				if err := websocket.JSON.Send(client, &msg); err != nil {
					log.Printf("Cant Send error: %v\n", err)
					client.Close()
					rooms.Delete(roomName, client)
				}
			}(client)
		} // loop over clients in room

	} // end while
	log.Println("exit mChans functions", roomName)
}
