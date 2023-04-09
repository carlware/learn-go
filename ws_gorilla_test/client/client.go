package client

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	pongWait = 500 * time.Second
)

type baseClient struct {
	name string
	host string
	// buffer for bursts or spikes in data
	buffer            int
	connectionRetries int
	timeout           time.Duration
}

func NewClient(host string, opts ...Option) (*baseClient, error) {
	const op = "ws.New"

	c := &baseClient{
		host:              host,
		buffer:            1000,
		connectionRetries: 3,
		timeout:           time.Second * 60,
	}
	for _, opt := range opts {
		if err := opt.applyOption(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *baseClient) createConnection(ctx context.Context) (*websocket.Conn, error) {
	var err error
	var wsConn *websocket.Conn

	for i := 0; i < c.connectionRetries; i++ {
		wsConn, _, err = websocket.DefaultDialer.Dial(c.host, nil)
		if err == nil {
			break
		}
	}

	if err != nil {
		wrappedErr := fmt.Errorf("failed to dial websocket: %w", err)
		return wsConn, wrappedErr
	}

	for j := 0; j < c.connectionRetries; j++ {
		err = c.subscribe(ctx, wsConn)
		if err == nil {
			return wsConn, nil
		}
	}

	return wsConn, fmt.Errorf("failed to subscribe to websocket channel: %w", err)
}

func (c *baseClient) subscribe(ctx context.Context, wsConn *websocket.Conn) error {
	const op = "baseClient.subscribe"

	err := wsConn.WriteMessage(websocket.TextMessage, []byte(`subscribe`))
	if err != nil {
		return fmt.Errorf("failed to write subscription message: %w", err)
	}

	//ctxDeadline, _ := context.WithTimeout(ctx, time.Second*5)
	//
	//counter := len(requests)
	//
	//for {
	//	select {
	//	case <-ctxDeadline.Done():
	//		return ez.New(op, ez.EINTERNAL, "Failed to get subscription confirmations", nil)
	//
	//	default:
	//		_, bs, err := wsConn.ReadMessage()
	//		if err != nil {
	//			return ez.Wrap(op, err)
	//		}
	//
	//		err = c.handler.VerifySubscriptionResponse(bs)
	//		if err != nil {
	//			return ez.Wrap(op, err)
	//		}
	//
	//		counter--
	//
	//		if counter == 0 {
	//			return nil
	//		}
	//	}
	//}

	return nil
}

func (c *baseClient) Listen(ctx context.Context) (<-chan string, error) {
	wsConn, cErr := c.createConnection(ctx)
	if cErr != nil {
		return nil, cErr
	}

	strChan := make(chan string, c.buffer)
	//
	//wsConn.SetReadDeadline(time.Now().Add(pongWait))
	//wsConn.SetPongHandler(func(appData string) error {
	//	fmt.Printf("pong handler recv %s\n", appData)
	//	wsConn.SetReadDeadline(time.Now().Add(pongWait))
	//	return nil
	//})
	//wsConn.WriteMessage(websocket.PingMessage, []byte{})
	//
	//wsConn.SetPingHandler(func(appData string) error {
	//	fmt.Printf("ping handler recv %s\n", appData)
	//	return nil
	//})

	//wsConn.SetReadDeadline()
	//go func() {
	//	ticker := time.NewTicker(1000 * time.Millisecond)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			if err := wsConn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
	//				fmt.Println("Can't write to socket...disconnecting")
	//				return
	//			}
	//		}
	//	}
	//
	//}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(strChan)
				return
			default:
				if wsConn == nil {
					time.Sleep(100 * time.Millisecond)
					wsConn, _ = c.createConnection(ctx)
					continue
				}
				_, bs, bErr := wsConn.ReadMessage()
				if bErr != nil {
					fmt.Printf("type err %T\n", bErr)
					if _, ok := bErr.(*websocket.CloseError); ok {
						wsConn.Close()
						fmt.Println("closing connection")
						wsConn = nil
						continue
					}
					log.Printf("error reading message %s", bErr)
					continue
				}

				strChan <- string(bs)
			}
		}
	}()

	return strChan, nil
}
