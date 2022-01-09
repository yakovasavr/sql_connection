package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yakovasavr/sql_connection/internal/store"
)

type WEBServer struct{
	config	*Config
	router	*mux.Router
	store	*store.Store
}

func New(config *Config) *WEBServer {
	return &WEBServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *WEBServer) Start() error {
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *WEBServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *WEBServer) handleHello() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s.store.BookList)
	}
}

func (s *WEBServer) configureStore() error {
	st := store.New(s.config.Store)
	err := st.GetTable()
	if err != nil {
		return err
	}
	s.store = st

	return nil
}