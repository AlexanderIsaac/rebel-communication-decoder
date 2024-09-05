package model

type Distance struct {
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
