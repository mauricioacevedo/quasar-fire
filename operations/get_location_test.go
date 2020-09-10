package operations

import (
	"log"
	"testing"

	st "github.com/quasar-fire/structs"
)

func TestLocationOK(t *testing.T) {

	var (
		req = st.TopSecretRequest{
			Satellites: GetInicialSats(),
		}
	)

	satts, err := InitSatelites(req)

	if err != nil {
		t.Errorf("Threre is and error while initializing satellites.")

		return
	}

	distances, _ := ExtractDistancesAndMessages(satts)

	x, y := GetLocation(distances...)

	if x == 0 && y == 0 {
		t.Errorf("Imposible to get the Location")
	}
}

func TestLocationOKOnLocation(t *testing.T) {

	var (
		req = st.TopSecretRequest{
			Satellites: GetSatsOnLocation(),
		}
		challengex = float32(500.0)
		challengey = float32(500.0)
	)

	satts, err := InitSatelites(req)

	if err != nil {
		t.Errorf("Threre is and error while initializing satellites.")

		return
	}

	distances, _ := ExtractDistancesAndMessages(satts)

	x, y := GetLocation(distances...)

	resultx := x / challengex
	resulty := y / challengey

	log.Println("x: ", resultx, ", y: ", resulty)

	//target accuracy: 95%
	if resultx < 0.95 && resulty < 0.95 {
		t.Errorf("Imposible to get the Location")
	}
}

func TestLocationFailed(t *testing.T) {

	var (
		req = st.TopSecretRequest{
			Satellites: GetSatsOnFail(),
		}
		challengex = float32(500.0)
		challengey = float32(500.0)
	)

	satts, err := InitSatelites(req)

	if err != nil {
		t.Errorf("Threre is and error while initializing satellites.")

		return
	}

	distances, _ := ExtractDistancesAndMessages(satts)

	x, y := GetLocation(distances...)

	resultx := x / challengex
	resulty := y / challengey

	log.Println("x: ", resultx, ", y: ", resulty)

	//target accuracy: 95%
	if resultx != 0 || resulty != 0 {
		t.Errorf("Imposible to get the Location")
	}
}

func GetInicialSats() []st.Satellite {
	return []st.Satellite{
		{Name: "kenobi", Distance: 600},
		{Name: "skywalker", Distance: 600},
		{Name: "sato", Distance: 600},
	}
}

func GetSatsOnLocation() []st.Satellite {
	return []st.Satellite{
		{Name: "kenobi", Distance: 1221},
		{Name: "skywalker", Distance: 721},
		{Name: "sato", Distance: 400},
	}
}

func GetSatsOnFail() []st.Satellite {
	return []st.Satellite{
		{Name: "kenobi", Distance: 0},
		{Name: "skywalker", Distance: 100},
		{Name: "sato", Distance: 10},
	}
}
