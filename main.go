package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/siuyin/serversentevents_try/public"
)

func main() {
	log.Println("Starting to send events to clients...")
	http.Handle("/", http.FileServerFS(public.Content))
	http.HandleFunc("/events", eventStreamer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func eventStreamer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Accel-Buffering", "no")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(200)
	for {
		select {
		case <-r.Context().Done():
			fmt.Println("client disconnected")
			return
		default:
			io.WriteString(w, "data: "+time.Now().Format("15:04:05")+"\n\n")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
			w.(http.Flusher).Flush()
		}
	}
}
