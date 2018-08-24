// Webserver backend for showCal
package main

// Package externed methods have capital first letter
import {
	"fmt"
	"io"
	"net/http"
	"github.com/golang/glog"
}

var port = "3333"

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello from port %q \n", port))
}

// Entry point of the package main
func main() {
	duck := "QuackyQuack"
	// Handle a route
	http.HandleFunc("/", handle)
	glog.Info("%q on localhost:%q\n", duck, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}