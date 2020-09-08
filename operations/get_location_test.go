package operations

import (
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

func GetInicialSats() []st.Satellite {
	return []st.Satellite{
		{Name: "kenobi", Distance: 600},
		{Name: "skywalker", Distance: 600},
		{Name: "sato", Distance: 600},
	}
}
