package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/raphi011/scores/repo/sql"
	"github.com/raphi011/scores/test"
)

func testServices(t testing.TB) *handlerServices {
	repos, _ := sql.RepositoriesTest(t)

	return servicesFromRepositories(repos, false)
}

func SetupTestServer(t testing.TB) *gin.Engine {
	a := app{
		production: false,
		conf:       nil,
		log:        logrus.New(),
	}
	services := testServices(t)

	return initRouter(a, services)
}

type testClient struct {
	t             testing.TB
	router        *gin.Engine
	sessionCookie string
}

func newTestClient(t testing.TB) *testClient {
	return &testClient{
		t:      t,
		router: SetupTestServer(t),
	}
}

func (c *testClient) login() {
	w := c.post("/debug/new-admin", nil)

	test.Equal(c.t, "/debug/new-admin expected code %d, got %d", http.StatusNoContent, w.Code)

	w = c.post("/pw-auth", credentialsDto{
		Email:    "admin@scores.network",
		Password: "test123",
	})

	test.Equal(c.t, "Cannot login user, expected status %d, got %d", http.StatusOK, w.Code)

	headers := w.Header()

	setCookie := headers["Set-Cookie"][0]

	keyValue := strings.Split(setCookie, "=")

	c.sessionCookie = keyValue[0] + "=" + keyValue[1]
}

func (c *testClient) post(path string, body interface{}) *httptest.ResponseRecorder {
	c.t.Helper()

	var jsonBody io.Reader

	if body != nil {
		marshalled, err := json.Marshal(body)
		test.Check(c.t, "post() unmarshal: %v", err)
		jsonBody = bytes.NewReader(marshalled)
	}

	req, err := http.NewRequest("POST", path, jsonBody)

	test.Check(c.t, "new post-request: %v", err)

	c.setCookie(req)
	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)
	return w
}

func (c *testClient) setCookie(req *http.Request) {
	if c.sessionCookie != "" {
		req.Header.Set("cookie", c.sessionCookie)
	}
}

func (c *testClient) get(path string) *httptest.ResponseRecorder {
	c.t.Helper()

	req, err := http.NewRequest("GET", path, nil)
	test.Check(c.t, "new get-request: %v", err)

	c.setCookie(req)
	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)
	return w
}

func TestGetUnauthorizedTournament(t *testing.T) {
	client := newTestClient(t)

	w := client.get("/tournaments")

	test.Equal(t, "/tournaments expected status %d, got %d", http.StatusUnauthorized, w.Code)
}

func TestGetTournamentsWithoutFilters(t *testing.T) {
	client := newTestClient(t)
	client.login()

	w := client.get("/tournaments")

	test.Equal(t, "/tournaments expected status %d, got %d", http.StatusOK, w.Code)
}
