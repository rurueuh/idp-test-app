package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = "dev"

func main() {
	host, _ := os.Hostname()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"version\":%q,\"pod\":%q}\n", version, host)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, page, version, host)
	})

	log.Printf("idp-test-app version=%s listening on :%s", version, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

const page = `<!doctype html>
<html>
<head><meta charset="utf-8"><title>idp-test-app</title></head>
<body style="font-family:system-ui,sans-serif;background:#0d1226;color:#e5e7eb;margin:0;display:flex;min-height:100vh;align-items:center;justify-content:center">
  <div style="text-align:center">
    <h1 style="margin:0 0 .5rem">idp-test-app</h1>
    <p style="font-size:2rem;margin:.25rem 0">version <b style="color:#34d399">%s</b></p>
    <p style="color:#94a3b8">served by %s</p>
  </div>
</body>
</html>
`
