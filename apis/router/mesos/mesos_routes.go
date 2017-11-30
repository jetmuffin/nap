package mesos

import "net/http"

func (mr mesosRouter) handleUsage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("{\"cpus\": 3.5,\"mem\": 1024,\"disk\": 2048}"))
}
