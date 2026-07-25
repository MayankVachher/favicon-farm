// Serves the farm: index.html and the favicon SVGs.
package main

import (
	"embed"
	"flag"
	"log"
	"net/http"
)

//go:embed index.html *.svg
var site embed.FS

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "listen address")
	flag.Parse()

	log.Printf("favicon-farm serving on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, http.FileServerFS(site)))
}
