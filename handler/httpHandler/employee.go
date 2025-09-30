package httphandler

import (
	"backend/driver"
	"backend/models"
	"backend/repository"
	"backend/repository/employee"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

)

func NewEmployeeHandler(db *driver.DB) *Employee {
	return &Employee{
		repo: employee.NewSqlEmployeeRepo(db.SQL),
	}
}

type Employee struct {
	repo repository.EmployeeRepo
}

func (e *Employee) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, err := e.repo.Fetch(r.Context())
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, payload)
}

func (e *Employee) Create(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}
	json.NewDecoder(r.Body).Decode(&employee)

	newId, err := e.repo.Insert(r.Context(), &employee)

	if err != nil {
		respondwithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created", "employeeId": strconv.FormatInt(newId, 6)})
}

func (e *Employee) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["emp_Id"])
	data := models.Employee{}
	json.NewDecoder(r.Body).Decode(&data)
	data.Model.ID = uint(id);
	payload, err := e.repo.Update(r.Context(), &data)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, payload)
}

func (e *Employee) GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["emp_Id"])
	payload, err := e.repo.GetById(r.Context(), int64(id))
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, payload)
}

func (e *Employee) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["emp_Id"])
	_, err := e.repo.Delete(r.Context(), int64(id))
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Deleted Successfully"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondwithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
