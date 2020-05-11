package route_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/raphi011/scores/cmd/api/auth"
	"github.com/raphi011/scores/cmd/api/router"
	"github.com/raphi011/scores/test"
)

func SetupTestServer(t testing.TB) *router.Router {
	r := router.New(
		router.WithMode("debug"),
		router.WithTestRepository(t),
		router.WithEventQueue(),
	)

	return r
}

type testClient struct {
	t             testing.TB
	router        *gin.Engine
	sessionCookie string
	ip            string
}

func newTestClient(t testing.TB) *testClient {
	return &testClient{
		t:      t,
		router: SetupTestServer(t).Build(),
	}
}

func (c *testClient) login() {
	oldIP := c.ip
	c.ip = "127.0.0.1"
	w := c.post("/debug/new-admin", nil)
	c.ip = oldIP

	test.Equal(c.t, "/debug/new-admin expected code %d, got %d", http.StatusNoContent, w.Code)

	w = c.post("/pw-auth", auth.PasswordCredentials{
		Email:    "admin@scores.network",
		Password: "test123",
	})

	test.Equal(c.t, "Cannot login user, expected status %d, got %d", http.StatusOK, w.Code)

	headers := w.Header()

	c.sessionCookie = parseCookie(headers["Set-Cookie"][0])

}

func parseCookie(setCookie string) (cookie string) {
	parts := strings.Split(setCookie, ";")

	keyValue := strings.SplitN(parts[0], "=", 2)

	cookie = keyValue[0] + "=" + keyValue[1]

	return
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
	fillRequest(req, path)

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

func fillRequest(r *http.Request, path string) {
	r.Proto = "HTTP/1.1"
	r.RemoteAddr = "192.168.1.1:80"
	r.Host = "scores.network"
	r.RequestURI = path
}

func (c *testClient) get(path string) *httptest.ResponseRecorder {
	c.t.Helper()

	req, err := http.NewRequest("GET", path, nil)

	test.Check(c.t, "new get-request: %v", err)
	fillRequest(req, path)

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
