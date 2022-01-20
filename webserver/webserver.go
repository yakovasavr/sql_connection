package webserver

import (
	"fmt"
	"net/http"
	"database/sql"
	"log"

	"github.com/gorilla/mux"
	"github.com/yakovasavr/sql_connection/store"
	"github.com/yakovasavr/sql_connection/usecase"
	_ "github.com/lib/pq"
)

type WEBServer struct{
	usecase usecase.BookService
}

func (s *WEBServer) GetLastBook(w http.ResponseWriter, r *http.Request) {

	book, err := s.usecase.GetLastBook(r.Context())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, book)
}

func (s *WEBServer) GetYura(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello Yura!")
}

func (s *WEBServer) Start() error {
	//database
	connStr := "user=postgres dbname=mytest host=xx password=xxx sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
    }
	repo := store.Repo{Conn: db}
	///

	useCase := usecase.BookUsecase{BookRepo: repo}
	server := WEBServer{usecase: useCase}

	router := mux.NewRouter()

	router.HandleFunc("/books", server.GetLastBook).Methods("GET")
	router.HandleFunc("/", server.GetYura).Methods("GET")

	fmt.Println("starting server at :8000")
	return http.ListenAndServe(":8000", router)
}
