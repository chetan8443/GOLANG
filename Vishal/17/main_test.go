package main

import (
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vishal/studentInfo/routes"
)

type TStudent struct {
	Id     int
	Name   string
	Marks  int
	Result string
}

func Test_GetById(t *testing.T) {

	req, err := http.NewRequest("GET", "/result/GetById?id=99", nil)
	if err != nil {
		t.Error(err) 
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetById)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Id":4,"Name":"StudentD","Marks":82,"Result":"Paas With Dist"}`
	fmt.Println(rr.Body.String())
	ir:=`{"message":"Student  details Does Not Found Please Enter valid Id!"}`
	if rr.Body.String() != expected && rr.Body.String()!=ir{
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	// bd:=`{"message":"Student  details Does Not Found Please Enter valid Id!"}`
	// if rr.Body.String() != bd {
	// 	t.Errorf("expeceted %v but got %v",rr.Body.String(),bd)
	// }
}

func Test_GetDataBase(t *testing.T) {
	var db = routes.GetMYSQLDB()

	if db == nil {
		t.Errorf("could not connect to database")
	}
}

