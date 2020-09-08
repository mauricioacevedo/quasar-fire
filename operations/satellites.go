package operations

import (
	"fmt"

	st "github.com/quasar-fire/structs"
)

var (
	kenobi     = []float32{-500, -200}
	skywalker  = []float32{100, -100}
	sato       = []float32{500, 100}
	sats       = StartResistance()
	onlineSats = []st.Satellite{}
)

// StartResistance initialize satellites with the default config
func StartResistance() []st.Satellite {
	return []st.Satellite{
		NewSatellite("kenobi", 0, kenobi, nil),
		NewSatellite("skywalker", 0, skywalker, nil),
		NewSatellite("sato", 0, sato, nil),
	}
}

// InitSatelites configure the satellites with the request info for topsecret
func InitSatelites(req st.TopSecretRequest) ([]st.Satellite, error) {

	satellites := []st.Satellite{}

	/**** TODO: Note to future me: i think we need this in the future.
	if len(req.Satellites) != len(sats) {
		err := fmt.Errorf("Number of satellites is different.. is lord Vader trying to hack into our systems?")

		return nil, err
	}*/

	for _, s := range req.Satellites {

		sat, err := FindSatelliteByName(s.Name, sats)

		// NO!!!
		if err != nil {
			return nil, err
		}

		sat.Distance = s.Distance
		sat.Message = s.Message
		satellites = append(satellites, sat)
	}

	onlineSats = satellites

	return satellites, nil
}

// NewSatellite constructor for type Satellite
func NewSatellite(n string, d float32, p []float32, m []string) st.Satellite {
	return st.Satellite{
		Name:     n,
		Distance: d,
		Location: st.Point{X: p[0], Y: p[1]},
		Message:  m,
	}
}

// FindSatelliteByName find a satellite by name from the satellite list
func FindSatelliteByName(n string, satss []st.Satellite) (st.Satellite, error) {

	for _, s := range satss {
		if s.Name == n {
			return s, nil
		}
	}

	err := fmt.Errorf("Satellite %s does not exists", n)

	return st.Satellite{}, err
}

// ExtractDistancesAndMessages function to get distances and messages from already configured satellites
func ExtractDistancesAndMessages(satss []st.Satellite) ([]float32, [][]string) {

	var (
		distances = []float32{}
		msgs      = [][]string{}
	)

	for _, s := range satss {

		distances = append(distances, s.Distance)
		msgs = append(msgs, s.Message)
	}

	return distances, msgs
}
