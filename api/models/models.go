package models

type Coordinates struct {
	X   string `json:"x"`
	Y   string `json:"y"`
	Z   string `json:"z"`
	Vel string `json:"vel"`
}

type Location struct {
	Loc float64 `json:"loc"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}
