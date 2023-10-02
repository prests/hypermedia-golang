package navigation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestLandingPage(t *testing.T) {
	approvals.UseFolder("testdata")

	nr := &navigationRouter{
		templateBasePath: "../../..",
	}

	t.Run("it returns index template as html", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		nr.landingPage(response, request)

		approvals.VerifyString(t, response.Body.String())
	})
}