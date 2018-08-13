package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"os"
)

type Random struct {
	Time string
    Host string
    Number int
}

type RandomNumberAPI struct {
}

func (p *RandomNumberAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/random" {
		random := getRandom(w, r)
		json.NewEncoder(w).Encode(random)
		return
	}
	http.NotFound(w, r)
	return
}

func getRandom(w http.ResponseWriter, r *http.Request) (Random) {
	rand.Seed(time.Now().UTC().UnixNano())
	max := 1000
	min := 0
	host, _ := os.Hostname()
	msg := Random{ time.Now().Format(time.RFC850), host, rand.Intn(max-min) }
	return msg
}

func main() {
	mux := &RandomNumberAPI{}
	http.ListenAndServe(":8000", mux)
}
