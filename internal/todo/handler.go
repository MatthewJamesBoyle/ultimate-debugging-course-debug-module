package todo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	Description string `json:"description"`
}

type Response struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type server struct {
	svc *Service
}

func NewServer(_ *Service) (server, error) {

	return server{}, nil
}

func (s *server) CreateToDo(w http.ResponseWriter, r *http.Request) {
	var todo Request

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nID, err := s.svc.CreateTODO(r.Context(), todo.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Description: todo.Description, ID: nID})
}

func (s *server) GetToDoHandler(w http.ResponseWriter, r *http.Request) {
	// Extracting the ID from the URL
	path := strings.TrimPrefix(r.URL.Path, "/todo/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todoItem, err := s.svc.GetTODO(r.Context(), id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoItem)
}
