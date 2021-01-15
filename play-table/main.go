package main

import (
	"fmt"

	"encoding/json"
)

type Table struct {
	Ranges []Range `json:"ranges"`
	Name   string  `json:"name"`
	Max    int64   `json:"max"`
}

type Range struct {
	Min  int64 `json:"min"`
	Max  int64 `json:"max"`
	Rate int64 `json:"rate"`
}

const cent = 100

var RateAltTable = []Range{
	{
		Min:  0,
		Max:  25 * 1000 * cent,
		Rate: 800,
	},
	{
		Min:  25 * 1000 * cent,
		Max:  40 * 1000 * cent,
		Rate: 850,
	},
	{
		Min:  40 * 1000 * cent,
		Max:  50 * 1000 * cent,
		Rate: 900,
	},
	{
		Min:  50 * 1000 * cent,
		Max:  75 * 1000 * cent,
		Rate: 950,
	},
	{
		Min:  75 * 1000 * cent,
		Max:  100 * 1000 * cent,
		Rate: 1000,
	},
	{
		Min:  100 * 1000 * cent,
		Max:  150 * 1000 * cent,
		Rate: 1050,
	},
	{
		Min:  150 * 1000 * cent,
		Max:  200 * 1000 * cent,
		Rate: 1100,
	},
	{
		Min:  200 * 1000 * cent,
		Max:  350 * 1000 * cent,
		Rate: 1150,
	},
	{
		Min:  350 * 1000 * cent,
		Max:  500 * 1000 * cent,
		Rate: 1200,
	},
	{
		Min:  500 * 1000 * cent,
		Max:  750 * 1000 * cent,
		Rate: 1250,
	},
	{
		Min:  750 * 1000 * cent,
		Max:  1 * 1000 * 1000 * cent,
		Rate: 1300,
	},
	{
		Min:  1 * 1000 * 1000 * cent,
		Max:  1750 * 1000 * cent,
		Rate: 1350,
	},
	{
		Min:  1750 * 1000 * cent,
		Max:  2500 * 1000 * cent,
		Rate: 1400,
	},
	{
		Min:  2500 * 1000 * cent,
		Max:  3750 * 1000 * cent,
		Rate: 1450,
	},
}

var RateTable = []Range{
	{
		Min:  0,
		Max:  15 * 1000 * cent,
		Rate: 500, // 0.05 (5%)
	},
	{
		Min:  15 * 1000 * cent,
		Max:  25 * 1000 * cent,
		Rate: 600, // 0.06 (6%)
	},
	{
		Min:  25 * 1000 * cent,
		Max:  40 * 1000 * cent,
		Rate: 700,
	},
	{
		Min:  40 * 1000 * cent,
		Max:  60 * 1000 * cent,
		Rate: 800,
	},
	{
		Min:  60 * 1000 * cent,
		Max:  80 * 1000 * cent,
		Rate: 900,
	},
	{
		Min:  80 * 1000 * cent,
		Max:  100 * 1000 * cent,
		Rate: 950,
	},
	{
		Min:  100 * 1000 * cent,
		Max:  150 * 1000 * cent,
		Rate: 1000,
	},
	{
		Min:  150 * 1000 * cent,
		Max:  250 * 1000 * cent,
		Rate: 1100,
	},
	{
		Min:  250 * 1000 * cent,
		Max:  500 * 1000 * cent,
		Rate: 1200,
	},
	{
		Min:  500 * 1000 * cent,
		Max:  1 * 1000 * 1000 * cent,
		Rate: 1300,
	},
	{
		Min:  1 * 1000 * 1000 * cent,
		Max:  2 * 1000 * 1000 * cent,
		Rate: 1400,
	},
}

const Year float64 = 365

const MAX_RATE = 1500

var AltRater = NewRater(MAX_RATE, RateAltTable)
var BasicRater = NewRater(MAX_RATE, RateTable)

type Rater func(int64) int64

func NewRater(max int64, ranges []Range) Rater {
	return func(val int64) int64 {
		for _, r := range ranges {
			if val >= r.Min && val < r.Max {
				return r.Rate
			}
		}
		return max
	}
}

var autoInvestMonths = []int{
	1,
	4,
	7,
	10,
}

func IsCollectionMonth(month int) bool {
	for _, m := range autoInvestMonths {
		if month == m {
			return true
		}
	}
	return false
}

func main() {
	out, err := json.Marshal(&Table{
		Name:   "",
		Max:    100,
		Ranges: RateTable,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", out)
}
