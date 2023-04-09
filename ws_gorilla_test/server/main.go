package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic("error")
	}

	defer wsConn.Close()

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {

		timeout := time.Second * 5
		wsConn.SetWriteDeadline(time.Now().Add(timeout))

		err = wsConn.WriteJSON(struct {
			Name string `json:"name"`
			ID string `json:"id"`
		}{
			Name: "carlos",
			ID: time.Now().String(),
		})
		if err != nil {
			fmt.Println("breaking", err)
			break
		}
	}


	//c, err := upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//	log.Print("upgrade:", err)
	//	return
	//}
	//defer c.Close()
	//
	//

	//c.SetPingHandler(func(appData string) error {
	//	fmt.Printf("ping handler recv %s\n", appData)
	//	wErr := c.WriteMessage(websocket.PongMessage, []byte(`pong`))
	//	if wErr != nil {
	//		log.Printf("failed to write ping message %s\n", wErr)
	//	}
	//	return wErr
	//})
	//
	//c.SetPongHandler(func(appData string) error {
	//	fmt.Printf("pong handler recv %s\n", appData)
	//	return nil
	//})
	//
	//for {
	//	mt, message, err := c.ReadMessage()
	//	if err != nil {
	//		log.Println("read:", err)
	//		break
	//	}
	//	log.Printf("recv: %s", message)
	//	err = c.WriteMessage(mt, message)
	//	if err != nil {
	//		log.Printf("write error: %s", err)
	//		break
	//	}
	//
	//	if string(message) == "subscribe" {
	//		go func() {
	//			ticker := time.NewTicker(3000 * time.Millisecond)
	//			for {
	//				select {
	//				case <-ticker.C:
	//					cErr := c.WriteMessage(mt, []byte(fmt.Sprintf("new message at %s", time.Now().Format("2006-01-02 15:04:05"))))
	//					if cErr != nil {
	//						log.Printf("write error: %s", cErr)
	//						return
	//					}
	//					//time.Sleep(3 * time.Second)
	//				}
	//			}
	//		}()
	//
	//	}
	//}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", ws)
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
