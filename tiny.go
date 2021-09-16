package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	dirPtr := flag.String("serve", ".", "Select a directory to serve!")
	portPtr := flag.Int("port", -1, "Select a port to run server on!")

	flag.Parse()

	if *portPtr == -1 {
		log.Println("Require port number to run!")
		os.Exit(-1)
	}

	fs := http.FileServer(http.Dir(*dirPtr))
	http.Handle("/", fs)

	port := ":" + strconv.Itoa(*portPtr)
	log.Println("Listening on "+port, "... and serving directory: "+*dirPtr)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
