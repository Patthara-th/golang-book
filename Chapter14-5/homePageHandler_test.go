package home

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/json", strings.NewReader(`{
	"first_name" : "patthara" ,
	"last_name" : "thanathada" ,
	"email" : "patthara.th@tnis.com"}`))

	home := HomePageHandler{}
	home.ServeHTTP(res, req)

	if res.Code != 201 {
		t.Fatalf("Expected status to == 201, but got %d", res.Code)
	}

	user := new(User)
	json.NewDecoder(res.Body).Decode(user)

	if user.Firstname != "patthara" {
		t.Fatalf("Expected Firstname to == patthara, but got sd", user.Firstname)
	}

}
