package response

type NearestDriver struct {
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
}
