package model

// Your model here
import (
	"time"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Your model here
type OptionsContract struct {
	Bid            float64   `json:"bid" validate:"gt=0"`            // bid price, should be lower than ask price
	Ask            float64   `json:"ask" validate:"gt=0"`            // ask price
	Type           string    `json:"type" validate:"oneof=call put"` // only accept call or put
	LongShort      string    `json:"long_short" validate:"oneof=long short"`
	StrikePrice    float64   `json:"strike_price" validate:"gt=0"`
	ExpirationDate time.Time `json:"expiration_date" validate:"gt"`
}

func (o *OptionsContract) Validate() error {
	validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(o); err != nil {
		return err
	}

	return nil
}
