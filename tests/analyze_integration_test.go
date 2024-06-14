package tests

import (
	"GOLANG-DEV-LOGIC-CHALLENGE-NOPECODE96/controllers"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestIntegration_Analyze(t *testing.T) {
	samples := []struct {
		body       string
		desc       string
		statusCode int
	}{
		{
			statusCode: http.StatusOK,
			desc:       "It should be OK",
			body: `[
  {
    "strike_price": 100, 
    "type": "call", 
    "bid": 10.05, 
    "ask": 12.04, 
    "long_short": "long", 
    "expiration_date": "2025-12-17T00:00:00Z"
  },
  {
    "strike_price": 102.50, 
    "type": "call", 
    "bid": 12.10, 
    "ask": 14, 
    "long_short": "long", 
    "expiration_date": "2025-12-17T00:00:00Z"
  },
  {
    "strike_price": 103, 
    "type": "put", 
    "bid": 14, 
    "ask": 15.50, 
    "long_short": "short", 
    "expiration_date": "2025-12-17T00:00:00Z"
  },
  {
    "strike_price": 105, 
    "type": "put", 
    "bid": 16, 
    "ask": 18, 
    "long_short": "long", 
    "expiration_date": "2025-12-17T00:00:00Z"
  }
]`,
		},
		{
			statusCode: http.StatusBadRequest,
			desc:       "It should be Bad Request",
			body: `[
  {
    "strike_price": 0, 
    "type": "Call", 
    "bid": 0, 
    "ask": 0, 
    "long_short": "long", 
    "expiration_date": "2025-12-17T00:00:00Z"
  }
]`,
		},
		{
			statusCode: http.StatusInternalServerError,
			desc:       "It should be Internal Server Error",
			body:       ``,
		},
	}

	for _, s := range samples {
		t.Run(s.desc, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/analyze", strings.NewReader(s.body))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Add("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(controllers.AnalyzeController)

			hr.ServeHTTP(rr, req)

			if status := rr.Code; status != s.statusCode {
				t.Errorf("/analyze status = %v, want = %v", status, s.statusCode)
			}
		})
	}
}
