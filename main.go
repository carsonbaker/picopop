package main

import (
  "flag"
  "log"
  "text/template"
  "net/http"
  "os"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("index.html"))

func homeHandler(c http.ResponseWriter, req *http.Request) {
  homeTempl.Execute(c, req.Host)
}

func main() {

  flag.Parse()
  go h.run()
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/ws", wsHandler)

  if os.Getenv("PORT") != "" {
    *addr = ":"+os.Getenv("PORT")
  }

  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }

}
