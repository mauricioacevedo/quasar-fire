package structs

type TopSecretRequest struct {
	Satellites []Satellite `json:"satellites"`
}

type TopSecretRequestSplit struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
	Location Point    `json:"location"`
}

type SatelliteDB struct {
	ID        int16   `json:"id"`
	Name      string  `json:"name"`
	Distance  float32 `json:"distance"`
	Message   string  `json:"message"`
	LocationX float32 `json:"locationx"`
	LocationY float32 `json:"locationy"`
}
type TopSecretResponse struct {
	Position Point  `json:"position"`
	Message  string `json:"message"`
}

type TopSecretPostResponse struct {
	Message string `json:"message"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
