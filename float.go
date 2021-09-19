package gotypes

import (
	"fmt"
	"math"
)

type Float float64

type FloatConfig struct {
	Decimals int
	Symbol   string
}

func (f *Float) ToFloor(config *FloatConfig) float64 {
	if f != nil {
		if config != nil {
			pow := getPow(config.Decimals)
			return math.Floor(float64(*f)*pow) / pow
		}
		return math.Floor(float64(*f))
	}
	return 0
}

func (f *Float) ToRound(config *FloatConfig) float64 {
	if f != nil {
		if config != nil {
			pow := getPow(config.Decimals)
			return math.Round(float64(*f)*pow) / pow
		}
		return math.Round(float64(*f))
	}
	return 0
}

func (f *Float) ToCeil(config *FloatConfig) float64 {
	if f != nil {
		if config != nil {
			pow := getPow(config.Decimals)
			return math.Ceil(float64(*f)*pow) / pow
		}
		return math.Ceil(float64(*f))
	}
	return 0
}

func (f *Float) ToMoney(config *FloatConfig) string {
	if f != nil {
		if config != nil {
			config = &FloatConfig{
				Decimals: 2,
				Symbol:   "$",
			}
		}
		pow := getPow(config.Decimals)
		result := math.Round(float64(*f)*pow) / pow
		return fmt.Sprintf("%v%v", config.Symbol, result)
	}
	return ""
}

func (f *Float) ToPrint(message string) {
	if f != nil {
		fmt.Println(message, *f)
	}
}

func getPow(d int) float64 {
	return math.Pow(10, float64(d))
}
