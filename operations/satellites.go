package operations

import (
	"fmt"

	st "github.com/quasar-fire/structs"
)

var (
	kenobi       = []float32{-500, -200}
	skywalker    = []float32{100, -100}
	sato         = []float32{500, 100}
	sats         = StartResistance()
	onlineSats   = []st.Satellite{}
	onlineSatsV2 = StartResistance()
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

func GetOneSatelliteSplit(name string) (st.Satellite, error) {
	return FindSatelliteByName(name, onlineSatsV2)
}

func UpdateSatellitesSplit(satUpdated st.Satellite) []st.Satellite {
	satsNew := []st.Satellite{}

	for _, s := range onlineSatsV2 {
		if s.Name == satUpdated.Name {
			s = satUpdated
		}

		satsNew = append(satsNew, s)
	}

	onlineSatsV2 = satsNew

	return onlineSatsV2
}

func ResetSatelliteSplit() ([]st.Satellite, error) {
	onlineSatsV2 = StartResistance()

	for _, s := range onlineSatsV2 {

		sat, err := FindSatelliteByName(s.Name, onlineSatsV2)

		// NO!!!
		if err != nil {
			return nil, err
		}

		sat.Distance = s.Distance
		sat.Message = s.Message
		onlineSatsV2 = append(onlineSatsV2, sat)
	}

	return onlineSatsV2, nil

}

func GetAllSatelliteSplit() []st.Satellite {

	return onlineSatsV2
}

func InitSatelitesFromDB(dbSats []st.Satellite) ([]st.Satellite, error) {

	if len(dbSats) != len(sats) {
		err := fmt.Errorf("Number of satellites is different.. is lord Vader trying to hack into our systems?")

		return nil, err
	}

	onlineSats = dbSats

	return dbSats, nil
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
