package model

import (
	"testing"
	"time"
)

func TestValidateSucceed(t *testing.T) {
  oc := &OptionsContract{
    Ask: 12.04,
    Bid: 10.05,
    Type: "call",
    LongShort: "long",
    StrikePrice: 100,
    ExpirationDate: time.Now().Add(1 * time.Hour),
  }

  err := oc.Validate()
  if err != nil {
    t.Errorf("oc.Validate() = %v, want = nil", err)
  }
}

func TestValidateFailed(t *testing.T) {
  oc := &OptionsContract{
    Ask: 0,
    Bid: 0,
    Type: "other",
    LongShort: "other",
    StrikePrice: 0,
    ExpirationDate: time.Now().Add(-10 * time.Hour),
  }

  err := oc.Validate()
  if err == nil {
    t.Errorf("oc.Validate() = %v, want = OptionsContract validation failed", err)
  }
}

