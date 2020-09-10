package services

import (
	"fmt"

	"github.com/quasar-fire/operations"
	st "github.com/quasar-fire/structs"
)

func TopSecretService(req st.TopSecretRequest) (*st.TopSecretResponse, error) {

	sats, err := operations.InitSatelites(req)

	//They are trying to deceive us..
	if err != nil {

		msg := fmt.Errorf("Mission is compromised!!! %s", err.Error())

		return nil, msg
	}

	distances, msgs := operations.ExtractDistancesAndMessages(sats)

	x, y := operations.GetLocation(distances...)

	if x == 0 && y == 0 {
		msg := fmt.Errorf("No Location found or Satellites not Online, try to refine distances and make sure you have 3 sats..")

		return nil, msg
	}

	message := operations.GetMessage(msgs...)

	if message == "" {
		msg := fmt.Errorf("No Message detected")

		return nil, msg
	}

	res := st.TopSecretResponse{
		Position: st.Point{
			X: x,
			Y: y,
		},
		Message: message,
	}

	return &res, nil
}

func TopSecretSplitPostService(satelliteName string, req st.TopSecretRequestSplit) (*st.TopSecretPostResponse, error) {

	sat, err := operations.GetOneSatelliteSplit(satelliteName)

	if err != nil {

		msg := fmt.Errorf("Mission is compromised!! Satellite %s does not exists", err.Error())

		return nil, msg
	}
	//just store the thing..

	sat.Distance = req.Distance
	sat.Message = req.Message

	operations.UpdateSatellitesSplit(sat)

	res := st.TopSecretPostResponse{
		Message: fmt.Sprintf("Ok"),
	}

	return &res, nil
}

func TopSecretSplitResetService() (*st.TopSecretPostResponse, error) {
	_, err := operations.ResetSatelliteSplit()

	if err != nil {

		msg := fmt.Errorf("Mission is compromised!! there was an error: %s", err.Error())

		return nil, msg
	}

	res := st.TopSecretPostResponse{
		Message: fmt.Sprintf("Ok"),
	}

	return &res, nil
}

func TopSecretSplitGetService() (*st.TopSecretResponse, error) {

	sats := operations.GetAllSatelliteSplit()

	distances, msgs := operations.ExtractDistancesAndMessages(sats)

	x, y := operations.GetLocationSplit(distances...)

	if x == 0 && y == 0 {
		msg := fmt.Errorf("No Location found or Satellites not Online, try to refine distances and make sure you have 3 satellites online. Han")

		return nil, msg
	}

	message := operations.GetMessage(msgs...)

	if message == "" {
		msg := fmt.Errorf("No Message detected")

		return nil, msg
	}

	res := st.TopSecretResponse{
		Position: st.Point{
			X: x,
			Y: y,
		},
		Message: message,
	}

	return &res, nil
}
