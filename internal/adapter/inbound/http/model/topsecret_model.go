package model

type Healthy struct {
	Sucess bool `json:"success"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type TopsecretResponse struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

type TopSecretSplitResponse struct {
	SavedReceivedMessage bool `json:"savedReceivedMessage"`
}
