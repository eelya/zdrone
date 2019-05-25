package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/eelya/zdrone/pkg/proc"
	log "github.com/sirupsen/logrus"
)

var count = 0

func handler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("Handle %s request to %s from %s", r.Method, r.URL, r.RemoteAddr)
	r.ParseForm()
	f, _ := strconv.Atoi(r.Form.Get("x"))
	p := proc.NewProc(19)
	fmt.Fprintf(w, "hey %v %v, count is %d", p.Add(f), p.Mul(f), count)
	count++
}

func main() {
	log.SetLevel(log.DebugLevel)

	port, ok := os.LookupEnv("Z_PORT")
	if !ok {
		port = "8989"
	}

	http.HandleFunc("/foo", handler)

	addr := net.JoinHostPort("0.0.0.0", port)
	log.Infof("Starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
