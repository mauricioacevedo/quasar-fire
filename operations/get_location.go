package operations

import (
	"fmt"

	"github.com/quasar-fire/structs"
	u "github.com/quasar-fire/utils"
)

// GetLocation returns the ship location
func GetLocation(distances ...float32) (x, y float32) {

	x, y, err := solution(onlineSats)

	if err != nil {
		return 0, 0
	}

	return x, y
}

func solution(sats []structs.Satellite) (x, y float32, err error) {

	if !conditions(sats) {
		err := fmt.Errorf("Satellites not in range")

		return 0, 0, err
	}

	x1, y1, r1 := sats[0].Location.X, sats[0].Location.Y, sats[0].Distance
	x2, y2, r2 := sats[1].Location.X, sats[1].Location.Y, sats[1].Distance
	x3, y3, r3 := sats[2].Location.X, sats[2].Location.Y, sats[2].Distance

	//kenobi and skywalker
	A := 2*x2 - 2*x1
	B := 2*y2 - 2*y1
	C := u.Pow(r1) - u.Pow(r2) - u.Pow(x1) + u.Pow(x2) - u.Pow(y1) + u.Pow(y2)
	D := 2*x3 - 2*x2
	E := 2*y3 - 2*y2
	F := u.Pow(r2) - u.Pow(r3) - u.Pow(x2) + u.Pow(x3) - u.Pow(y2) + u.Pow(y3)

	x = (C*E - F*B) / (E*A - B*D)
	y = (C*D - A*F) / (B*D - A*E)

	return x, y, nil
}

func conditions(sats []structs.Satellite) bool {

	if len(sats) < 3 {
		return false
	}

	d1 := u.Distance(sats[0].Location, sats[1].Location)
	if d1 > sats[0].Distance+sats[1].Distance {
		return false
	}
	d2 := u.Distance(sats[1].Location, sats[2].Location)
	if d2 > sats[1].Distance+sats[2].Distance {
		return false
	}

	d3 := u.Distance(sats[0].Location, sats[2].Location)
	if d3 > sats[0].Distance+sats[2].Distance {
		return false
	}

	return true
}
