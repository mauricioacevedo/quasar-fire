package structs

type TopSecretRequest struct {
	Satellites []Satellite `json:"satellites"`
}

type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
	Location Point    `json:"location"`
}

type TopSecretResponse struct {
	Position Point  `json:"position"`
	Message  string `json:"message"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}