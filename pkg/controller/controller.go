package controller

import (
	"encoding/json"
	"net/http"
	"web_service_ko/pkg/model"
	"web_service_ko/pkg/utils"

	"github.com/gorilla/mux"
)

func ShowAllBookDetails(w http.ResponseWriter, r *http.Request) {
	showBook := model.ShowAllBookDetails()
	res, _ := json.Marshal(showBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ShowBookDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID := vars["id"]
	empDetails := model.ShowBookDetail(empID)
	res, _ := json.Marshal(empDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBookDetail(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBookDetail()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBookDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID := vars["id"]
	UpdateBook := &model.Book{}
	utils.ParseBody(r, UpdateBook)
	empDetailstoUpdate := UpdateBook.UpdateBookDetail(empID)
	res, _ := json.Marshal(empDetailstoUpdate)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBookDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID := vars["id"]
	empDetails := model.DeleteBookDetail(empID)
	res, _ := json.Marshal(empDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
