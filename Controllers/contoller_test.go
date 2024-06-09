package controllers

import (
	"bytes"
	"empdb/db"
	"empdb/models"
	"encoding/json"
	"strconv"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHeartBeat(t *testing.T) {
	router := gin.Default()
	router.GET("/heartbeat", Heart_Beat)

	w := performRequest(router, "GET", "/heartbeat", nil)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Working" {
		t.Errorf("Expected body 'Working' but got '%s'", w.Body.String())
	}
}

func TestAddEmployee(t *testing.T) {

	router := gin.Default()
	router.POST("/addEmployee", AddEmployee)

	var emp =models.Employee{Name:"Jothsdbfsh",Position:1,Salary:12000.00}

	jsonEmp,_:=json.Marshal(emp)

	w := performRequest(router, "POST", "/addEmployee", jsonEmp)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var res models.Employee
	err:=json.Unmarshal(w.Body.Bytes() ,&res)

	if err!=nil{
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if res.Name != emp.Name {
		t.Errorf("Expected body 'success' but got '%s'", w.Body.String())
	}
}
func TestGetEmployeeByID(t *testing.T) {

	router := gin.Default()
	router.GET("/employee/:id", GetEmployeeByID)

	// Insert a sample employee for testing
	conn := new(db.Connection)
	conn.CreateDB()
	emp := models.Employee{
		Name:  "Janfsdfsdfe Doe",
		Position:   28,
		Salary: 450000.58,
	}
	conn.Insert(&emp)

	id := strconv.Itoa(int(emp.ID))

	w := performRequest(router, "GET", "/employee/"+id, nil)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var responseEmp models.Employee
	err := json.Unmarshal(w.Body.Bytes(), &responseEmp)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if emp.Name != responseEmp.Name {
		t.Errorf("Expected name %s but got %s", emp.Name, responseEmp.Name)
	}
}
func TestUpdateEmployee(t *testing.T) {
	

	router := gin.Default()
	router.PUT("/employee/:id", UpdateEmployee)

	// Insert a sample employee for testing
	conn := new(db.Connection);
	conn.CreateDB()
	emp := models.Employee{
		Name:  "Jane Doe",
		Position:   28,
		Salary: 450700.58,
	}

	ok:=conn.Insert(&emp)
	if !ok {
		t.Errorf("Expected success but got %v", ok)
	}

	id := strconv.Itoa(int(emp.ID))
	updateEmp := map[string]interface{}{
		"name": "Updated Name",
		"position":  30,
	}

	body, err := json.Marshal(updateEmp)
	if err != nil {
		t.Errorf("Failed to Marshal response body: %v", err)
	}

	w := performRequest(router, "PUT", "/employee/"+id, body)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var responseEmp models.Employee

	json.Unmarshal(w.Body.Bytes(), &responseEmp)


	if emp.Name == responseEmp.Name {
		t.Errorf("Expected name %s but got %s", emp.Name, responseEmp.Name)
	}
}
func TestDeleteEmployee(t *testing.T) {

	router := gin.Default()
	router.DELETE("/employee/:id", DeleteEmployee)

	// Insert a sample employee for testing
	conn := new(db.Connection);
	conn.CreateDB()
	emp := models.Employee{
		Name:  "Jane Doe",
		Position:   28,
		Salary: 450700.58,
	}

	ok:=conn.Insert(&emp)
	if !ok{
		t.Errorf("Failed to insert employee")
	}

	id := strconv.Itoa(int(emp.ID))
	w := performRequest(router, "DELETE", "/employee/"+id, nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

}

func TestGetEmployees(t *testing.T) {

	router := gin.Default()
	router.GET("/employees/:page/:perpage", GetEmployees)

	// Insert a sample employee for testing
	emp := models.Employee{
		Name:  "Jane Doe",
		Position:   28,
		Salary: 556644.55,
	}
	conn:=new(db.Connection);
	conn.CreateDB()
	conn.Insert(&emp)

	page := "1"
	perpage := "10"
	w := performRequest(router, "GET", "/employees/"+page+"/"+perpage, nil)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}
