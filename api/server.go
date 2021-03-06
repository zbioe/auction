package api

import (
	"net/http"

	"github.com/zbioe/auction/handlers"
	"github.com/zbioe/auction/db"
)

type server struct {
	mux	*http.ServeMux
	Db	*db.Mem
}

func NewServer() *server {
	return &server{
		mux: http.NewServeMux(),
		Db: db.NewMem(),
	}
}

func (s *server) SetupHandlers() {
	s.mux.HandleFunc("/bid", handlers.NewBid(s.Db).Serve)
	s.mux.HandleFunc("/stats", handlers.NewStats(s.Db).Serve)
}

func (s *server) Start(port string) error {
	return http.ListenAndServe(port, s.mux)
}

func (s *server) Run(port string) error {
	s.SetupHandlers()
	return s.Start(port)
}
