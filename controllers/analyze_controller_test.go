package controllers

import (
	"GOLANG-DEV-LOGIC-CHALLENGE-NOPECODE96/model"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var ops = []model.OptionsContract{
	{
		Ask:            12.04,
		Bid:            10.05,
		Type:           "call",
		LongShort:      "long",
		StrikePrice:    100,
		ExpirationDate: time.Date(2025, time.December, 17, 23, 0, 0, 0, time.UTC),
	},
	{
		Ask:            14,
		Bid:            12.10,
		Type:           "call",
		LongShort:      "long",
		StrikePrice:    102.50,
		ExpirationDate: time.Date(2025, time.December, 17, 23, 0, 0, 0, time.UTC),
	},
	{
		Ask:            15.50,
		Bid:            14,
		Type:           "put",
		LongShort:      "short",
		StrikePrice:    103,
		ExpirationDate: time.Date(2025, time.December, 17, 23, 0, 0, 0, time.UTC),
	},
	{
		Ask:            18,
		Bid:            16,
		Type:           "put",
		LongShort:      "long",
		StrikePrice:    105,
		ExpirationDate: time.Date(2025, time.December, 17, 23, 0, 0, 0, time.UTC),
	},
}

func TestAnalyzeController(t *testing.T) {
	br, err := json.Marshal(ops)
	if err != nil {
		t.Fatal(err)
	}

	body := bytes.NewReader(br)
	req, err := http.NewRequest(http.MethodPost, "/analyze", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	hr := http.HandlerFunc(AnalyzeController)

	hr.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Logf("response: %v", rr.Body.String())
		t.Errorf("AnalyzeController().StatusCode = %v, want = %v", status, http.StatusOK)
	}
}

func TestAnalyze(t *testing.T) {
	ctx := context.TODO()
	_, err := Analyze(ctx, ops)
	if err != nil {
		t.Errorf("Analyze() = %v, want = nil", err)
	}
}
