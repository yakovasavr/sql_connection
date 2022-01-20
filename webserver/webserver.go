package webserver

import (
	"fmt"
	"net/http"
	"database/sql"


	"github.com/gorilla/mux"
	"github.com/yakovasavr/sql_connection/store"
	"github.com/yakovasavr/sql_connection/usecase"
)

type WEBServer struct{
	usecase usecase.BookService
}

func (s *WEBServer) Get(w http.ResponseWriter, r *http.Request) {

	book, err := s.usecase.GetBook(r.Context())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, book)
}

func (s *WEBServer) Start() error {
	connStr := "user=postgres dbname=mytest host=localhost password=xxx sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	repo := store.Repo{Conn: db}
	useCase := usecase.BookUsecase{BookRepo: repo}
	server := WEBServer{usecase: useCase}

	router := mux.NewRouter()

	router.HandleFunc("/books", server.Get).
	Methods("GET")

	fmt.Println("starting server at :8000")
	return http.ListenAndServe(":8000", router)
}