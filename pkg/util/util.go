package util

import (
	"errors"
	"math"
)

func Find(s []string, val string) (string, error) {
	for _, item := range s {
		if val == item {
			return val, nil
		}
	}
	return val, errors.New("unrecognized unit")
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
