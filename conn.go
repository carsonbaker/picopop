package main

import (
  "github.com/gorilla/websocket"
  "net/http"
  "./songs"
)

type connection struct {
  // The websocket connection.
  ws *websocket.Conn

  // Buffered channel of outbound messages.
  send chan []byte
}

func (c *connection) reader() {
  for {
    _, message, err := c.ws.ReadMessage()
    
    if err != nil {
      break
    }

    buf := songs.RenderJSONSong(message)
    // buf := songs.RenderBeefyFartSong()

    c.send <- buf

  }
  c.ws.Close()
}

func (c *connection) writer() {
  for audio := range c.send {
    err := c.ws.WriteMessage(websocket.BinaryMessage, audio)
    if err != nil {
        break
    }
  }
  c.ws.Close()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
  ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
  if _, ok := err.(websocket.HandshakeError); ok {
    http.Error(w, "Not a websocket handshake", 400)
    return
  } else if err != nil {
    return
  }
  c := &connection{send: make(chan []byte, 256), ws: ws}
  h.register <- c
  defer func() { h.unregister <- c }()
  go c.writer()
  c.reader()
}
