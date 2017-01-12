package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTemplate = template.Must(template.ParseFiles("home.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	fmt.Println("Server started...")
	err := http.ListenAndServe(*addr, nil)
	defer fmt.Println("Server stoped!!!")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
