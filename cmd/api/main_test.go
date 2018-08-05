package main

import (
	"os"
	"testing"
)

var a app

func TestMain(m *testing.M) {
	// a = app{
	// 	production: false,
	// }
	// a.Initialize()

	// ensureTableExists()

	code := m.Run()

	// clearTable()

	os.Exit(code)
}

// func executeRequest(req *http.Request) *httptest.ResponseRecorder {
// 	rr := httptest.NewRecorder()
// 	a.Router.ServeHTTP(rr, req)

// 	return rr
// }

// func checkResponseCode(t *testing.T, expected, actual int) {
// 	if expected != actual {
// 		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
// 	}
// }

// func TestEmptyTable(t *testing.T) {
// 	// clearTable()

// 	req, _ := http.NewRequest("GET", "/matches", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	if body := response.Body.String(); body != "[]" {
// 		t.Errorf("Expected an empty array. Got %s", body)
// 	}
// }
