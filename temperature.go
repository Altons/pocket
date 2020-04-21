package pocket

import (
	"errors"
	"strings"
)

//Temperature
type temperature struct{}

//Convert is function converting different temperature units
// It can convert Celsius(c), Kelvin(k) and Fahrenheit(f)
func (t temperature) Convert(value float64, unitIn string, unitOut string) (float64, error) {
	units := []string{"c", "k", "f"}
	unitI, err := find(units, strings.ToLower(unitIn))
	if err != nil {
		return 0.0, errors.New("invalid unit:" + unitIn)
	}
	unitO, err := find(units, strings.ToLower(unitOut))
	if err != nil {
		return 0.0, errors.New("invalid unit:" + unitOut)
	}

	switch {
	case unitI == unitO:
		return value, nil
	case unitI == "c" && unitO == "f":
		return value*9.0/5.0 + 32.0, nil
	case unitI == "k" && unitO == "f":
		return (value-273.15)*9.0/5.0 + 32.0, nil
	case unitI == "f" && unitO == "c":
		return (value - 32.0) * 5.0 / 9.0, nil
	case unitI == "c" && unitO == "k":
		return value + 273.15, nil
	case unitI == "k" && unitO == "c":
		return value - 273.15, nil
	case unitI == "f" && unitO == "k":
		return (value-32.0)*5.0/9.0 + 273.15, nil
	default:
		return 0.0, errors.New("Unknown error")
	}

}
