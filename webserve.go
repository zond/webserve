package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	port := flag.Int("port", 8080, "What port to listen to")
	dir := flag.String("dir", wd, "What dir to serve")

	flag.Parse()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), http.FileServer(http.Dir(*dir))))
}
