package main

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"os"
	"time"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	// Get the current epoch time
	epochTime := time.Now().Unix()

	// Create a response with the epoch time
	w.Header().Set("Content-Disposition", "attachment; filename=utc")
	w.Header().Set("Content-Type", "text/plain")

	// Write the epoch time to the response
	fmt.Fprintf(w, "%d\n", epochTime)
}

func main() {

	addrPtr := flag.String("addr", "127.0.0.1", "listen address")
	portPtr := flag.Int("port", 2013, "listen port")
	flag.Parse()

	http.HandleFunc("/utc", downloadFile)

	// Start the server on port 8080
	listenArg := fmt.Sprintf("%s:%d", *addrPtr, *portPtr)
	fmt.Fprintf(os.Stderr, "utcd started at %s\n", listenArg)
	if err := http.ListenAndServe(listenArg, nil); err != nil {
		log.Fatal(err)
	}
}
