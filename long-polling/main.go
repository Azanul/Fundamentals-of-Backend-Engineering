package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	jobs := map[string]bool{}
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		jobId := fmt.Sprintf("job:%d", time.Now().Unix())
		jobs[jobId] = false
		go func() {
			time.Sleep(15 * time.Second)
			jobs[jobId] = true
		}()
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(jobId))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		body := r.URL.Query().Get("jobId")
		if completed, ok := jobs[body]; completed && ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("job completed"))
		} else if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		for !jobs[body] {
			time.Sleep(time.Second)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("job completed"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
