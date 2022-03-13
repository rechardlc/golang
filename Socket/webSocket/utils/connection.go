package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

// Myws 定义ws
func Myws(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
	H.r <- c
	go c.writer()
	c.reader()
	defer func() {
		c.data.Type = "logout"
		userList = del(userList, c.data.User)
		c.data.UserList = userList
		c.data.Content = c.data.User
		dataB, _ := json.Marshal(c.data)
		H.b <- dataB
		H.r <- c
	}()
}

func (c *connection) writer() {
	for message := range c.sc {
		if err := c.ws.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}
	if err := c.ws.Close(); err != nil {
		return
	}
}

var userList []string

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			H.r <- c
			break
		}
		if err := json.Unmarshal(message, &c.data); err != nil {
			return
		}
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			userList = append(userList, c.data.User)
			c.data.UserList = userList
			dataB, _ := json.Marshal(c.data)
			H.b <- dataB
		case "user":
			c.data.Type = "user"
			dataB, _ := json.Marshal(c.data)
			H.b <- dataB
		case "logout":
			c.data.Type = "logout"
			userList = del(userList, c.data.User)
			dataB, _ := json.Marshal(c.data)
			H.b <- dataB
			H.r <- c
		default:
			fmt.Print("========default================")
		}
	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var nSlice []string
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			nSlice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(nSlice)
	return nSlice
}
