package test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	for _, tc := range []serviceTestCases{
		{name: "ping", urlFormat: "http://localhost:8000/%s/template_service.Service/Ping"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, service := range services {
				t.Run(service, func(t *testing.T) {
					url := fmt.Sprintf(tc.urlFormat, service)
					t.Logf("Запрос к %s", url)

					//nolint: gosec
					resp, err := http.Post(url,
						"application/json",
						strings.NewReader(fmt.Sprintf(`{"ping": "%s"}`, service)),
					)
					require.NoError(t, err)
					t.Log(resp.Status)
					require.Equal(t, http.StatusOK, resp.StatusCode)

					bytes, err := io.ReadAll(resp.Body)
					require.NoError(t, err)
					t.Log(string(bytes))
				})
			}
		})
	}
}
