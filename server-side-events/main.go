package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "text/event-stream")
		for i := 0; i < 30; i++ {
			w.Write([]byte(fmt.Sprintf("data: Hello from the server --- %d\n\n", i)))
			flusher.Flush()
			time.Sleep(time.Second)
		}
		return
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
