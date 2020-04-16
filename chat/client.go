package chat

import (
	"fmt"
	"io"
	"websocketsProject/app"
	"websocketsProject/models"

	"golang.org/x/net/websocket"
)

const channelBufSize = 100

var maxId uint = 0

// Chat client.
type Client struct {
	id            uint
	ws            *websocket.Conn
	server        *Server
	chListenWrite chan *MessageToUser
	doneCh        chan bool
	authorized    bool
}

// Create new chat client.
func NewClient(ws *websocket.Conn, server *Server) *Client {

	if ws == nil {
		panic("ws cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	maxId++
	ch := make(chan *MessageToUser, channelBufSize)
	doneCh := make(chan bool)

	return &Client{maxId, ws, server, ch, doneCh, false}
}

func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

func (c *Client) Write(msg *MessageToUser) {
	select {
	case c.chListenWrite <- msg:
	default:
		c.server.Del(c)
		err := fmt.Errorf("client %d is disconnected.", c.id)
		c.server.Err(err)
	}
}

func (c *Client) Done() {
	c.doneCh <- true
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	for {
		select {

		// send message to the client
		case msg := <-c.chListenWrite:
			//msgJson, _ := json.Marshal(msg)
			websocket.JSON.Send(c.ws, msg)

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Client) listenRead() {
	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenWrite method
			return

		// read data from websocket connection
		default:
			var action MessageFromUser
			err := websocket.JSON.Receive(c.ws, &action)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.server.Err(err)
			} else {
				switch action.Action {
				case "AUTH":
					userId, err := app.GetUserId(action.Data["token"].(string))
					if err != nil {
						c.Write(&MessageToUser{Action: "AUTH", Data: map[string]interface{}{"result": false, "error": err.Error()}})
						break
					}

					anonymousId := c.id
					c.id = userId
					c.authorized = true

					c.server.AuthorizeClient(anonymousId, c)
					break
				case "NEW_MESSAGE":
					if !c.authorized {
						break
					}
					message := &models.ChatMessage{
						ChatID:    uint(action.Data["chat_id"].(float64)),
						UserID:    c.id,
						Text:      action.Data["text"].(string),
					}

					models.GetDB().Save(message)

					rows, _ := models.GetDB().
						Raw("select user_id as uid from user_chats where chat_id = ?", action.Data["chat_id"]).
						Rows()

					var receivers []uint
					var uid uint
					for rows.Next() {
						rows.Scan(&uid)
						receivers = append(receivers, uid)
					}
					rows.Close()

					c.server.SendToClients(&MessageToUser{
						Action: "NEW_MESSAGE",
						Data: message,
					}, receivers)
					break
				}
			}
		}
	}
}
