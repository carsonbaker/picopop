package main

import (
  "flag"
  "log"
  "text/template"
  "net/http"
  "os"
  "github.com/carsonbaker/copenhagen/util"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("index.html"))

func homeHandler(c http.ResponseWriter, req *http.Request) {
  homeTempl.Execute(c, req.Host)
}

func jqueryHandler(c http.ResponseWriter, req *http.Request) {
  template.Must(template.ParseFiles("assets/js/jquery.min.js")).Execute(c, req.Host)
}

func main() {

  util.Initialize()

  flag.Parse()
  go h.run()
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/ws", wsHandler)
  http.HandleFunc("/assets/js/jquery.min.js", jqueryHandler)

  if os.Getenv("PORT") != "" {
    *addr = ":"+os.Getenv("PORT")
  }

  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }

}
