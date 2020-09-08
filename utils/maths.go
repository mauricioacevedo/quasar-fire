package utils

import (
	"math"

	"github.com/quasar-fire/structs"
)

func Pow(val float32) float32 {

	result := math.Pow(float64(val), 2)

	return float32(result)
}

func Sqrt(val float32) float32 {

	result := math.Sqrt(float64(val))

	return float32(result)
}

func Distance(p1, p2 structs.Point) float32 {

	X := Pow(p1.X - p2.X)
	Y := Pow(p1.Y - p2.Y)

	return Sqrt(X + Y)
}
