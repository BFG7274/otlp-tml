package main

import (
	"bfg7274/otlp-tml-gateway/pkg/logs"
	"bfg7274/otlp-tml-gateway/pkg/store"
	"bfg7274/otlp-tml-gateway/pkg/users"
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
)

type server struct {
	u users.UserClient
	s store.StoreClient
	l logs.LogClient
}

var tracer = otel.Tracer("otlp-tml-gateway")

var s *server

func getUser(w http.ResponseWriter, req *http.Request) {
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, span := tracer.Start(ctxt, "insert-user")
	defer span.End()
	
}

func getItem(w http.ResponseWriter, req *http.Request) {

}

func buyItem(w http.ResponseWriter, req *http.Request) {

}

func main() {
	s = &server{
		u: *users.NewUserClient(),
		s: *store.NewStoreClient(),
		l: *logs.NewLogClient(),
	}
	h := http.NewServeMux()
	h.HandleFunc("/getUser", getUser)
	h.HandleFunc("/getItem", getItem)
	h.HandleFunc("/buy", buyItem)
	httpd := &http.Server{Addr: "0.0.0.0:9381", Handler: h}
	httpd.ListenAndServe()
	s.l.WriteLog("Hello", "info")
}
