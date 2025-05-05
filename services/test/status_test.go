package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServicesStatus(t *testing.T) {
	for _, tc := range []serviceTestCases{
		{name: "version", urlFormat: "http://localhost:8001/%s/version"},
		{name: "liveness", urlFormat: "http://localhost:8001/%s/liveness"},
		{name: "readiness", urlFormat: "http://localhost:8001/%s/readiness"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, service := range services {
				t.Run(service, func(t *testing.T) {
					url := fmt.Sprintf(tc.urlFormat, service)
					t.Logf("Запрос к %s", url)

					//nolint: gosec
					resp, err := http.Get(url)
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
