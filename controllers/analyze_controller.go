package controllers

import (
	"GOLANG-DEV-LOGIC-CHALLENGE-NOPECODE96/model"
	"context"
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"strings"
)

var errTypeNotDefined = errors.New("type not defined properly")

// GraphPoint structure for X & Y values of the risk & reward graph
type GraphPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// AnalysisResult structure for the response body
type AnalysisResult struct {
	MaxLoss         float64      `json:"max_loss"`
	MaxProfit       float64      `json:"max_profit"`
	GraphData       []GraphPoint `json:"graph_data"`
	BreakEvenPoints []float64    `json:"break_even_points"`
}

func AnalyzeController(w http.ResponseWriter, r *http.Request) {
	// Max options contract could be receive is up to four
	ops := make([]model.OptionsContract, 4)

	body, err := r.GetBody()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err = json.NewDecoder(body).Decode(&ops); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for _, o := range ops {
		if err := o.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}

	rs, err := Analyze(r.Context(), ops)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := json.Marshal(&rs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Analyze contains bussiness logic of analyzing the model.OptionsContract
func Analyze(ctx context.Context, ops []model.OptionsContract) (*AnalysisResult, error) {

	// Options Premium
	// it should be noted that the
	// Options Premium was assumed since
	// there's no source that could be use to
	// determine its real value so just randomize it
	opr := rand.New(rand.NewSource(2)).Float64()

	// Set the graph points and break-event points
	bep := make([]float64, len(ops))
	gps := make([]GraphPoint, len(ops))
	for i, op := range ops {
		up, err := calculateUAP(op.Type, op.StrikePrice)
		if err != nil {
			return nil, err
		}

		gps[i].X = up

		pl, err := calculatePOL(op.Type, op.StrikePrice, up, opr)
		if err != nil {
			return nil, err
		}

		gps[i].Y = pl

		bep[i], err = calculateBEP(op.Type, op.StrikePrice, opr)
		if err != nil {
			return nil, err
		}
	}

	// Find max profit and max loss
	mp := gps[0].Y
	ml := gps[0].Y

	for i := 1; i < len(gps); i++ {
		y := gps[i].Y
		mp = math.Max(mp, y)
		ml = math.Min(ml, y)
	}

	// Assign to result
	ar := new(AnalysisResult)
	ar.MaxLoss = ml
	ar.MaxProfit = mp
	ar.GraphData = gps
	ar.BreakEvenPoints = bep

	return ar, nil
}

// calculatePOL calculate loss/profit (Y) of GraphData
// The options premium is randomize with 0 > p >= 2
func calculatePOL(t string, sp, up, opr float64) (float64, error) {
	if strings.Compare(t, "put") == 0 {
		return math.Max(sp-up, 0), nil
	}

	if strings.Compare(t, "call") == 0 {
		c := math.Max(up-sp, 0)
		return c - opr, nil
	}
	return 0, errTypeNotDefined
}

// calculateUAP calcuate the underlying asset price
// Underlying asset price is based on assumption by
// multiply the strike price with random number 0 > r >= 2
//
// Rule: The underlying asset price for put option should
// be lower than strike price while call option is the oppsite
func calculateUAP(t string, sp float64) (float64, error) {
	r := rand.New(rand.NewSource(2)).Float64()

	if strings.Compare(t, "put") == 0 {
		return sp - r, nil
	}

	if strings.Compare(t, "call") == 0 {
		return sp + r, nil
	}

	return 0, errTypeNotDefined
}

// calculateBEP calcuate break-even points based on
// given type of option. Call type option would sum strike
// price and option premium whereare the put would divide it
func calculateBEP(t string, sp, opr float64) (float64, error) {
	if strings.Compare(t, "put") == 0 {
		return sp - opr, nil
	}

	if strings.Compare(t, "call") == 0 {
		return sp + opr, nil
	}

	return 0, errTypeNotDefined
}
