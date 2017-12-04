package main

import (
  "flag"
  "log"
  "net/http"
  "os"
  "./util"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {

  util.Initialize()

  flag.Parse()
  go h.run()
  http.HandleFunc("/ws", wsHandler)

  if os.Getenv("PORT") != "" {
    *addr = ":"+os.Getenv("PORT")
  }

  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }

}
