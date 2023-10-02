package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prests/hypermedia-golang/router"
)

type testRouter struct {
	routes []router.Route
}

func (tr *testRouter) Routes() []router.Route {
	return tr.routes
}

type customFS struct{}

func (ch *customFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestNewServer(t *testing.T) {
	tr := &testRouter{}
	tr.routes = []router.Route{
		router.NewRoute("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	}

	fsHandler := &customFS{}

	server := NewServer(fsHandler, tr)

	t.Run("it return 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, 200)
	})

	t.Run("it return 200 for assets request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/assets/htmx.min.js", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, 200)
	})
}

func assertStatus(t testing.TB, received int, expected int) {
	t.Helper()
	if (received != expected) {
		t.Errorf("Incorrect statuses: received %d expected %d", received, expected)
	}
}