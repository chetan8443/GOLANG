package UnitTest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	a "v1/WebServer"
)

func TestAdd(t *testing.T) {
	req, err := http.NewRequest("GetDetails", "/marks/:id", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(req)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.GetDetails)
	handler.ServerHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned: got %v want %v", status, http.StatusOK)
	}
	body := rr.Body.String()
	fmt.Println(body)
}
