package utils

import (
	"testing"

	"github.com/quasar-fire/structs"
)

func TestDistanceOK(t *testing.T) {

	var (
		p1  = structs.Point{X: 0, Y: 0}
		p2  = structs.Point{X: 1, Y: 0}
		ans = float32(1)
	)

	answer := Distance(p1, p2)

	if answer != ans {
		t.Errorf("Distancia invalida")
	}
}

func TestDistanceSamePoint(t *testing.T) {

	var (
		p1  = structs.Point{X: 0, Y: 0}
		p2  = structs.Point{X: 0, Y: 0}
		ans = float32(0)
	)

	answer := Distance(p1, p2)

	if answer != ans {
		t.Errorf("Distancia invalida")
	}
}

func TestPow(t *testing.T) {

	var (
		val    = float32(2)
		answer = float32(4)
	)

	ans := Pow(val)

	if answer != ans {
		t.Errorf("Distancia invalida")
	}
}

func TestSqrt(t *testing.T) {

	var (
		val    = float32(4)
		answer = float32(2)
	)

	ans := Sqrt(val)

	if answer != ans {
		t.Errorf("Distancia invalida")
	}
}
